package database

import "fmt"

func CreateProduct(name string) error {
	query := `INSERT INTO products (name) VALUES ($1)`
	_, err := db.Exec(query, name)
	if err != nil {
		return fmt.Errorf("Error creating product: %v", err)
	}

	fmt.Println("Product created!")
	return nil
}

func UpdateProduct(oldProductName, newProductName string) error {
	query := `UPDATE products SET name = $1 WHERE name = $2`
	_, err := db.Exec(query, newProductName, oldProductName)
	if err != nil {
		return fmt.Errorf("Error updating product: %v", err)
	}

	fmt.Println("Product updated!")
	return nil
}
