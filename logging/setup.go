package logging

var logger Logger

func init() {
	logger = &NullLogger{}
}

func SetLogger(newLogger Logger) {
	logger = newLogger
}

func GetLogger() Logger {
	return logger
}
