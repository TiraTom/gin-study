package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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
	DB_NAME                     string        `required:"true"`
	DB_CONNECTION_MAX_LIFE_TIME time.Duration `default:"10s"`
	DB_MAX_OPEN_CONNECTION      int           `default:"10"`
	DB_CONNECTION_MAX_IDLE_TIME time.Duration `default:"10s"`
	DB_DNS                      string        // envファイルには項目として不要。各項目から生成し設定される。
}

// ENVの値として設定可能な値
var acceptedENVValue = []string{"local", "test1", "test2"}

// 開発環境とみなすENVの値
var debugENVValue = []string{"local", "test1", "test2"}

// 設定ファイルで環境名を表す項目名
const keyForENV = "ENV"

// .env.xxxファイルが格納されているフォルダ名
const envFileFolder = "conf-files"

// IsDebugEnvは、ローカル開発環境や検証環境などの非本番環境の場合にtrueを返す
func (e *Environment) IsDebugEnv() bool {
	if isNotEnvSet() {
		panic("環境変数ENVがセットされていません")
	}
	if !isAcceptedEnvValue() {
		panic(fmt.Errorf("想定外の値(%v)が環境変数ENVにセットされています", os.Getenv(keyForENV)))
	}

	env := os.Getenv(keyForENV)
	for _, v := range debugENVValue {
		if v == env {
			return true
		}
	}
	return false
}

// IsTestEnvは、テスト環境の場合にtrueを返却する
func IsTestEnv() (bool, error) {
	if isNotEnvSet() {
		return false, fmt.Errorf("環境変数ENVがセットされていません")
	}
	if !isAcceptedEnvValue() {
		return false, fmt.Errorf("想定外の値(%v)が環境変数ENVにセットされています", os.Getenv(keyForENV))
	}
	return strings.Contains(os.Getenv(keyForENV), "test"), nil
}

// isNotEnvSetは、環境変数ENVに値が設定されている場合にtrueを返却する
func isNotEnvSet() bool {
	return os.Getenv(keyForENV) == ""
}

// isAcceptedEnvValueは、ENVの値として許容されている値が設定されている場合にtrueを返却する
func isAcceptedEnvValue() bool {
	env := os.Getenv(keyForENV)
	for _, v := range acceptedENVValue {
		if v == env {
			return true
		}
	}
	return false
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
	err = godotenv.Load(fmt.Sprintf("%s/%s/.env.%s", projectRoot, envFileFolder, os.Getenv("ENV")))
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
	env.DB_DNS = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", env.DB_USER, env.DB_PASSWORD, env.DB_ADDRESS, env.DB_NAME)

	return &env
}
