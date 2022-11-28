package main

import "fmt"

type Address struct {
	street string
	number int
	city string
	state string
}

type Person interface{
	// Não podemos colocar propriedades em interfaces. 
	// Os métodos tem que ter a mesma assinatura, para a interface funcionar
	Deactivate()
}

type Company struct {
	name string
}

func (company Company) Deactivate() {
}

type Client struct {
	name     string
	age      int
	isActive bool
	Address
}

func (client Client) Deactivate() {
	client.isActive = false
	fmt.Printf("O cliente %s foi desativado", client.name)
}

func Deactivation(person Person){
	person.Deactivate()
}

func main() {
	client := Client{
		name:     "Caio",
		age:      29,
		isActive: true,
	}
	myCompany := Company{}

	Deactivation(client)
}
