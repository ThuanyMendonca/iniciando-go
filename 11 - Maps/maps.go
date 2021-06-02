package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Deve ser do mesmo tipo
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[int]string{
		1: "Pedro",
		2: "Silva",
	}
	fmt.Println(usuario2)

	// Chave é uma string que possui uma chave de string que tem o valor de string
	usuario3 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"faculdade": "Fatec",
			"nome":      "ADS",
		},
	}
	fmt.Println(usuario3)
	// Deletar uma chave
	delete(usuario3, "nome")
	fmt.Println(usuario3)

	// Adicionar uma chave
	usuario3["signo"] = map[string]string{
		"nome": "Capricornio",
	}
	fmt.Println(usuario3)
}
