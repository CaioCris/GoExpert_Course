package math

// Iniciado com letra MAIUSCULA ele é exportado
// Iniciado com letra minuscula é privado
func Sum[T int | float64](a, b T) T {
	return a + b
}
