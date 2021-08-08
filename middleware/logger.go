package middleware

import (
	"go.uber.org/zap"
)

func GetZapLogger() *zap.Logger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return zapLogger
}
