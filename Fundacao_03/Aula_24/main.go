package main

func main() {

	a := 1
	b := 2
	c := 3

	if a > b {
		println(a)
	} else {
		println(b)
	}

	if a > b && c > a || b < c {
		println("a > b AND c > a OR b < c")
	}

	switch a {
	case 1:
		println("isso")
	case 2:
		println("não")
	default:
		println("nenhuma anterior")
	}

}
