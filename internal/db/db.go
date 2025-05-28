package database

import (
	logger "github.com/MdSadiqMd/mail.send/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(dataSourceName string) {
	db := logger.New("database")

	var err error
	DB, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		db.Fatal("error in db connection: %v", err)
	}
	db.Info("Database connected successfully")

	err = DB.AutoMigrate()
	if err != nil {
		db.Fatal("error in db migration: %v", err)
	}
	db.Info("Database migrated successfully")
}

func GetDB() *gorm.DB {
	return DB
}
