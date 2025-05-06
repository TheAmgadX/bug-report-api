package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
	sslmode  string
	timeZone string
}

func initDBConfig() *dbConfig {
	_ = godotenv.Load(".env")

	return &dbConfig{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
		sslmode:  os.Getenv("DB_SSLMODE"),
		timeZone: os.Getenv("APP_TIMEZONE"),
	}
}

func ConnectDB() (*gorm.DB, error) {
	cnf := initDBConfig()
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cnf.host, cnf.user, cnf.password, cnf.dbname, cnf.port, cnf.sslmode, cnf.timeZone)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		QueryFields: true,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

