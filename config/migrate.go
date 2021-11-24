package config

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	mm "github.com/golang-migrate/migrate/v4/database/mysql"
)

const dbDefinitionFolder = "migrations/definitions"
const dbDummyDataFolder = "migrations/dummyData"

// DoMigrateはDBのマイグレーションを実施する
// マイグレーションはプロジェクトルート配下のmirations
func DoMigrate(dsn string, needDummyData bool) error {
	err := migrateDifinitions(dsn)
	if err != nil {
		return err
	}

	if needDummyData {
		err = insertDummyData(dsn)
		if err != nil {
			return err
		}
	}

	return nil
}

func migrateDifinitions(dsn string) error {
	return doMigrate(dsn, fmt.Sprint("file://", dbDefinitionFolder))
}

func insertDummyData(dsn string) error {
	// TODO これだと、definitionsフォルダ配下で置いてるバージョンのファイルが見つからなくてエラーという結果になる
	// TODO 別の方法を考える必要がある
	// return doMigrate(dsn, fmt.Sprint("file://", dbDummyDataFolder))
	return fmt.Errorf("not yet implemented for inserting data from %s", dbDummyDataFolder)
}

func doMigrate(dsn string, mFolderPath string) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s&multiStatements=true", dsn))
	if err != nil {
		return err
	}

	driver, err := mm.WithInstance(db, &mm.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		mFolderPath,
		"mysql",
		driver,
	)
	if err != nil {
		return err
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
