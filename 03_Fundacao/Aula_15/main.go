package main

func main() {

	// Memória -> Endereço -> Valor
	 a := 10
	 println(a) // Valor
	 println(&a) // Endereço

	var pointer *int = &a
	println(pointer)
	*pointer = 20
	println(a)

	b := &a
	println(b)
	println(*b)
}
