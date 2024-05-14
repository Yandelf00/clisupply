package queries

import (
	"fmt"
	"strconv"
	"strings"
	"supplycli/database"
	"supplycli/models"
	"time"
)

func CreateProduct() {
	var name string
	var description string
	var category string
	var price uint
	var stockQuantity uint
	var supplierIdsStg string
	var suppliers []models.Supplier
	fmt.Print("Enter the product name :")
	fmt.Scanln(&name)
	fmt.Print("Enter the product description : ")
	fmt.Scanln(&description)
	fmt.Print("Enter the product category :")
	fmt.Scanln(&category)
	fmt.Print("Enter the product price : ")
	fmt.Scanln(&price)
	fmt.Print("Enter the stock quantity : ")
	fmt.Scanln(&stockQuantity)
	fmt.Printf("Enter the ids of the suppliers of this product(space separated) :")
	tmp := strings.Fields(supplierIdsStg)
	for _, supplierIdstr := range tmp {
		var suppliertmp models.Supplier
		supplierId, err := strconv.ParseUint(supplierIdstr, 10, 64)
		if err != nil {
			fmt.Printf("Invalid ID: %s\n", supplierIdstr)
			return
		}
		database.DB.Where("id = ?", supplierId).First(&suppliertmp)
		suppliers = append(suppliers, suppliertmp)
	}
	product := models.Product{
		Name:          name,
		Description:   description,
		Category:      category,
		Price:         price,
		StockQuantity: stockQuantity,
		Suppliers:     suppliers,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	database.DB.Create(&product)
}

func GetAllProducts() {
	var products []models.Product
	database.DB.Find(&products)
	for _, product := range products {
		fmt.Printf("Name : %s, Category : %s, Price : %d, stock Quantity: %d", product.Name, product.Category, product.Price, product.StockQuantity)
	}
}

func GetProductsByCategory(category string) {
	var products []models.Product
	database.DB.Where("category = ?", category).Find(&products)
	for _, product := range products {
		fmt.Printf("Name : %s, Category : %s, Price : %d, stock Quantity: %d", product.Name, product.Category, product.Price, product.StockQuantity)
	}
}

func GetProductsByName(name string) {
	var products []models.Product
	database.DB.Where("name = ?", name).Find(&products)
	for _, product := range products {
		fmt.Printf("Name : %s, Category : %s, Price : %d, stock Quantity: %d", product.Name, product.Category, product.Price, product.StockQuantity)
	}
}

func GetProductByMaxPrice(price uint) {
	var products []models.Product
	database.DB.Where("price <= ?", price).Find(&products)
	for _, product := range products {
		fmt.Printf("Name : %s, Category : %s, Price : %d, stock Quantity: %d", product.Name, product.Category, product.Price, product.StockQuantity)
	}
}

func GetProductByMinPrice(price uint) {
	var products []models.Product
	database.DB.Where("price >= ?", price).Find(&products)
	for _, product := range products {
		fmt.Printf("Name : %s, Category : %s, Price : %d, stock Quantity: %d", product.Name, product.Category, product.Price, product.StockQuantity)
	}
}

func GetProductById(id uint) {
	var products []models.Product
	database.DB.Where("id = ?", id).Find(&products)
	for _, product := range products {
		fmt.Printf("Name : %s, Category : %s, Price : %d, stock Quantity: %d", product.Name, product.Category, product.Price, product.StockQuantity)
	}
}
