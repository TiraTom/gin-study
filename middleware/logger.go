package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	CONTEXT_KEY_FOR_REQUEST_ID = "RequestID"
)

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
func GetZapLoggerUnaryInterceptor(zapLogger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var requestID = ""
		uuid, err := uuid.NewRandom()

		if err != nil {
			requestID = "cannot create requestID"
		} else {
			requestID = uuid.String()
		}

		newCtx := context.WithValue(ctx, CONTEXT_KEY_FOR_REQUEST_ID, requestID)
		zapLogger.Info(fmt.Sprintf("[REQUEST_START] RequestID=%s", requestID))

		resp, err := handler(newCtx, req)

		zapLogger.Info(fmt.Sprintf("[REQUEST_END] RequestID=%s", requestID))

		return resp, err
	}
}
