package database

import (
	"fmt"
	"log"

	"github.com/nkuros/invoices-api/initializers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB(config *initializers.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Connected Successfully to the Database")
}

func GetDatabase() *gorm.DB {
	return db
}