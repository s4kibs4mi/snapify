package log

import (
	"github.com/s4kibs4mi/snapify/log/hooks"
	"github.com/sirupsen/logrus"
	"os"
)

type IAppLogger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Fatal(args ...interface{})
}

type AppLogger struct {
	logger *logrus.Logger
}

func (l *AppLogger) Info(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *AppLogger) Error(args ...interface{}) {
	l.logger.Errorln(args...)
}

func (l *AppLogger) Warning(args ...interface{}) {
	l.logger.Warningln(args...)
}

func (l *AppLogger) Fatal(args ...interface{}) {
	l.logger.Fatalln(args...)
}

func New() IAppLogger {
	defLogger := logrus.New()
	defLogger.Out = os.Stdout
	defLogger.AddHook(hooks.NewHook())

	return &AppLogger{
		logger: defLogger,
	}
}
