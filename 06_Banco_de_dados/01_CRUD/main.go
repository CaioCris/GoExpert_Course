package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
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
	err = insertProduct(db, NewProduct("Notebook", 1899.90))
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
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
