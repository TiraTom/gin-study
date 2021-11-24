package config

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Gdb *gorm.DB
}

func NewDB(env *Environment) *DB {
	gdb, err := gorm.Open(mysql.Open(env.DB_DNS), getGormConf(env))
	if err != nil {
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

func getGormConf(env *Environment) *gorm.Config {
	var gormConf *gorm.Config
	if env.IsDebugEnv() {
		// デバッグ用のログレベル設定中。環境変数読み込みにできるといいかも
		gormConf = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	} else {
		gormConf = &gorm.Config{}
	}

	return gormConf
}
