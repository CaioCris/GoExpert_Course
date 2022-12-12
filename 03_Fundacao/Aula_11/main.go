package main

import "fmt"

type Client struct {
	name     string
	age      int
	isActive bool
}

func main() {
	client := Client{
		name:     "Caio",
		age:      29,
		isActive: true,
	}

	fmt.Println(client.name)
	fmt.Println(client.isActive)

	client.isActive = false
	fmt.Println(client.isActive)

}
