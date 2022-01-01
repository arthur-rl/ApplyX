package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// postgres_uri := os.Getenv("POSTGRES_URI")
	dsn := "host=localhost user=postgres password=pass dbname=applyx port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connectiong to database")
		fmt.Println(err)
		os.Exit(1)
	}
	return db
}
