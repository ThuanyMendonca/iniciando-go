package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// Quando a variavel é criada dessa forma, ela não é acessível fora do escopo
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que 0")
	}

	// fmt.Println(outroNumero) // Não funciona, pois está apenas no escopo do if
}
