package buttplug

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/pidurentry/buttplug-go/handshakemsg"

	"github.com/gorilla/websocket"
)

type Buttplug interface {
	Send(messages ...Message) error
	Wait()
	Close(timeout time.Duration) error
}

type buttplug struct {
	done chan interface{}
	conn *websocket.Conn
}

func Dial(url string) (Buttplug, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	buttplug := &buttplug{
		done: make(chan interface{}, 0),
		conn: conn,
	}

	go buttplug.listen()

	if err := buttplug.Send(&handshakemsg.RequestServerInfo{1, "buttplug-go", handshakemsg.MESSAGE_VERSION}); err != nil {
		return buttplug, err
	}

	return buttplug, nil
}

func (buttplug *buttplug) listen() {
	defer close(buttplug.done)

	for {
		messageType, message, err := buttplug.conn.ReadMessage()
		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
			default:
				fmt.Printf("%#v\n", err)
			}
			return
		}
		go buttplug.processMessage(messageType, message)
	}
}

func (buttplug *buttplug) processMessage(messageType int, message []byte) {
	fmt.Printf("%d: %s\n", messageType, message)
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

func (buttplug *buttplug) Wait() {
	<-buttplug.done
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
