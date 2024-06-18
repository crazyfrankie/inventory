package config

import (
	"log"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"inventory/model"
)

var (
	DB       *gorm.DB
	DbConfig *DatabaseConfig
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func LoadConfig(file string) (*DatabaseConfig, error) {
	cfg, err := ini.Load(file)
	if err != nil {
		return nil, err
	}

	dbConfig := &DatabaseConfig{
		User:     cfg.Section("database").Key("user").String(),
		Password: cfg.Section("database").Key("password").String(),
		Host:     cfg.Section("database").Key("host").String(),
		Port:     cfg.Section("database").Key("port").String(),
		Name:     cfg.Section("database").Key("name").String(),
	}

	return dbConfig, nil
}

func InitDB(dbConfig *DatabaseConfig) (*gorm.DB, error) {
	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Name
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
