package buttplug

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type Buttplug interface {
	Send(messages ...Message) error
	Recieve() <-chan Message
	Error() <-chan error
	Close(timeout time.Duration) error
}

type buttplug struct {
	conn *websocket.Conn
	done chan interface{}
	msg  chan Message
	err  chan error
}

func Dial(url string) (Buttplug, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	buttplug := &buttplug{
		conn: conn,
		done: make(chan interface{}, 0),
		msg:  make(chan Message, 0),
		err:  make(chan error, 0),
	}

	go buttplug.listen()
	return buttplug, nil
}

func (buttplug *buttplug) listen() {
	defer close(buttplug.done)
	for {
		messageType, message, err := buttplug.conn.ReadMessage()
		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				return
			default:
				buttplug.err <- err
			}
		}
		go buttplug.processMessage(messageType, message)
	}
}

func (buttplug *buttplug) processMessage(messageType int, message []byte) {
	switch messageType {
	case websocket.TextMessage:
		buttplug.processJSON(message)
	default:
		buttplug.err <- errors.New("unexpected websocket message type")
	}
}

func (buttplug *buttplug) processJSON(message []byte) {
	dec := json.NewDecoder(bytes.NewReader(message))
	for dec.More() {
		token, err := dec.Token()
		if err != nil {
			buttplug.err <- err
			return
		}

		msgType, ok := token.(string)
		if !ok {
			continue
		}

		msg, err := NewMessage(msgType)
		if err != nil {
			buttplug.err <- err
			continue
		}

		if err := dec.Decode(msg); err != nil {
			buttplug.err <- err
			continue
		}

		buttplug.msg <- msg.(Message)
	}
}

func (buttplug *buttplug) Send(messages ...Message) error {
	data := make([]interface{}, len(messages))
	for index, message := range messages {
		data[index] = message.Serilize()
	}

	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", json)
	return buttplug.conn.WriteMessage(websocket.TextMessage, json)
}

func (buttplug *buttplug) Recieve() <-chan Message {
	return buttplug.msg
}

func (buttplug *buttplug) Error() <-chan error {
	return buttplug.err
}

func (buttplug *buttplug) Close(timeout time.Duration) error {
	err := buttplug.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}

	select {
	case <-buttplug.done:
	case <-time.After(timeout):
		return errors.New("timeout reached")
	}

	return nil
}
