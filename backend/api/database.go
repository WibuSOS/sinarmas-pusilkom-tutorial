package api

import (
	"fmt"
	"github.com/WibuSOS/sinarmas/models"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func SetupDb() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	dbUrl := os.Getenv("DATABASE_URL")

	if os.Getenv("ENVIRONMENT") == "PROD" {
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := db.AutoMigrate(&models.Todos{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, err
}
