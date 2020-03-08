package logging

import (
	"github.com/apex/log"
	logging "github.com/helstern/kommol/internal/core/logging/app"
)

type ApexLoggerAdapter struct {
	logging.Logger
	entry *log.Entry
}

func (l ApexLoggerAdapter) Debug(msg string) {
	l.entry.Debug(msg)
}

func (l ApexLoggerAdapter) Debugf(msg string, v ...interface{}) {
	l.entry.Debugf(msg, v...)
}

func (l ApexLoggerAdapter) Error(msg string) {
	l.entry.Error(msg)
}

func (l ApexLoggerAdapter) Errorf(msg string, v ...interface{}) {
	l.entry.Errorf(msg, v...)
}

func (l ApexLoggerAdapter) Fatal(msg string) {
	l.entry.Fatal(msg)
}

func (l ApexLoggerAdapter) Fatalf(msg string, v ...interface{}) {
	l.entry.Fatalf(msg, v...)
}

func (l ApexLoggerAdapter) Info(msg string) {
	l.entry.Info(msg)
}

func (l ApexLoggerAdapter) Infof(msg string, v ...interface{}) {
	l.entry.Infof(msg, v...)
}

func (l *ApexLoggerAdapter) Trace(msg string) logging.Logger {
	entry := l.entry.Trace(msg)
	return &ApexLoggerAdapter{
		entry: entry,
	}
}

func (l ApexLoggerAdapter) Warn(msg string) {
	l.entry.Warn(msg)
}

func (l ApexLoggerAdapter) Warnf(msg string, v ...interface{}) {

}

func (l ApexLoggerAdapter) WithError(err error) logging.Logger {
	entry := l.entry.WithError(err)
	return &ApexLoggerAdapter{
		entry: entry,
	}
}

func (l ApexLoggerAdapter) WithField(key string, value interface{}) logging.Logger {
	entry := l.entry.WithField(key, value)
	return &ApexLoggerAdapter{
		entry: entry,
	}
}

func (l ApexLoggerAdapter) WithFields(fields logging.Fields) logging.Logger {
	entry := l.entry.WithFields(log.Fields(fields))
	return &ApexLoggerAdapter{
		entry: entry,
	}
}

func NewLogger(entry *log.Entry) logging.Logger {
	return &ApexLoggerAdapter{
		entry: entry,
	}
}

func NewLoggerWithFields(fields []logging.Fields) logging.Logger {
	entry := log.WithFields(log.Fields{})
	for _, item := range fields {
		entry = entry.WithFields(log.Fields(item))
	}
	return NewLogger(entry)
}
