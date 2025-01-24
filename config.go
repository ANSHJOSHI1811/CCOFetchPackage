package config

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres" // PostgreSQL driver for GORM
	"gorm.io/gorm"            // Core GORM package
)

// DB is the global database connection object
var DB *gorm.DB

// InitializeDatabase initializes the database connection and automigrates models
func InitializeDatabase() {
	dsn := "host=localhost user=postgres password=password dbname=cco port=5432 sslmode=disable" // Database connection string
	var err error

	// Open the database connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Database connected successfully!")

	// Automigrate models
	err = autoMigrateModels(DB)
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	fmt.Println("Database migration completed successfully!")
}

// autoMigrateModels handles automigration of all models
func autoMigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Provider{},
		&models.Service{},
		&models.Region{},
		&models.SKU{},
		&models.Price{},
		&models.Term{},
		&models.SavingPlan{},
	)
}
