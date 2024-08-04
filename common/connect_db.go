package common

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Cannot get env")
	}
	// Corrected DSN

	// dsn := "root:root@tcp(192.168.145.57:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("MYSQL_CONN_STRING")
	if dsn == "" {
		return nil, errors.New("MYSQL_CONN_STRING not set in environment variables")
	}

	// Set up a timeout context for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var db *gorm.DB
	for retries := 3; retries > 0; retries-- {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			// Check if the connection is valid
			sqlDB, err := db.DB()
			if err != nil {
				return nil, err
			}

			err = sqlDB.PingContext(ctx)
			if err == nil {
				log.Println("Connected to the database successfully")
				return db, nil
			}
		}

		log.Printf("Failed to connect to the database, retries left: %d, error: %s", retries-1, err)
		time.Sleep(2 * time.Second) // Wait for 2 seconds before retrying
	}

	return nil, errors.New("Failed to connect to the database after multiple attempts")
}
