package main

import "fmt"

func main() {
	// Aritméticos
	soma := 1 + 2
	subtracao := 3 - 1
	divisao := 25 / 5
	multiplicacao := 5 * 8
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Para qualquer tipo de comparação e/ou contas
	var numero1 int16 = 10
	var numero2 int16 = 4

	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	fmt.Println(1 == 2)

	fmt.Println("---------------------")

	// Operadores lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// Operadores unários
	numero := 10
	numero++     // incrementa +1
	numero += 15 // incrementa valor digitado

	println(numero)

	numero--    // subtrai 1
	numero -= 5 // subtrai 5

	numero *= 3 // multiplica por 3
	numero /= 2 // divide por 2
	numero %= 2 // resto da divisão

	fmt.Println(numero)

	// Operador ternário (não possui)

}
