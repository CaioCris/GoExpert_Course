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
	Endereco
}

func (client Client) Deactivate() {
	client.isActive = false
	fmt.Printf("O cliente %s foi desativado", client.name)
}

func main() {
	client := Client{
		name:     "Caio",
		age:      29,
		isActive: true,
	}
	client.Cidade = "São Paulo"

	fmt.Println(client.name)
	fmt.Println(client.Cidade)
	fmt.Println(client.isActive)

	client.Deactivate()
}
