package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	env *Environment
}

func (d *DB) GormConnect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/gin_study", d.env.DB_USER, d.env.DB_PASSWORD, d.env.DB_ADDRESS)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// コネクションプールの設定
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxLifetime(d.env.DB_CONNECTION_MAX_LIFE_TIME)
	sqlDB.SetMaxOpenConns(d.env.DB_MAX_OPEN_CONNECTION)
	sqlDB.SetConnMaxIdleTime(d.env.DB_CONNECTION_MAX_IDLE_TIME)

	return db
}

func NewDB(env *Environment) *DB {
	return &DB{env}
}
