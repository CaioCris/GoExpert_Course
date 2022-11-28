package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Client struct {
	name     string
	age      int
	isActive bool
	Endereco Endereco
}

func main() {
	client := Client{
		name:     "Caio",
		age:      29,
		isActive: true,
	}
	client.Endereco.Cidade = "São Paulo"

	fmt.Println(client.name)
	fmt.Println(client.isActive)

	client.isActive = false
	fmt.Println(client.isActive)

}
