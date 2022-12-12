package main

// Apenas FOR para usar loop
func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	// podemos usar um _ para esconder o K (chave/index) ou o V (value)
	numbers := []string{"um", "dois", "três"}
	for k, v := range numbers {
		println(k, v)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("loop infinito, usado em algumas situações")
	}

}
