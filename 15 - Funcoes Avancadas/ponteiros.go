package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido) // é pego uma cópia do valor
	fmt.Println(numero)          // A variável continua com o mesmo valor

	novoNumero := 40
	fmt.Println(novoNumero)               // 40
	inverterSinalComPonteiro(&novoNumero) // passando o endereço de memória onde a variável está salva
	fmt.Println(novoNumero)               // -40 (o valor negativo foi atribuido a variavel novoNumero)
}
