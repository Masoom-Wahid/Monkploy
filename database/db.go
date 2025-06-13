package database

import (
	"fmt"
	"platform/pkg/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	host := env.GetEnv("DB_HOST", "localhost")
	port := env.GetEnv("DB_PORT", "5432")
	user := env.GetEnv("DB_USER", "postgres")
	password := env.GetEnv("DB_PASSWORD", "postgres")
	databaseName := env.GetEnv("DB_NAME", "postgres")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, databaseName)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil
	}

	DB = db.Debug()

	return DB
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = Connect()
	}

	return DB
}
