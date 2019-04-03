package buttplug

type MessageId int

type Message interface {
	Id() MessageId
	Serilize() interface{}
}

var MessageRepository = make(map[string]func() interface{})

func NewMessage(msgType string) (interface{}, error) {
	factory, ok := MessageRepository[msgType]
	if !ok {
		return nil, &UnknownMessageType{}
	}
	return factory(), nil
}
