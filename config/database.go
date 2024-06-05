package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Password string
	User     string
	Name     string
	Port     string
}

func SetupDatabase(cfg *DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
