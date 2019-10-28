package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/s4kibs4mi/snapify/config"
	"github.com/s4kibs4mi/snapify/log"
)

var instance *gorm.DB

func ConnectSQLDB() error {
	db, err := gorm.Open("postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
			config.DB().Username, config.DB().Password,
			config.DB().Host, config.DB().Port, config.DB().Name))
	if err != nil {
		return err
	}

	db.DB().SetMaxIdleConns(config.DB().MaxIdleConnections)
	db.DB().SetMaxOpenConns(config.DB().MaxOpenConnections)
	db.DB().SetConnMaxLifetime(config.DB().MaxConnectionLifetime)

	db.LogMode(true)
	db.SetLogger(log.Log())

	instance = db
	return nil
}

func DB() *gorm.DB {
	return instance
}
