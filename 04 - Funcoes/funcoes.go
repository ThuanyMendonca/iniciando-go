package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

// Função com mais de 1 retorno
func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(4, 5)
	println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Teste")
	fmt.Println(resultado)

	// Executando para utilizar os 2 retornos
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Executando para utilizar apenas 1 retorno
	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)
}
