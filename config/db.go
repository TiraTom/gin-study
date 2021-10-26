package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Gdb *gorm.DB
}

func NewDB(env *Environment) *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/gin_study?parseTime=true", env.DB_USER, env.DB_PASSWORD, env.DB_ADDRESS)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// デバッグ用のログレベル設定中。環境変数読み込みにできるといいかも
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// コネクションプールの設定
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxLifetime(env.DB_CONNECTION_MAX_LIFE_TIME)
	sqlDB.SetMaxOpenConns(env.DB_MAX_OPEN_CONNECTION)
	sqlDB.SetConnMaxIdleTime(env.DB_CONNECTION_MAX_IDLE_TIME)

	return &DB{db}
}
