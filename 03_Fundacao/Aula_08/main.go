package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(50, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)

}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior do que 50")
	} else {
		return a + b, nil
	}
}
