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
