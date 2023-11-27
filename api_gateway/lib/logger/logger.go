package logger

import "go.uber.org/zap"

type Logger struct {
	log *zap.SugaredLogger
}

func NewLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	return logger.Sugar()
}
