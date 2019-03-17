package statusmsg

//go:generate go run ../message/generator.go -- $GOFILE

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
	ID         int        `json:"Id"`
	LogLevel   LogLevel   `json:"LogLevel"`
	LogMessage LogMessage `json:"LogMessage"`
}
