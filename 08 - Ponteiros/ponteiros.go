package main

import "fmt"

func main() {
	fmt.Println("Ponteiros") // É uma referencia de memória

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // Endereço de memória onde variavel3 está salvo

	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, ponteiro)  // output -> 100 0xc000012088
	fmt.Println(variavel3, *ponteiro) // (desreferenciação) output -> 100 100
}
