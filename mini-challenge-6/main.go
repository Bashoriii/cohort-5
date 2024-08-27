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

	// GET PRODUCT BY ID
	// err = database.GetProductById(3)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// CREATE A VARIANT
	// err = database.CreateVariant("Pink", 9, 4)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// UPDATE VARIANT BY ID
	// err = database.UpdateVariant("Yellow Black", 3)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// DELETE VARIANT BY ID
	// err = database.DeleteVariantById(6)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// GET PRODUCT WITH VARIANT BY ID
	err = database.GetProductWithVariant2(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
