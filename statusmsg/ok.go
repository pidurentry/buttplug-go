package statusmsg

import "github.com/pidurentry/buttplug-go"

type Ok struct {
	id buttplug.MessageId
}

func (ok *Ok) Id() buttplug.MessageId {
	return ok.id
}
