package main

import (
	"errors"
	"fmt"
)

func main() {
	// int - de acordo com a arquitetura do computador, se for 64 bits é o mesmo que int64
	// int8, int16, int32, int64, int
	var numero int = 1000000000
	fmt.Println(numero)

	// unsyned
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias - nome (rune = int32)
	var numero3 rune = 12345
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	// Não pode ser declarado apenas float
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12556653.45
	fmt.Println(numeroReal2)

	// string
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// pegar número da tabela ascii (somente um caracter com aspas simples)
	char := 'B'
	fmt.Println(char)

	// valores default
	var texto string
	fmt.Println(texto) // vazio

	var num string
	fmt.Println(num) // 0

	// boolean
	var booleano1 bool = true
	fmt.Println(booleano1)

	// criando um erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro) // default = nil

}
