package controller

import (
	"fmt"
	"strconv"
	"strings"
	"supplycli/queries"
)

type Input struct {
	command string
	args    []string
}

var input_instance Input

func SetInputInstance(inpt string) {
	if len(inpt) > 0 {
		tmp := strings.Fields(inpt)
		input_instance.command = tmp[0]
		input_instance.args = tmp[1:]
		Controller()
	}
}

func Controller() {
	switch input_instance.command {
	case "get":
		GetController()
	case "create":
		CreateController()
	case "make":
		MakeController()
	case "help":
		HelpController()
	default:
		fmt.Println("command unknown")
	}
}

func GetController() {
	if len(input_instance.args) > 0 {
		switch input_instance.args[0] {
		case "employees":
			Employees()
		case "customers":
			Customers()
		case "suppliers":
			Suppliers()
		case "products":
			Products()
		default:
			fmt.Println("idk what are you trying to get")
		}
	} else {
		fmt.Println("not enough arguments")
	}
}

func Products() {
	if len(input_instance.args) > 1 {
		switch input_instance.args[1] {
		case "all":
			queries.GetAllProducts()
		case "category":
			if len(input_instance.args) > 2 {
				queries.GetProductsByCategory(input_instance.args[2])
			} else {
				fmt.Println("not enough arguments")
			}
		case "name":
			if len(input_instance.args) > 2 {
				queries.GetProductsByName(input_instance.args[2])
			} else {
				fmt.Println("not enough arguments")
			}
		case "id":
			if len(input_instance.args) > 2 {
				prodid, err := strconv.ParseUint(input_instance.args[2], 10, 64)
				if err != nil {
					fmt.Println("there was an error parsing the id")
					return
				}
				queries.GetProductById(uint(prodid))
			} else {
				fmt.Println("not enough arguments")
			}
		default:
			fmt.Println("nothing")
		}
	} else {
		fmt.Println("not enough arguments")
	}
}

func Employees() {
	if len(input_instance.args) > 1 {
		switch input_instance.args[1] {
		case "all":
			queries.CheckAllEmployees()
		case "employees":
			queries.CheckEmployees()
		case "managers":
			queries.CheckManagers()
		case "admins":
			queries.CheckAdmins()
		default:
			fmt.Println("no role as such")
		}
	} else {
		fmt.Println("not enough arguments")
	}
}

func Customers() {
	if len(input_instance.args) > 1 {
		switch input_instance.args[1] {
		case "all":
			queries.CreateCustomer()
		default:
			fmt.Println("not available")
		}
	} else {
		fmt.Println("not enough arguments")
	}
}

func Suppliers() {
	if len(input_instance.args) > 1 {
		switch input_instance.args[1] {
		case "all":
			queries.CheckSuppliers()
		default:
			fmt.Println("idk man")
		}
	} else {
		fmt.Println("not enough arguments")
	}
}

func CreateController() {

}
func MakeController() {}
func HelpController() {}
