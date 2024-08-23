package main

import (
	"fmt"
	"go-sql/database"
)

func main() {
	db := database.DatabaseConnection()
	defer db.Close()

	var err error
	// CREATE A PRODUCT
	// err = database.CreateProduct("Hand Puff")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// UPDATE PRODUCT
	// err = database.UpdateProduct("Shampoo", "Zebracross")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// CREATE A VARIANT
	// err = database.CreateVariant("Magic White", 8, 3)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// UPDATE VARIANT
	err = database.UpdateVariant(2, "Yellow Black")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// func CreateProduct() {
// 	query := `INSERT INTO products (name) VALUES ($1)`
// 	_, err := db.Exec(query, "Odol pepsodent")

// 	if err != nil {
// 		fmt.Println("Error creating product", err)
// 	}
// 	fmt.Println("Product created!")
// }
