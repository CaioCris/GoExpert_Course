package main

import "fmt"

func main() {

	salaries := map[string]int{"Caio": 1000, "Rebeca": 2000, "Wesley": 3000}
	fmt.Println(salaries["Caio"])
	delete(salaries, "Wesley")
	salaries["Antonio"] = 4000
	fmt.Println(salaries["Antonio"])

	sal := make(map[string]int)
	sal["Caio"] = 2000
	fmt.Println(sal["Caio"])

	sal1 := map[string]int{}
	sal1["Rebeca"] = 3000
	fmt.Println(sal1["Rebeca"])

	for name, salary := range salaries {
		fmt.Printf("O salario de %s é %d\n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("O salario é %d\n", salary)
	}
}
