package buttplug

type MessageId int

type Message interface {
	Id() MessageId
}
