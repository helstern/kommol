package app

type Logger interface {
	Debug(msg string)
	Debugf(msg string, v ...interface{})
	Error(msg string)
	Errorf(msg string, v ...interface{})
	Fatal(msg string)
	Fatalf(msg string, v ...interface{})
	Info(msg string)
	Infof(msg string, v ...interface{})
	Trace(msg string) Logger
	Warn(msg string)
	Warnf(msg string, v ...interface{})
	WithError(err error) Logger
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
}

type LoggerFactory func(fields []Fields) Logger
