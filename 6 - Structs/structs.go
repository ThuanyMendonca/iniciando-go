package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u) // Quando ainda não foi passado valor para a variável, é exibido os valores default
	u.nome = "Maria"
	u.idade = 14
	fmt.Println(u) // {Maria 14}

	enderecoExemplo := endereco{"Rua de teste", 12}
	usuario2 := usuario{"Maria", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "José"} // Passando apenas um atributo
	fmt.Println(usuario3)
}
