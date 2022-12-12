package main

import (
	"fmt"

	"github.com/GoLang/math"
	"github.com/google/uuid"
)

func main() {
	s := math.Sum(10, 20)
	id := uuid.New()
	fmt.Printf("Resultado: %v\n", s)
	fmt.Printf("UUID: %v", id)
}
