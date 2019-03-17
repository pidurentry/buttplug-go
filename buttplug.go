package buttplug

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Buttplug interface {
	Send(messages ...Message) error
}

type buttplug struct {
	conn *websocket.Conn
}

func Dial(url string) (Buttplug, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	buttplug := &buttplug{
		conn: conn,
	}

	go buttplug.listen()

	return buttplug, nil
}

func (buttplug *buttplug) listen() {
	for {
		go buttplug.processMessage(buttplug.conn.ReadMessage())
	}
}

func (buttplug *buttplug) processMessage(messageType int, message []byte, err error) {
	fmt.Printf("%d: %s (%#v)\n", messageType, message, err)
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
