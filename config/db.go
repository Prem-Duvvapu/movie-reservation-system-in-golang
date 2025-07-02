package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"movie-reservation-system/models"
)

func NewDB(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Enable UUID extension
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return nil, fmt.Errorf("failed to create UUID extension: %w", err)
	}

	// Auto-migrate all models
	if err := db.AutoMigrate(
				&models.Movie{},
				&models.User{},
				&models.Showtime{},
				&models.Reservation{},
	); err != nil {
		return nil, fmt.Errorf("Failed to auto-migrate: %w", err)
	}

	return db, nil
}