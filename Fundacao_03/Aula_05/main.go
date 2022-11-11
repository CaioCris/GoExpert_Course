package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool = true
	c int
	d string
	e float64
	f ID = 10
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(len(meuArray))
	fmt.Println(meuArray[len(meuArray)-1])

	for index, value := range meuArray {
		fmt.Printf("O indice é %d e o valor é %d\n", index, value)
	}
}
