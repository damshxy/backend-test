package migrations

import (
	"fmt"

	"github.com/damshxy/manajemen-backend-test/database"
	"github.com/damshxy/manajemen-backend-test/models"
)

func RunMigration() {
	err := database.DB.AutoMigrate(
		&models.Product{},
		&models.Order{},
		&models.User{},
	)
	
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Migrated")
}