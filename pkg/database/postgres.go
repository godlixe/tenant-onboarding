package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseCredentials struct {
	Host     string
	User     string
	Password string
	Port     string
	DBName   string
}

func NewPostgresClient(
	creds DatabaseCredentials,
) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		creds.Host,
		creds.User,
		creds.Password,
		creds.DBName,
		creds.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
