package queries

import (
	"fmt"
	"supplycli/database"
	"supplycli/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// function to add employees
func AddEmployee() {
	var firstname string
	var lastname string
	var email string
	var password string
	var role string
	var phone_number string
	var address string
	fmt.Print("Enter employee first name : ")
	fmt.Scanln(&firstname)
	fmt.Print("Enter employee last name : ")
	fmt.Scanln(&lastname)
	fmt.Print("Enter employee email address : ")
	fmt.Scanln(&email)
	fmt.Print("Enter employee password : ")
	fmt.Scanln(&password)
	fmt.Print("Enter employee role : ")
	fmt.Scanln(&role)
	fmt.Print("Enter employee phone number : ")
	fmt.Scanln(&phone_number)
	fmt.Print("Enter employee address (not required): ")
	fmt.Scanln(&address)
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("there was an error hashing the password")
	}
	employee := models.Employee{
		FirstName:   firstname,
		LastName:    lastname,
		Email:       email,
		Password:    string(hashedpassword),
		Role:        role,
		PhoneNumber: phone_number,
		Address:     address,
		CreatedAT:   time.Now(),
		UpdatedAT:   time.Now(),
	}
	database.DB.Create(&employee)
}

//func updateEmployee() {
//}

//func deleteEmployee() {
//}

// this function shows all the employees
func CheckAllEmployees() {
	var employees []models.Employee
	database.DB.Find(&employees)
	for _, employee := range employees {
		fmt.Printf("First name : %s, Last name : %s, Email : %s, Phone number : %s \n", employee.FirstName, employee.LastName, employee.Email, employee.PhoneNumber)
	}
}

// this function shows only the employees with the role "employee"
func CheckEmployees() {
	role := "employee"
	var employees []models.Employee
	database.DB.Where("role = ?", role).Find(&employees)
	for _, employee := range employees {
		fmt.Printf("First name : %s, Last name : %s, Email : %s, Phone number : %s \n", employee.FirstName, employee.LastName, employee.Email, employee.PhoneNumber)
	}
}

// this function shows only the managers
func CheckManagers() {
	role := "manager"
	var managers []models.Employee
	database.DB.Where("role = ?", role).Find(&managers)
	for _, manager := range managers {
		fmt.Printf("First name : %s, Last name : %s, Email : %s, Phone number : %s \n", manager.FirstName, manager.LastName, manager.Email, manager.PhoneNumber)
	}
}

// this function shows only the admins
func CheckAdmins() {
	role := "admin"
	var admins []models.Employee
	database.DB.Where("role = ?", role).Find(&admins)
	for _, admin := range admins {
		fmt.Printf("First name : %s, Last name : %s, Email : %s, Phone number : %s \n", admin.FirstName, admin.LastName, admin.Email, admin.PhoneNumber)
	}
}
