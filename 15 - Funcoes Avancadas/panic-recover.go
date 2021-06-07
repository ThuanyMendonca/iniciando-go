package main

import "fmt"

func recuperarExecucao() {
	// recover está pegando exatamente a mensagem que iria quebrar o código
	// r -> output (A média é exatamente 6)
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar a execução")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	// Mata a execução do programa
	panic("A média é exatamente 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 8))
	fmt.Println("Pós execução!")
}
