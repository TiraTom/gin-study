package usecase_impl_test

import (
	"os"
	"testing"

	"github.com/Tiratom/gin-study/config"
)

// SetUpForUsecaseDBTestはDB接続テストをするに当たっての前提となる共通処理を実施する。
// テスト記述部分をシンプルにするため、このメソッド内でエラーが起きた場合はそこでテストを失敗させている。
func SetUpForUsecaseDBTest(t *testing.T) (*config.Environment, *config.DB) {
	// infrastructureのrepository_implのテスト（DBと接続する）と競合しないようにするため環境はtest2としている
	err := os.Setenv("ENV", "test2")
	if err != nil {
		t.Errorf("テストの共通前処理でエラー発生: %v", err)
		t.FailNow()
	}

	conf := config.NewEnvironment()

	return conf, config.NewDB(conf)
}

// BeforeEachForDBTestはテスト用DBを初期状態にする。
// テスト記述部分をシンプルにするため、このメソッド内でエラーが起きた場合はそこでテストを失敗させている。
func BeforeEachForDBTest(t *testing.T, conf *config.Environment, db *config.DB) {
	err := config.ResetMigrate(conf.DB_DNS)
	if err != nil {
		t.Errorf("DBのリセット処理においてエラーが発生しました: %v", err)
		t.FailNow()
	}

	err = config.DoMigrate(conf.DB_DNS)
	if err != nil {
		t.Errorf("DBのマイグレーション処理においてエラーが発生しました: %v", err)
		t.FailNow()
	}
}
