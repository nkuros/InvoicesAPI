package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nkuros/invoices-api/controllers"
	"github.com/nkuros/invoices-api/database"
)



func InvoiceRouter(router *gin.RouterGroup) *gin.RouterGroup {
	{
		books := router.Group("invoices")
		{
			db := database.GetDatabase()
			invoiceController := controllers.NewInvoiceController(db)
			books.GET("/", invoiceController.GetAllInvoicesFromPeriod)
			books.POST("/", invoiceController.CreateInvoice)
		}
	}

	return router
}