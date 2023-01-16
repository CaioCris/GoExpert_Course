package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&product{})
	if err != nil {
		return
	}

	//create
	db.Create(&product{Name: "Notebook", Price: 2500.00})

	//create batch
	products := []product{
		{Name: "Mouse", Price: 250.00},
		{Name: "Keyboard", Price: 350.00},
	}
	db.Create(&products)

	//select one
	var prod product

	db.First(&prod, 3)
	fmt.Println(prod)

	db.First(&prod, "name = ?", "Notebook")
	fmt.Println(prod)

	//select all
	var prods []product

	db.Find(&prods)
	for _, prod := range prods {
		fmt.Println(prod)
	}
}
