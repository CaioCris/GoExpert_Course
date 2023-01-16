package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type product struct {
	ID    string
	Name  string
	Price float64
}

func newProduct(name string, price float64) *product {
	return &product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	product := newProduct("Notebook", 1899.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = 3000
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	prod, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de %.2f", prod.Name, prod.Price)
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, prod := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", prod.Name, prod.Price)
	}
	err = deleteProduct(db, product.ID)
}

func insertProduct(db *sql.DB, product *product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?,?,?)")
	if err != nil {
		panic(err)
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		panic(err)
	}
	return nil
}

func updateProduct(db *sql.DB, product *product) error {
	stmt, err := db.Prepare("update products set name = ? , price = ? where id = ?")
	if err != nil {
		panic(err)
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		panic(err)
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)
	var prod product
	err = stmt.QueryRow(id).Scan(&prod.ID, &prod.Name, &prod.Price)
	if err != nil {
		panic(err)
	}
	return &prod, nil
}

func selectAllProducts(db *sql.DB) ([]product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var products []product
	for rows.Next() {
		var prod product
		err = rows.Scan(&prod.ID, &prod.Name, &prod.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, prod)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	return nil
}
