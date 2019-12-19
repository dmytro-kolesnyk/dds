package logger

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	ownerName string
}

func NewMainLogger() *Logger {
	logger := Logger{""}

	return &logger
}

func NewLogger(owner interface{}) *Logger {
	logger := Logger{getTypeName(owner)}

	return &logger
}

func getTypeName(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (logger *Logger) Info(args ...interface{}) {
	log.Info(logger.ownerName, args)
}
