package main

// ESTRUCTURAS COMO OBJETOS
import (
	"fmt"

	"github.com/goCurso/cursoGolangGO/ejercicio10_estructurasPunteros/model"
)

func main() {

	coche := model.NewCoche(2, 600, 2)

	coche.IncrementarPotencia(10)
	coche.PincharRueda()
	//coche.PincharRueda() // Con esto el coche no arranca

	fmt.Println(coche.Arranca())

}
