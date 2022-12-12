package main

import "fmt"

type MyIntNumber int

// Usar o ~ele considera que os tipos criados com tipos "primitivos" sejam considerados
type Number interface {
	~int | float64
}

func Sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false

}

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 1000.20, "João": 2000.10, "Maria": 3000}
	m3 := map[string]MyIntNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(Sum(m))
	fmt.Println(Sum(m2))
	fmt.Println(Sum(m3))
	fmt.Println(Compare(10, 10))
}
