package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)

	}
	fmt.Println(i)

	fmt.Println("---------------------")
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}
	//fmt.Println(j) // Nesse caso será undefined, pois a variável j está somente no escopo do for

	nomes := [3]string{"João", "Davi", "Lucas"}
	// Mostrando os dois indices e valores, caso queira mostrar apenas um valor usar o _
	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}

	usuario := map[string]string{
		"nome":      "Marciel",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	// loop infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
