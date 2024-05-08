package models

import (
	"time"
	"gorm.io/gorm"
   
)
type Company struct {
	gorm.Model
	LegalName     string
	RepresentativeName string
	PhoneNumber   string
	PostalCode    string
	Address       string
	Users         []User `gorm:"foreignKey:ID;references:ID"`
	Customers     []Customer `gorm:"foreignKey:ID;references:ID"`
	Invoices      []Invoice `gorm:"foreignKey:ID;references:ID"`
}

type User struct {
	gorm.Model
	FullName      string
	Email         string
	Password      string
}

type Customer struct {
	gorm.Model
	LegalName     string
	RepresentativeName string
	PhoneNumber   string
	PostalCode    string
	Address       string
	BankAccount   BankAccount
	Invoices      []Invoice `gorm:"foreignKey:ID;references:ID"`
}

type BankAccount struct {
	gorm.Model
	CustomerID    string
	BankName      string
	BranchName    string
	AccountNumber string
	AccountName   string
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
