package main

import "fmt"

func main() {

	//CLOSURES
	tablaDel := 2
	MiTabla := tabla(tablaDel)

	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}

}

func tabla(numero int) func() int {

	valor := numero
	secuencia := 0

	return func() int {
		secuencia++
		return valor * secuencia
	}

}
