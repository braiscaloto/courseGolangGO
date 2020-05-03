package main

import "fmt"

func run(num int, callback func(int)) {
	for i := 0; i < 10; i++ {
		fmt.Printf("GoRutina %v %v\n", num, i)
	}
	callback(num)
}

func main() {

	Callback := func(numeroDeGoRutina int) {

		fmt.Printf("Ha finalizado la GoRutina %v\n", numeroDeGoRutina)

	}

	for i := 0; i < 100; i++ {
		go run(i, Callback)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
