package buttplug

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pidurentry/buttplug-go/logging"
	"github.com/pidurentry/buttplug-go/message"

	// Load all message classes
	_ "github.com/pidurentry/buttplug-go/enumerationmsg"
	_ "github.com/pidurentry/buttplug-go/handshakemsg"
	_ "github.com/pidurentry/buttplug-go/statusmsg"
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
	logging.GetLogger().Infof("Connecting to %s", url)

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		logging.GetLogger().Warningf("Failed to connect to %s: %v", url, err)
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

	logging.GetLogger().Info("Listening for messages...")
	defer logging.GetLogger().Info("Listener stopped")

	for {
		messageType, message, err := buttplug.conn.ReadMessage()
		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				logging.GetLogger().Info("Recieved close message")
				return
			default:
				logging.GetLogger().Errorf("Failed to read message: %v", err)
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
		logging.GetLogger().Errorf("Recieved unexpected %s message: %v", buttplug.lookupMessageType(messageType))
		buttplug.err <- &UnexpectedWebsocketMessageType{}
	}
}

func (buttplug *buttplug) lookupMessageType(messageType int) string {
	switch messageType {
	case websocket.TextMessage:
		return "TextMessage"
	case websocket.BinaryMessage:
		return "BinaryMessage"
	case websocket.CloseMessage:
		return "CloseMessage"
	case websocket.PingMessage:
		return "PingMessage"
	case websocket.PongMessage:
		return "PongMessage"
	default:
		return "<unknown>"
	}
}

func (buttplug *buttplug) processJSON(message []byte) {
	logging.GetLogger().Tracef("Processing json:\n%s", message)

	dec := json.NewDecoder(bytes.NewReader(message))
	for dec.More() {
		token, err := dec.Token()
		if err != nil {
			logging.GetLogger().Errorf("Unexpected message token: %v", err)
			buttplug.err <- err
			return
		}

		msgType, ok := token.(string)
		if !ok {
			continue
		}

		msg, err := buttplug.create(msgType)
		if err != nil {
			logging.GetLogger().Errorf("Failed to create %s message object", msgType, err)
			buttplug.err <- err
			continue
		}

		if err := dec.Decode(msg); err != nil {
			logging.GetLogger().Errorf("Failed to decode %s message", msgType, err)
			buttplug.err <- err
			continue
		}

		logging.GetLogger().Debugf("Recieved %s message: %#v", msgType, msg.(Message))
		buttplug.msg <- msg.(Message)
	}
}

func (buttplug *buttplug) create(msgType string) (interface{}, error) {
	factory, ok := message.Repository[msgType]
	if !ok {
		return nil, &UnknownMessageType{}
	}
	return factory(), nil
}

func (buttplug *buttplug) Send(messages ...Message) error {
	data := make([]interface{}, len(messages))
	for index, message := range messages {
		logging.GetLogger().Debugf("Sending message: %#v", message)
		data[index] = message.Serilize()
	}

	json, err := json.Marshal(data)
	if err != nil {
		logging.GetLogger().Errorf("Failed to encode messages: %v", err)
		return err
	}

	logging.GetLogger().Tracef("Sending json message:\n%s", json)
	return buttplug.conn.WriteMessage(websocket.TextMessage, json)
}

func (buttplug *buttplug) Recieve() <-chan Message {
	return buttplug.msg
}

func (buttplug *buttplug) Error() <-chan error {
	return buttplug.err
}

func (buttplug *buttplug) Close(timeout time.Duration) error {
	logging.GetLogger().Info("Closing connection...")
	defer logging.GetLogger().Info("Connection closed")

	err := buttplug.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		logging.GetLogger().Errorf("Failed to send close messages: %v", err)
		return err
	}

	select {
	case <-buttplug.done:
	case <-time.After(timeout):
		logging.GetLogger().Warning("Connection failed to close")
		return &Timeout{}
	}

	return nil
}
