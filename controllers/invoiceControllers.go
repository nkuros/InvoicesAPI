package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nkuros/invoices-api/constants"
	"github.com/nkuros/invoices-api/models"
	"gorm.io/gorm"
)


type InvoiceController struct {
	Database *gorm.DB
}

func NewInvoiceController(db *gorm.DB) *InvoiceController {
	return &InvoiceController{Database: db}
}


func (ic InvoiceController) GetAllInvoicesFromPeriod(c *gin.Context) {
	return
}


func (ic InvoiceController) CreateInvoice(c *gin.Context) {
	return
}


