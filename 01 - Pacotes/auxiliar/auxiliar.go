package auxiliar

import "fmt"

// Escrever registra uma mensagem na tela
// Obs: dentro do próprio pacote consegue chamar outra função que começa com letra minuscula, fora não
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
