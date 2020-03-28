package logger

import "go.uber.org/zap"

func NewDefault() *zap.Logger {
	return zap.NewExample()
}
