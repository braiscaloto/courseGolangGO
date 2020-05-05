package main

import "fmt"

type sumaElements struct {
	a int
	b int
}

func sumaProceso(sumasElementos chan sumaElements, resultados chan int, quit chan int) {

	for {
		select {
		case sumaArealizar := <-sumasElementos:
			resultados <- sumaArealizar.a + sumaArealizar.b

		case <-quit:
			return
		}
	}
}
func main() {

	canalElementos := make(chan sumaElements)
	canalResultados := make(chan int, 50)
	canalQuit := make(chan int)
	go sumaProceso(canalElementos, canalResultados, canalQuit)
	canalElementos <- sumaElements{1, 2}

	i := 0
FR:
	for {
		select {
		case resultado := <-canalResultados:
			i++
			fmt.Println(resultado)
			if i == 100000 {
				canalQuit <- 1
				break FR
			} else {
				canalElementos <- sumaElements{i, 2}
			}

		}
	}

}
