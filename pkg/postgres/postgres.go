package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes a connection to the PostgreSQL database using a URL.
func InitDB(dbURL string) (*gorm.DB, error) {
	// Open the database connection using the URL
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Return the DB instance
	return db, nil
}
