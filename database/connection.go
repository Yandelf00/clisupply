package database

import (
	"fmt"
	"supplycli/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := "user:passowrd@/database?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("there was an error in the connection")
		return
	}
	DB = db
	db.AutoMigrate(&models.Employee{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Supplier{})
	db.AutoMigrate(&models.Purchase{})
	db.AutoMigrate(&models.Supply{})
}
