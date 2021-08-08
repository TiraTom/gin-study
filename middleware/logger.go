package middleware

import (
	"log"

	"go.uber.org/zap"
)

func GetZapLogger() *zap.Logger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	return zapLogger
}
