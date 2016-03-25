package memberlist

import "log"

type Logger interface {
	Printf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
}

type _DefaultLogger struct {}

func defaultLogger() Logger {
	return &_DefaultLogger{}
}

func (l *_DefaultLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *_DefaultLogger) Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (l *_DefaultLogger) Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *_DefaultLogger) Warnf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (l *_DefaultLogger) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *_DefaultLogger) Fatalf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *_DefaultLogger) Panicf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
