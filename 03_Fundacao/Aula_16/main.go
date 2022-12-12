package main

func sum(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	a := 10
	b := 20
	println(sum(&a,&b))
	println(a)
}
