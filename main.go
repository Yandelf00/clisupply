package main

import (
	"bufio"
	"fmt"
	"os"
	"supplycli/controller"
	"supplycli/database"
	"supplycli/logged"
	"supplycli/models"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	database.Connection()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter your username : ")
	scanner.Scan()
	username := scanner.Text()
	fmt.Print("Please enter your password : ")
	scanner.Scan()
	password := scanner.Text()
	var employee models.Employee
	if err := database.DB.Where("first_name = ?", username).Find(&employee).Error; err != nil {
		fmt.Println("Error finding the employee: ", err)
	}
	err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(password))
	if err != nil {
		fmt.Println("thats not the correct password")
		return
	}
	// sets the logged in user
	logged.SetLoggedAcc(employee.ID, employee.Role)
	fmt.Println("you are connected !")
	for {
		fmt.Print(">>> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}
		// Call the controller inside the loop
		controller.SetInputInstance(input)
	}
}
