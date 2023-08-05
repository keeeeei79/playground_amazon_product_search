package logging

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger = NewLogger()

func NewLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}