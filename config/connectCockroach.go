package config

import (
	"fmt"
	"myapp/logger"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectCockroach() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbDatabase, dbPort),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), logger.InitConfig())

	if err != nil {
		panic(err)
	}

	return db
}
