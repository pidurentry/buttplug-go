package buttplug

import "github.com/gorilla/websocket"

type Buttplug interface{}

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

	// TODO: Initilize listener

	return buttplug, nil
}
