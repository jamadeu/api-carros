package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Connect to database
	db, err = ConnectDb()
	if err != nil {
		return fmt.Errorf("Error connecting db: %v", err)
	}
	return nil
}

func GetConnDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize logger
	logger = NewLogger(p)
	return logger
}
