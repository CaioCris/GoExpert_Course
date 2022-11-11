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
	fmt.Printf("O tipo de E é %T \n", e)
	fmt.Printf("O valor de E é %v \n", e)
	fmt.Printf("O tipo de F é %T \n", f)
	fmt.Printf("O valor de F é %v", f)
}
