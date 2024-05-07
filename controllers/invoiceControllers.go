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
	var p []models.Invoice

	
	fromDate := c.GetHeader("from_date")
	toDate := c.GetHeader("to_date")
	switch {
	case fromDate != "" && toDate == "":
		err := ic.Database.Where("due_date >= ?", fromDate).Find(&p).Error
		if err != nil {
			c.JSON(400, gin.H{
				"Error": "Query failed: " + err.Error(),
			})
			return
		}
	case fromDate == "" && toDate != "":
		err := ic.Database.Where("due_date <= ?", toDate).Find(&p).Error
		if err != nil {
			c.JSON(400, gin.H{
				"Error": "Query failed: " + err.Error(),
			})
			return
		}
	case fromDate != "" && toDate != "":
		err := ic.Database.Where("due_date BETWEEN ? AND ?", fromDate, toDate).Find(&p).Error
		if err != nil {
			c.JSON(400, gin.H{
				"Error": "Query failed: " + err.Error(),
			})
			return
		}
	default:
		err := ic.Database.Find(&p).Error
		if err != nil {
			c.JSON(400, gin.H{
				"Error": "Query failed: " + err.Error(),
			})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, p)




}


func (ic InvoiceController) CreateInvoice(c *gin.Context) {
	
	var p models.Invoice

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	type Invoice struct {
		gorm.Model
		IssuedDate    time.Time
		PaymentAmount float64
		Fee           float64
		FeeRate       float64
		Tax           float64
		TaxRate       float64
		TotalAmount   float64
		DueDate       time.Time
		Status        string
	}

	p.FeeRate = constants.ADMINISTRATIVE_FEE
	p.TaxRate = constants.GOVERNMENT_TAXES

	p.Tax = ic.calculateInvoiceTaxAmount(p.PaymentAmount)
	p.Fee = ic.calculateInvoiceFeeAmount(p.PaymentAmount)
	p.TotalAmount = ic.calculateInvoiceTotalAmount(p.PaymentAmount)

	err = ic.Database.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error creating Invoice: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}



func (ic InvoiceController) calculateInvoiceTotalAmount(paymentAmount float64) float64 {
	return paymentAmount * (1 + (1+constants.GOVERNMENT_TAXES)*constants.ADMINISTRATIVE_FEE)
}

func (ic InvoiceController) calculateInvoiceTaxAmount(paymentAmount float64) float64 {
	return paymentAmount * constants.GOVERNMENT_TAXES
}

func (ic InvoiceController) calculateInvoiceFeeAmount(paymentAmount float64) float64 {
	return paymentAmount * (1+constants.GOVERNMENT_TAXES)*constants.ADMINISTRATIVE_FEE
}