package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
)

const (
	host      = "localhost"
	user      = "aj_e"
	password  = "eng464999"
	dbname    = "bugs_reports"
	port      = "5432"
	sslmode   = "disable"
	TiemeZone = "Africa/Cairo"
)

func ConnectDB() (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, TiemeZone)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		QueryFields:     true,
	})
	
	if err != nil {
		return nil, err
	}

	return db, nil	
}
