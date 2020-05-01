package main

import "fmt"

//PUNTEROS

func main() {

	a := 2
	pA := &a // & Cogemos la dirección de memoria de a

	fmt.Println(pA)
	fmt.Println(*pA) // * accede dentro de pA

	main2(pA)
	fmt.Println(a)

	pNil := returnNil()
	/*Siempre que recibimos un puntero de una función,
	  realizamos un if para ver si lo que contiene es nil*/
	if pNil == nil {
		fmt.Println("Esto es nil y si lo ejecutamos nos lanza un panic")
	} else {
		fmt.Println(*pNil)
	}
}

func main2(pA *int) {
	*pA = 8
}

func returnNil() *int {
	return nil
}
