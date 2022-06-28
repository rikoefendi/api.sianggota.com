package database

import (
	"fmt"

	"api.sianggota.com/config"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Connect(cfg config.Database) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s application_name=%s", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode, cfg.TimeZone, cfg.Name)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		log.Error("Failed To Connect Database", err.Error())
		panic(err)
	}
	if cfg.Logger {
		// db = db.Debug()
		db.Logger.LogMode(logger.Error)
	}
	return db
}

func Session() *gorm.DB {
	return db
}
