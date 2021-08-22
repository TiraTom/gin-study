package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Gdb *gorm.DB
}

func NewDB(env *Environment) *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/gin_study", env.DB_USER, env.DB_PASSWORD, env.DB_ADDRESS)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
