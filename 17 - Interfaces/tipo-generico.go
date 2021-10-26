package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	// Interface genérica recebe qualquer tipo de dados
	generica("String")
	generica(1)
	generica(true)

	// O map aceita qualquer tipo de chave e qualquer tipo de valor
	mapa := map[interface{}]interface{}{
		1:            "String",
		int64(12321): true,
		"String":     "ABC",
	}

	println(mapa)
}
