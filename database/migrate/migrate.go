package main

import (
	"fmt"
	"log"

	"github.com/nkuros/invoices-api/models"
	"github.com/nkuros/invoices-api/database"
	"github.com/nkuros/invoices-api/initializers"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	database.ConnectDB(&config)
}

func main() {
	db := database.GetDatabase()
	db.AutoMigrate(
		&models.Invoice{},
		&models.User{},
		&models.Customer{},
		&models.BankAccount{},
		&models.Company{},
	)
	fmt.Println("Migration complete")
}

