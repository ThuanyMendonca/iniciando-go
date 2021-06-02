package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string // é necessário especificar tamanho
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// Slice
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7} // ... é uma flexibilização atribuir qualquer quantidade de atributos
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18) // Adicionando item no slice e retorna um novo slice
	fmt.Println(slice)

	slice2 := array2[1:3] // Pegando os valores da posição 1 até a posição 3 (não entra a última posição)
	fmt.Println(slice2)   // Output -> [Posição 2 Posição 3]

	array2[1] = "Posição alterada" // Quando altera aqui, já altera o valor na slice2 pois está apontando para ela
	fmt.Println(slice2)            // Output -> [Posição alterada Posição 3]

}
