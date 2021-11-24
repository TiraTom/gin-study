package config

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	mm "github.com/golang-migrate/migrate/v4/database/mysql"
)

const dbDefinitionFolder = "migrations"

// DoMigrateはDBのマイグレーションを実施する
// マイグレーションファイルはプロジェクトルート配下のmirationsフォルダを参照する
func DoMigrate(dsn string) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s&multiStatements=true", dsn))
	if err != nil {
		return err
	}

	driver, err := mm.WithInstance(db, &mm.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprint("file://", dbDefinitionFolder),
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
