package api

import "go.uber.org/zap"

type TgLogger struct {
	logger *zap.Logger
}

func (l *TgLogger) Println(v ...interface{}) {
	l.logger.Sugar().Info(v...)
}

func (l *TgLogger) Printf(format string, v ...interface{}) {
	l.logger.Sugar().Infof(format, v...)
}
