package statusmsg

import "github.com/pidurentry/buttplug-go"

type RequestLog struct {
	id       buttplug.MessageId
	logLevel LogLevel
}

func (requestLog *RequestLog) Id() buttplug.MessageId {
	return requestLog.id
}

func (requestLog *RequestLog) LogLevel() LogLevel {
	return requestLog.logLevel
}
