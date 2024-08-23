package database

import "fmt"

func CreateVariant(variant string, qty, product_id int) error {
	query := `INSERT INTO variants (variant_name, quantity, product_id) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, variant, qty, product_id)
	if err != nil {
		return fmt.Errorf("Error creating product: %v", err)
	}

	fmt.Println("Variant created!")
	return nil
}

func UpdateVariant(prodId int, param string) error {
	query := `UPDATE variants SET product_id = $1 WHERE variant_name = $2`
	_, err := db.Exec(query, prodId, param)
	if err != nil {
		return fmt.Errorf("Error updating product: %v", err)
	}

	fmt.Println("Variant updated!")
	return nil
}
