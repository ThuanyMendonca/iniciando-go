package main

import "fmt"

// Não é exatamente herança, apenas próximo

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // coloca o nome da outra struct sem passar o tipo *Pode passar o tipo, porém para acessar ficaria estudante.pessoa.nome*
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	el := estudante{p1, "ADS", "USP"}
	fmt.Println(el)      // {{João Pedro 20 178} ADS USP}
	fmt.Println(el.nome) // João
}
