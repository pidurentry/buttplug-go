package statusmsg

import "github.com/pidurentry/buttplug-go"

type Ping struct {
	id buttplug.MessageId
}

func (ping *Ping) Id() buttplug.MessageId {
	return ping.id
}
