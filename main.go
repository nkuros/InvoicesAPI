package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/nkuros/invoices-api/database"
	"github.com/nkuros/invoices-api/initializers"
	"github.com/nkuros/invoices-api/routers"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	database.ConnectDB(&config)

}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

    server := gin.Default()

	router := server.Group("/api")
	routers.InvoiceRouter(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}


