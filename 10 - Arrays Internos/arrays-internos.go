package main

import "fmt"

func main() {
	// make - Aloca um espaço na memória para algo
	slice := make([]float64, 10, 11)
	fmt.Println(slice)
	fmt.Println(len(slice)) // Verificar o tamanho length
	fmt.Println(cap(slice)) // Verificar a capacidade

	// Caso ultrapasse a capacidade do slice, ele dobra de tamanho

	slice2 := make([]float32, 5)
	fmt.Println(slice2)
	slice2 = append(slice2, 10)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}
