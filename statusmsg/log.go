package statusmsg

import "github.com/pidurentry/buttplug-go"

type LogLevel string
type LogMessage string

const (
	LOG_LEVEL_OFF   LogLevel = "Off"
	LOG_LEVEL_FATAL LogLevel = "Fatal"
	LOG_LEVEL_ERROR LogLevel = "Error"
	LOG_LEVEL_WARN  LogLevel = "Warn"
	LOG_LEVEL_INFO  LogLevel = "Info"
	LOG_LEVEL_DEBUG LogLevel = "Debug"
	LOG_LEVEL_TRACE LogLevel = "Trace"
)

type Log struct {
	id         buttplug.MessageId
	logLevel   LogLevel
	logMessage LogMessage
}

func (log *Log) Id() buttplug.MessageId {
	return log.id
}

func (log *Log) LogLevel() LogLevel {
	return log.logLevel
}

func (log *Log) LogMessage() LogMessage {
	return log.logMessage
}
