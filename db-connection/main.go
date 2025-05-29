package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	name string
	price float64
}

func main() {
	connStr := ""

	db , err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	
	createProductTable(db)

	product := Product{name: "Book", price: 15.55}
	pk := InsertIntoProduct(db, product)

	fmt.Printf("ID = %d\n", pk)
}

func createProductTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

func InsertIntoProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name , price)
		VALUES ($1, $2) RETURNING id`

	var pk int 
	err := db.QueryRow(query, product.name, product.price).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}

	return pk
}