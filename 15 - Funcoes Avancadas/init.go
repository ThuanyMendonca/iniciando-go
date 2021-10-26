package main

import "fmt"

// Executa a função init primeiro, pode ter somente uma por arquivo
func init() {
	fmt.Println("Executando a função init")
}

func main() {
	fmt.Println("Função main sendo executada")
}
