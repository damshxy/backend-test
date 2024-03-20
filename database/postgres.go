package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsn := "host=localhost user=postgres password=root dbname=db_manajemen-backend-test port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}