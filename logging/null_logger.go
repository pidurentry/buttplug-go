package logging

type NullLogger struct{}

func (*NullLogger) Trace(interface{}) {}

func (*NullLogger) Tracef(string, ...interface{}) {}

func (*NullLogger) Debug(interface{}) {}

func (*NullLogger) Debugf(string, ...interface{}) {}

func (*NullLogger) Info(interface{}) {}

func (*NullLogger) Infof(string, ...interface{}) {}

func (*NullLogger) Warning(interface{}) {}

func (*NullLogger) Warningf(string, ...interface{}) {}

func (*NullLogger) Error(interface{}) {}

func (*NullLogger) Errorf(string, ...interface{}) {}

func (*NullLogger) Fatal(interface{}) {}

func (*NullLogger) Fatalf(string, ...interface{}) {}
