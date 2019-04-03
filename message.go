package buttplug

import "github.com/pidurentry/buttplug-go/message"

type Message interface {
	Id() message.Id
	Serilize() interface{}
}
