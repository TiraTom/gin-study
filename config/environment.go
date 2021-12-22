package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Environment struct {
	ENV                         string        `required:"true"`
	APP_PORT_NUM                string        `default:"8081"`
	DB_USER                     string        `required:"true"`
	DB_PASSWORD                 string        `required:"true"`
	DB_ADDRESS                  string        `required:"true"`
	DB_CONNECTION_MAX_LIFE_TIME time.Duration `default:"10s"`
	DB_MAX_OPEN_CONNECTION      int           `default:"10"`
	DB_CONNECTION_MAX_IDLE_TIME time.Duration `default:"10s"`
	DB_DNS                      string        // envファイルには項目として不要。各項目から生成し設定される。
}

// IsDebugEnvは、ローカル開発環境や検証環境などの非本番環境の場合にtrueを返す
func (e *Environment) IsDebugEnv() bool {
	if e.ENV == "local" || e.ENV == "test" {
		return true
	}

	panic(fmt.Errorf("想定しないENVの値が設定されています"))
}

func IsTestEnv() (bool, error) {
	if isNotEnvSet() {
		return false, fmt.Errorf("環境変数ENVがセットされていません")
	}
	return os.Getenv("ENV") == "test", nil
}

func isNotEnvSet() bool {
	return os.Getenv("ENV") == ""
}

// SetEnvValues 環境変数を設定ファイルから取得しその後変数として保持する。アプリ起動時に呼び出す処理のため、環境変数取得時にエラーが発生した場合はpanicを起こすようにしている。
func NewEnvironment() *Environment {
	projectRoot := "."

	isTest, err := IsTestEnv()
	if err != nil {
		panic(fmt.Errorf("テスト環境かどうかの判別処理でエラーが発生しました; %w", err))
	}

	if isTest {
		_, testSourceFilePath, _, ok := runtime.Caller(0)
		if !ok {
			zap.L().Fatal("現在ディレクトリ取得処理でエラー")
			panic(fmt.Errorf("現在ディレクトリ取得処理でエラー"))
		}
		projectRoot = filepath.Dir(filepath.Dir(testSourceFilePath))
	}

	// 指定した環境に対応した環境変数ファイルを読み込む
	err = godotenv.Load(fmt.Sprintf("%s/.env.%s", projectRoot, os.Getenv("ENV")))
	if err != nil {
		zap.L().Fatal(fmt.Sprint("環境変数ファイル読み込みでエラー; ", err))
		panic(err)
	}

	var env Environment
	err = envconfig.Process("", &env)
	if err != nil {
		zap.L().Fatal(err.Error())
		panic(err)
	}

	// 各項目から生成できるのでenvファイルの項目としては用意せずここで設定する
	env.DB_DNS = fmt.Sprintf("%s:%s@tcp(%s)/gin_study?parseTime=true", env.DB_USER, env.DB_PASSWORD, env.DB_ADDRESS)

	return &env
}
