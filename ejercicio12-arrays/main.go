package main

import (
	"fmt"
)

func main() {
	// Arrays con número de variables definido

	var holaMundo [2]string
	holaMundo[0] = "Hola"
	holaMundo[1] = "Mundo"

	/*fmt.Println(holaMundo)
	var numeros []int
	numeros = append(numeros, 0)
	fmt.Println(len(numeros))
	fmt.Println(cap(numeros))
	numeros = append(numeros, 1, 2, 3, 4, 5)
	fmt.Println(numeros)*/

	letras := []string{"a", "b", "c"}
	letras = append(letras, "d")
	fmt.Println(len(letras))
	fmt.Println(cap(letras[0:2]))

	for _, letra := range letras {
		fmt.Println(letra)
	}

}
