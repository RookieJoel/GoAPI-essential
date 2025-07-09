package main

import (
	"database/sql"
	"fmt"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int `json:"price"`
}

func createProduct(product *Product) error {
	// Insert product into the database
	// query := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`
	_ , err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("could not create product: %v", err)
	}
	return nil

}

func getProductByID(id int) (*Product, error) {
	var product Product
	// query := `SELECT id, name, price FROM products WHERE id = $1`
	row := db.QueryRow("SELECT id, name, price FROM products WHERE id = $1", id)
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no product found with ID %d", id)
		}
		return nil, fmt.Errorf("could not get product: %v", err)
	}
	return &product, nil
}

func getAllProducts() ([]Product, error) {
	var products []Product
	// query := `SELECT id, name, price FROM products`
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, fmt.Errorf("could not get products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, fmt.Errorf("could not scan product: %v", err)
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over products: %v", err)
	}

	return products, nil
}

func updateProduct(id int, product *Product) error {
	// Update product in the database
	// query := `UPDATE products SET name = $1, price = $2 WHERE id = $3`
	_, err := db.Exec("UPDATE products SET name = $1, price = $2 WHERE id = $3", product.Name, product.Price, id)
	if err != nil {
		return fmt.Errorf("could not update product: %v", err)
	}
	return nil
}

func deleteProduct(id int) error {
	// Delete product from the database
	// query := `DELETE FROM products WHERE id = $1`
	_, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("could not delete product: %v", err)
	}
	return nil
}