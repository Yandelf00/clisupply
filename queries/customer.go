package queries

import (
	"fmt"
	"supplycli/database"
	"supplycli/models"
)

// function to add a customer
func CreateCustomer() {
	var first_name string
	var last_name string
	var email string
	var phone_number string
	var shipping_address string
	var billing_address string
	fmt.Print("enter customers first name : ")
	fmt.Scanln(&first_name)
	fmt.Print("enter customers last name : ")
	fmt.Scanln(&last_name)
	fmt.Print("enter customers email: ")
	fmt.Scanln(&email)
	fmt.Print("enter customers phone number : ")
	fmt.Scanln(&phone_number)
	fmt.Print("enter customers shipping address: ")
	fmt.Scanln(&shipping_address)
	fmt.Print("enter customers billing address: ")
	fmt.Scanln(&billing_address)
	customer := models.Customer{
		FirstName:       first_name,
		LastName:        last_name,
		Email:           email,
		PhoneNumber:     phone_number,
		ShippingAddress: shipping_address,
		BillingAddress:  billing_address,
	}
	database.DB.Create(&customer)
}

// function to get all the customers
func GetCustomers() {
	var customers []models.Customer
	database.DB.Find(&customers)
	for _, customer := range customers {
		fmt.Printf("first name : %s, last name : %s, email : %s, phone number : %s \n", customer.FirstName, customer.LastName, customer.Email, customer.PhoneNumber)
	}
}
