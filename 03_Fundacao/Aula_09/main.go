package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 3, 45, 6, 34, 654, 100))

}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
