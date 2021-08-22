package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	APP_PORT_NUM                string        `default:"8081"`
	DB_USER                     string        `required:"true"`
	DB_PASSWORD                 string        `required:"true"`
	DB_ADDRESS                  string        `required:"true"`
	DB_CONNECTION_MAX_LIFE_TIME time.Duration `default:"10s"`
	DB_MAX_OPEN_CONNECTION      int           `default:"10"`
	DB_CONNECTION_MAX_IDLE_TIME time.Duration `default:"10s"`
}

// SetEnvValues 環境変数を設定ファイルから取得しその後変数として保持する。アプリ起動時に呼び出す処理のため、環境変数取得時にエラーが発生した場合はpanicを起こすようにしている。
func NewEnvironment() *Environment {
	// 指定した環境に対応した環境変数ファイルを読み込む
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("ENV")))
	if err != nil {
		panic(err)
	}

	var env Environment
	err = envconfig.Process("", &env)
	if err != nil {
		panic(err)
	}

	return &env
}
