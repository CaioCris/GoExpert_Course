package main

import "fmt"

func main() {
	doubleTotal := func() int {
		return sum(1, 3, 45, 6, 34, 654, 100) * 2
	}()

	fmt.Println(doubleTotal)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
