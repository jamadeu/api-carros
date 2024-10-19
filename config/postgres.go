package config

import (
	"github.com/jamadeu/api-cars/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	logger := GetLogger("InitializeDb")
	// Connect DB
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error to connect database: %v", err)
		return nil, err
	}

	// Migrate the schema
	if err = db.AutoMigrate(&schemas.Car{}); err != nil {
		logger.Errorf("Automigratoin error: %v", err)
		return nil, err
	}
	return db, nil
}
