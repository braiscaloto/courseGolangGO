package main

import "fmt"

var estado bool

func main() {

	estado = true
	if estado = false; estado == true {
		fmt.Println("Es verdadero")
	} else {
		fmt.Println("Es falso")
	}

	switch numero1 := 5; numero1 {
	case 3:
		println("Es tres")
	case 5:
		println("Es cinco")
	}

}
