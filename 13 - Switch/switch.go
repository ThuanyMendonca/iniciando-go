package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		//return "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
		//return "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
		//return "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
		//return "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
		//return "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
		//return "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
		return "Sábado"
	default:
		diaDaSemana = "Número inválido"
		//return "Número inválido"
	}

	return diaDaSemana
}

func main() {
	// No Go não possui o break, quando a condição é satisfeita
	fmt.Println("Switch")
	dia := diaDaSemana(7)
	fmt.Println(dia)
}
