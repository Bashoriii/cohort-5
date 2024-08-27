package database

import "fmt"

// =========== CREATE VARIANT ===========
func CreateVariant(variant string, qty, product_id int) error {
	query := `INSERT INTO variants (variant_name, quantity, product_id) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, variant, qty, product_id)
	if err != nil {
		return fmt.Errorf("Error creating variant: %v", err)
	}

	fmt.Println("Variant created!")
	return nil
}

// =========== UPDATE VARIANT BY ID ===========
func UpdateVariant(param string, id int) error {
	query := `UPDATE variants SET variant_name = $1 WHERE id = $2`
	_, err := db.Exec(query, param, id)
	if err != nil {
		return fmt.Errorf("Error updating variant: %v", err)
	}

	fmt.Println("Variant updated!")
	return nil
}

// =========== DELETE VARIANT BY ID ===========
func DeleteVariantById(id int) error {
	query := `DELETE FROM variants WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error deleting variant: %v", err)
	}

	fmt.Println("Variant deleted!")
	return nil
}
