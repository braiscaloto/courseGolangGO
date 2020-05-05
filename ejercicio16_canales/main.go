package main

import "fmt"

func suma(a, b int, canal chan int) {

	canal <- a + b
}

func main() {
	canal := make(chan int)
	go suma(10, 10, canal)
	go suma(10, 1, canal)
	resultado1, resultado2 := <-canal, <-canal

	fmt.Println(resultado1, resultado2)
}
