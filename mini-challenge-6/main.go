package main

import (
	"database/sql"
	"fmt"
	"go-sql/database"
)

var db *sql.DB

func main() {
	db := database.DatabaseConnection()
	defer db.Close()

	// CreateProduct()
	err := database.CreateProduct("Odol Pepsodent")
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
