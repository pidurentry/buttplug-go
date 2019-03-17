package buttplug

import "errors"

type Message interface {
	Serilize() interface{}
}

var MessageRepository = make(map[string]func() interface{})

func NewMessage(msgType string) (interface{}, error) {
	factory, ok := MessageRepository[msgType]
	if !ok {
		return nil, errors.New("unknown message type")
	}
	return factory(), nil
}
