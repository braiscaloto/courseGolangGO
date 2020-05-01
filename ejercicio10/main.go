package main

import (
	"fmt"

	"github.com/goCurso/cursoGolangGO/ejercicio10/model"
)

// ESTRUCTURAS

func main() {

	motor := model.Motor{
		NumeroCilindros: 3,
		Cilindrada:      1100,
	}

	coche := model.Coche{
		MotorCoche:   motor,
		NumeroRuedas: 4,
		NumeroSerie:  "AZD12344112213NV",
	}
	fmt.Println(motor.Cilindrada)
	fmt.Println(coche.MotorCoche.NumeroCilindros)
	fmt.Println(coche.NumeroSerie)
}
