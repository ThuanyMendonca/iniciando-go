package main

import (
	"fmt"
	"math"
)

type IForma interface {
	area() float64
}

func escreverArea(f IForma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

// o método precisa ter a mesma assinatura que a interface
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
