package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

var Default Logger

type (
	Logger interface {
		Info(xRequestID string, i ...interface{})
		Debug(xRequestID string, i ...interface{})
		Warn(xRequestID string, i ...interface{})
		Error(xRequestID string, i ...interface{})
	}
	loggerImpl struct {
		appName string
	}
)

func NewLogger(appName string) {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	Default = &loggerImpl{
		appName: appName,
	}
}

func (l loggerImpl) Info(xRequestID string, i ...interface{}) {
	position := loggerCaller()
	logrus.WithFields(logrus.Fields{
		"APP_NAME":     l.appName,
		"X_REQUEST_ID": xRequestID,
		"POSITION":     position,
	}).Info(i...)
}

func (l loggerImpl) Debug(xRequestID string, i ...interface{}) {
	position := loggerCaller()
	logrus.WithFields(logrus.Fields{
		"APP_NAME":     l.appName,
		"X_REQUEST_ID": xRequestID,
		"POSITION":     position,
	}).Debug(i...)
}

func (l loggerImpl) Warn(xRequestID string, i ...interface{}) {
	position := loggerCaller()
	logrus.WithFields(logrus.Fields{
		"APP_NAME":     l.appName,
		"X_REQUEST_ID": xRequestID,
		"POSITION":     position,
	}).Warn(i...)
}

func (l loggerImpl) Error(xRequestID string, i ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"APP_NAME":     l.appName,
		"X_REQUEST_ID": xRequestID,
	}).Error(i...)
}

func loggerCaller() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d", file, line)
}
