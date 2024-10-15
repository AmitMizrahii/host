package dal

import (
	"fmt"
	"gorm.io/gorm"
	"log"

	"gorm.io/driver/postgres"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func InitDB(config DBConfig) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	db.AutoMigrate(&UserModel{})

	log.Println("Successfully connected to the database")
	return db, nil
}