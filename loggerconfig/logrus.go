package loggerconfig

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

var Info = func(args ...interface{}) {
	InfoImpl(args...)
}

func InfoImpl(args ...interface{}) {
	message := buildMessage(args...)
	logger.Info(message)
}

var Warn = func(args ...interface{}) {
	WarnImpl(args...)
}

func WarnImpl(args ...interface{}) {
	message := buildMessage(args...)
	logger.Warn(message)
}

var Panic = func(args ...interface{}) {
	PanicImpl(args...)
}

func PanicImpl(args ...interface{}) {
	message := buildMessage(args...)
	logger.Panic(message)
}

var Error = func(args ...interface{}) {
	ErrorImpl(args...)
}

func ErrorImpl(args ...interface{}) {
	message := buildMessage(args...)

	logrus.Error(message)
}

func buildMessage(args ...interface{}) string {
	var message string
	for i, arg := range args {
		if i == 0 {
			message = fmt.Sprint(arg)
		} else {
			message += " " + fmt.Sprint(arg)
		}
	}
	return message
}

func InitLogrus () {
	logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)
}