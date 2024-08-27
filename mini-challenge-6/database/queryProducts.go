package database

import (
	"database/sql"
	"fmt"
)

// =========== CREATE PRODUCT ===========
func CreateProduct(name string) error {
	query := `INSERT INTO products (name) VALUES ($1)`
	_, err := db.Exec(query, name)
	if err != nil {
		return fmt.Errorf("Error creating product: %v", err)
	}

	fmt.Println("Product created!")
	return nil
}

// =========== UPDATE PRODUCT ===========
func UpdateProduct(oldProductName, newProductName string) error {
	query := `UPDATE products SET name = $1 WHERE name = $2`
	_, err := db.Exec(query, newProductName, oldProductName)
	if err != nil {
		return fmt.Errorf("Error updating product: %v", err)
	}

	fmt.Println("Product updated!")
	return nil
}

// =========== GET PRODUCT BY ID ===========
func GetProductById(productId int) error {
	query := `SELECT * FROM products WHERE id = $1`
	var id int
	var name string

	err := db.QueryRow(query, productId).Scan(&id, &name)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("No product found with ID %d", productId)
		}
		return fmt.Errorf("Can't get product by ID: %v", err)
	}

	indentation := "    "
	fmt.Printf("Getting product by ID ->\n")
	fmt.Printf("%s%-10s: %d\n", indentation, "ID", id)
	fmt.Printf("%s%-10s: %s\n", indentation, "Name", name)
	return nil

}

// =========== GET PRODUCT BY ID ===========
func GetProductWithVariant(productId int) error {
	query := `SELECT products.id, products.name, variants.variant_name, variants.quantity FROM products LEFT JOIN variants ON products.id = variants.product_id WHERE products.id = $1`
	var id, qty int
	var productName string
	var variantName string

	err := db.QueryRow(query, productId).Scan(&id, &productName, &variantName, &qty)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("No product found with ID %d", productId)
		}
		return fmt.Errorf("Can't get product by ID: %v", err)
	}

	indentation := "    "
	fmt.Printf("Getting product by ID ->\n")
	fmt.Printf("%s%-10s: %d\n", indentation, "ID", id)
	fmt.Printf("%s%-10s: %s\n", indentation, "Product Name", productName)
	fmt.Printf("%s%-10s: %s\n", indentation, "Variant Name", variantName)
	fmt.Printf("%s%-10s: %d\n", indentation, "Quantity", qty)
	return nil
}

func GetProductWithVariant2(productId int) error {
	query := `SELECT products.id, products.name, variants.variant_name, variants.quantity FROM products LEFT JOIN variants ON products.id = variants.product_id WHERE products.id = $1`
	rows, err := db.Query(query, productId)
	if err != nil {
		return fmt.Errorf("Can't get product with variants: %v", err)
	}
	defer rows.Close()

	var product struct {
		id   int
		name string
	}

	var variants []struct {
		name     string
		quantity int
	}

	for rows.Next() {
		var (
			id          int
			productName string
			variantName sql.NullString
			variantQty  sql.NullInt64
		)

		err := rows.Scan(&id, &productName, &variantName, &variantQty)
		if err != nil {
			return fmt.Errorf("Error scanning row: %v", err)
		}

		if product.id == 0 {
			product.id = id
			product.name = productName
		}

		if variantName.Valid && variantQty.Valid {
			variants = append(variants, struct {
				name     string
				quantity int
			}{
				name:     variantName.String,
				quantity: int(variantQty.Int64),
			})
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("Error iterating rows: %v", err)
	}

	if product.id == 0 {
		return fmt.Errorf("No product found with ID %d", productId)
	}

	indentation := "    "
	fmt.Printf("Product details:\n")
	fmt.Printf("%s%-14s: %d\n", indentation, "ID", product.id)
	fmt.Printf("%s%-14s: %s\n", indentation, "Name", product.name)
	fmt.Printf("\nVariants:\n")

	if len(variants) == 0 {
		fmt.Printf("%sNone\n", indentation)
	} else {
		for _, variant := range variants {
			fmt.Printf("%s%-14s: %s\n", indentation, "Variant Name", variant.name)
			fmt.Printf("%s%-14s: %d\n\n", indentation, "Quantity", variant.quantity)
		}
	}

	return nil
}
