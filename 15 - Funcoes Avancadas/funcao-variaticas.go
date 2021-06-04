package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
	//fmt.Println(numeros)
}

func main() {
	// Função que recebe quantos parametros precisar
	totalDaSoma := soma(1, 2, 3, 6, 5, 2, 3, 6, 6, 5, 56, 62, 613, 6161, 6161, 6165, 613, 6151, 855, 789)
	fmt.Println(totalDaSoma)
}
