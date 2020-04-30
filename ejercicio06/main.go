package main

import "fmt"

func main() {

	/*
		//Almacenamos los dos valores que devuelve la misma función,
		//en este caso un Int y un Boll
		numero, estado := dos(2)
		fmt.Printf("El resultado es %d \n", numero)
		fmt.Printf("El valor boolean es %t \n", estado)*/
	fmt.Println(tres(10, 20))
	fmt.Println(tres(11, 22, 45, 60))
	fmt.Println(tres(1, 1, 1, 1, 1))

}

func uno(numero int) int {
	return numero * 2
}

func dos(numero int) (int, bool) {

	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}

}
func tres(numero ...int) int {

	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d:", item)
	}
	return total

}
