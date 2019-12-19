package logger

import (
	"reflect"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	ownerName string
}

func NewMainLogger() *Logger {
	return NewLogger("")
}

func NewLogger(owner interface{}) *Logger {
	logger := Logger{getTypeName(owner)}

	log.SetLevel(log.InfoLevel) //TODO update to read from config.yaml

	return &logger
}

func getTypeName(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (logger *Logger) Trace(args ...interface{}) {
	log.Trace(logger.ownerName, args)
}

func (logger *Logger) Debug(args ...interface{}) {
	log.Debug(logger.ownerName, args)
}

func (logger *Logger) Info(args ...interface{}) {
	log.Info(logger.ownerName, args)
}

func (logger *Logger) Warn(args ...interface{}) {
	log.Warn(logger.ownerName, args)
}

func (logger *Logger) Error(args ...interface{}) {
	log.Error(logger.ownerName, args)
}

func (logger *Logger) Fatal(args ...interface{}) {
	log.Fatal(logger.ownerName, args)
}

func (logger *Logger) Panic(args ...interface{}) {
	log.Panic(logger.ownerName, args)
}
