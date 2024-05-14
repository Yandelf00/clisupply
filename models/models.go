package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// here I'm grouping the needed constants for the file

const (
	// Employee roles
	RoleAdmin    = "admin"
	RoleManager  = "manager"
	RoleEmployee = "employee"

	// Product Categories
	SmartphoneCat = "smartphone"
	TabletCat     = "tablet"
	LaptopCat     = "laptop"
	ComputerCat   = "computer"
	TvCat         = "tv"
)

// this is the product Model :
type Product struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null"`
	Description   string
	Category      string     `gorm:"not null"`
	Price         uint       `gorm:"not null"`
	StockQuantity uint       `gorm:"not null"`
	Suppliers     []Supplier `gorm:"foreignKey:ProductID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// this is the employee model
type Employee struct {
	ID          uint   `gorm:"primaryKey"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Role        string `gorm:"not null"`
	PhoneNumber string `gorm:"not null;unique"`
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// this function makes sure that employee is one of the three we defined
func (e *Employee) BeforeSave(tx *gorm.DB) (err error) {
	if e.Role != RoleAdmin && e.Role != RoleEmployee && e.Role != RoleManager {
		return fmt.Errorf("invalid role: %s", e.Role)
	}
	return nil
}

// this is the customer model
type Customer struct {
	ID              uint   `gorm:"primaryKey"`
	FirstName       string `gorm:"not null"`
	LastName        string `gorm:"not null"`
	Email           string `gorm:"unique;not null"`
	PhoneNumber     string `gorm:"unique"`
	ShippingAddress string
	BillingAddress  string
}

// this is the supplier struct
type Supplier struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Email       string `gorm:"not null;unique"`
	PhoneNumber string `gorm:"not null;unique"`
	Address     string
	City        string
	PostalCode  string
	Country     string
	ProductID   uint `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// this is the purchase struct
// Purchase model
type Purchase struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	CustomerID uint `gorm:"not null"`
	Customer   Customer

	ProductID uint `gorm:"not null"`
	Product   Product

	Quantity uint `gorm:"not null"`
}

// Supply model
type Supply struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	SupplierID uint `gorm:"not null"`
	Supplier   Supplier

	Quantity   uint `gorm:"not null"`
	TotalPrice uint `gorm:"not null"`
}
