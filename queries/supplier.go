package queries

import (
	"fmt"
	"supplycli/database"
	"supplycli/models"
	"time"
)

// func to add supplier
func AddSupplier() {
	var name string
	var email string
	var phone_number string
	var address string
	var city string
	var postalcode string
	var country string
	var productid uint
	fmt.Print("Enter supplier name :")
	fmt.Scanln(&name)
	fmt.Print("Enter supplier email : ")
	fmt.Scanln(&email)
	fmt.Print("Enter supplier phone number : ")
	fmt.Scanln(&phone_number)
	fmt.Print("Enter the productID it supplies : ")
	fmt.Scanln(&productid)
	fmt.Print("Enter supplier address(not mandatory) : ")
	fmt.Scanln(&address)
	fmt.Print("Enter supplier city(not mandatory) : ")
	fmt.Scanln(&city)
	fmt.Print("Enter suppplier postalcode(not mandatory) : ")
	fmt.Scanln(&postalcode)
	fmt.Print("Enter supplier country(not mandatory) : ")
	fmt.Scanln(&country)
	supplier := models.Supplier{
		Name:        name,
		Email:       email,
		PhoneNumber: phone_number,
		ProductId:   productid,
		Address:     address,
		City:        city,
		PostalCode:  postalcode,
		Country:     country,
		CreatedAT:   time.Now(),
		UpdatedAT:   time.Now(),
	}

	database.DB.Create(&supplier)
}

func CheckSuppliers() {
	var suppliers []models.Supplier
	database.DB.Find(&suppliers)
	for _, supplier := range suppliers {
		fmt.Printf("Name : %s, Email : %s, Phone number : %s", supplier.Name, supplier.Email, supplier.PhoneNumber)
	}
}
