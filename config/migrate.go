package config

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	mm "github.com/golang-migrate/migrate/v4/database/mysql"
	"go.uber.org/zap"
)

const migrationFolderName = "migrations"

// DoMigrateはDBのマイグレーションを実施する。
func DoMigrate(dsn string, isTest bool) error {
	m, err := connectToDB(dsn, isTest)
	if err != nil {
		return fmt.Errorf("DBマイグレーション用接続においてエラーが発生しました %v", err)
	}
	defer func() {
		err1, err2 := m.Close()
		if err1 != nil || err2 != nil {
			panic(fmt.Sprint("DBマイグレーション用接続のClose時にエラーが発生しました sourceErr=%w, databaseErr=%w", err1, err2))
		}
	}()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

// ResetMigrateはDBのマイグレーションを全てDOWN状態にする。
// 主目的はDBテスト用。
func ResetMigrate(dsn string, isTest bool) error {
	m, err := connectToDB(dsn, isTest)
	if err != nil {
		return fmt.Errorf("DBマイグレーション用接続においてエラーが発生しました %w", err)
	}
	defer func() {
		err1, err2 := m.Close()
		if err1 != nil || err2 != nil {
			panic(fmt.Sprint("DBマイグレーション用接続のClose時にエラーが発生しました sourceErr=%w, databaseErr=%w", err1, err2))
		}
	}()

	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

// connectToDBはマイグレーション用のDB接続を行う。
func connectToDB(dsn string, isTest bool) (*migrate.Migrate, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s&multiStatements=true", dsn))
	if err != nil {
		return nil, fmt.Errorf("DB接続(dsn=%s)においてエラーが発生しました: %w", dsn, err)
	}

	driver, err := mm.WithInstance(db, &mm.Config{})
	if err != nil {
		return nil, fmt.Errorf("マイグレーション用DB接続(dsn=%s)においてエラーが発生しました: %w", dsn, err)
	}

	fPath, err := getDBDefinitionFolderPath(isTest)
	if err != nil {
		return nil, fmt.Errorf("マイグレーション用フォルダパス(isTest=%v)取得においてエラーが発生しました: %w", isTest, err)
	}

	return migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", fPath),
		"mysql",
		driver,
	)
}

// getDBDefinitionFolderPathはマイグレーション用ファイル格納フォルダのパスを取得する。
// マイグレーションファイルはプロジェクトルート配下のmirationsフォルダを参照するが、
// テスト（go testによるユニットテスト）かどうかで現在位置のディレクトリが異なるので、引数で判断してパスを設定している。
func getDBDefinitionFolderPath(isTest bool) (string, error) {
	if isTest {
		_, testSourceFilePath, _, ok := runtime.Caller(0)
		if !ok {
			zap.L().Fatal("現在ディレクトリ取得処理でエラー")
			panic(fmt.Errorf("現在ディレクトリ取得処理でエラー"))
		}
		// プロジェクトルート/config/migrate.goというディレクトリ前提でマイグレーションフォルダのパスを取得している
		return filepath.Join(filepath.Dir(filepath.Dir(testSourceFilePath)), migrationFolderName), nil
	}

	return migrationFolderName, nil
}
