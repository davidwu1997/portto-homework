package log

import (
	"go.uber.org/zap"
)

var (
	global *Logger
)

type Logger struct {
	*zap.SugaredLogger
}

func New() *Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

func Global() *Logger {
	if global == nil {
		global = New()
	}

	return global
}
