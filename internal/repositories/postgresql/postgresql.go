package postgresql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// New creates a new GORM instance connected to a PostgreSQL database, with New Relic integration.
func New(config Config) (*gorm.DB, error) {
	// Configure the DSN for PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Host,
		config.Username,
		config.Password,
		config.DBName,
		config.Port,
	)

	// Initialize the GORM DB connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Adjust the log mode as needed
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
