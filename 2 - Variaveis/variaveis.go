package main

import "fmt"

// Quando declara variavel e não utiliza, não é possível rodar a aplicação
func main() {
	// explicito
	var variavel1 string = "Variavel 1"

	// implicito - inferência de tipo
	variavel2 := "Variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// trocar de valor entre variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6) 
}
