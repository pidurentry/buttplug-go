package logging

type Logger interface {
	Trace(interface{})
	Tracef(string, ...interface{})
	Debug(interface{})
	Debugf(string, ...interface{})
	Info(interface{})
	Infof(string, ...interface{})
	Warning(interface{})
	Warningf(string, ...interface{})
	Error(interface{})
	Errorf(string, ...interface{})
	Fatal(interface{})
	Fatalf(string, ...interface{})
}
