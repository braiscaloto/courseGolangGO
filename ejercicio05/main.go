package main

import (
	"fmt"
)

func main() {

	/*numero := 0
	for {
		fmt.Println("Continúo")
		fmt.Println("Ingrese el número clave")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}

	}*/
	/*
		var i = 0

		for i < 10 {
			fmt.Printf("\n Valor de i: %d", i)
			if i == 5 {
				fmt.Printf("    Multiplicamos por 2 \n")
				i = i * 2
				continue
			}
			fmt.Printf("   Pasó por aquí \n")
			i++

		}*/

	var i = 0

RUTINA:
	fmt.Println("Este es el inicio de la RUTINA")
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("Voy a RUTINA sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("El valor de i: %d \n", i)
		i++
	}

}
