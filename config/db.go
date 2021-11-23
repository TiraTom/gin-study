package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mm "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Gdb *gorm.DB
}

func NewDB(env *Environment) *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/gin_study?parseTime=true", env.DB_USER, env.DB_PASSWORD, env.DB_ADDRESS)
	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// デバッグ用のログレベル設定中。環境変数読み込みにできるといいかも
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// マイグレーション設定

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/gin_study?parseTime=true&multiStatements=true", env.DB_USER, env.DB_PASSWORD, env.DB_ADDRESS))
	if err != nil {
		panic(err)
	}

	driver, err := mm.WithInstance(db, &mm.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/definitions",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	// コネクションプールの設定

	sqlDB, err := gdb.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxLifetime(env.DB_CONNECTION_MAX_LIFE_TIME)
	sqlDB.SetMaxOpenConns(env.DB_MAX_OPEN_CONNECTION)
	sqlDB.SetConnMaxIdleTime(env.DB_CONNECTION_MAX_IDLE_TIME)

	return &DB{gdb}
}
