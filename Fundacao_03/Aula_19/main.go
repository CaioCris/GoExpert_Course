package main

import "fmt"

func main() {

	var name interface{} = "Caio Cris"
	print(name.(string))
	res, ok := name.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
	res2 := name.(int)
	fmt.Printf("O valor de res2 é %v", res2)

}
