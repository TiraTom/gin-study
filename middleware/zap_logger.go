package middleware

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/Tiratom/gin-study/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ZapLogger struct {
}

func (l ZapLogger) Info(ctx context.Context, msg string) {
	zap.L().Info(msg, zap.Any(config.LOG_KEY_NAME_FOR_REQUEST_ID, ctx.Value(config.CONTEXT_KEY_FOR_REQUEST_ID)))
}

func (l ZapLogger) Warn(ctx context.Context, msg string) {
	zap.L().Warn(msg, zap.Any(config.LOG_KEY_NAME_FOR_REQUEST_ID, ctx.Value(config.CONTEXT_KEY_FOR_REQUEST_ID)))
}

func (l ZapLogger) Error(ctx context.Context, msg string) {
	zap.L().Error(msg, zap.Any(config.LOG_KEY_NAME_FOR_REQUEST_ID, ctx.Value(config.CONTEXT_KEY_FOR_REQUEST_ID)))
}

// GetZapLogger zapのロガーを取得する
func GetZapLogger() *zap.Logger {

	zapLogger, err := readZapConfig().Build()
	if err != nil {
		panic(err)
	}

	return zapLogger
}

// readZapConfig zap用の設定ファイルを読み込む
// <https://qiita.com/emonuh/items/28dbee9bf2fe51d28153>を参考にした
func readZapConfig() *zap.Config {
	configJson, err := ioutil.ReadFile("config/zap_config.json")
	if err != nil {
		panic(err)
	}

	var zapConfig zap.Config
	err = json.Unmarshal(configJson, &zapConfig)
	if err != nil {
		panic(err)
	}

	return &zapConfig
}

// GetZapLoggerInterceptor リクエスト詳細をログに出せるようにしたインターセプター。
// <https://note.com/dd_techblog/n/nd902b7ef8088>や<https://www.sambaiz.net/article/174/>を参考にした
func GetZapLoggerUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestID := config.CreateNewRequestID()

		newCtx := context.WithValue(ctx, config.CONTEXT_KEY_FOR_REQUEST_ID, requestID.Value)
		zap.L().Info("[REQUEST_START]", zap.String(config.LOG_KEY_NAME_FOR_REQUEST_ID, requestID.Value))

		resp, err := handler(newCtx, req)

		zap.L().Info("[REQUEST_END]", zap.String(config.LOG_KEY_NAME_FOR_REQUEST_ID, requestID.Value))

		return resp, err
	}
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{}
}
