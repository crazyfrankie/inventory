package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"inventory/model"
	"log"
)

var (
	DB *gorm.DB
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bubble"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return DB, nil
}

func InitModel() {
	if err := DB.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
