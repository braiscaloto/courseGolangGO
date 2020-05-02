package model

import "fmt"

/* Para declarar una sctruct necesitamos la palabra reservada type,
nombre de la estructura y el tipo.
Para declarar las variables que la componen hay
que poner el nombre y el tipo de variable. */

type motor struct {
	NumeroCilindros int
	Cilindrada      int
}

type Coche struct {
	Motor        *motor
	NumeroRuedas int
	NumeroSerie  string
	roto         bool
}

// En las estructuras es bueno trabajar con punteros

func NewCoche(cilindros, cilindrada, numRuedas int) *Coche {

	// Necesitamos usar un puntero de coche para cambiar la referencia.

	var motorCoche *motor
	if cilindros > 0 {
		motorCoche = &motor{
			NumeroCilindros: cilindros,
			Cilindrada:      cilindrada,
		}
	}
	// Devolvemos el valor al espacio de memoria de Coche, modificándolo.

	return &Coche{
		Motor:        motorCoche,
		NumeroRuedas: numRuedas,
		NumeroSerie:  "ajaiijsias",
		roto:         false,
	}

}

func (c *Coche) Arranca() string {

	if c.roto == true {
		return "No ha podido arrancar, llévalo al mecánico"
	}
	return fmt.Sprintf("El motor ha arrancado correctamente con %v de cilindrada.", c.Motor.Cilindrada)
}

func (c *Coche) IncrementarPotencia(incremento int) {
	c.Motor.Cilindrada += incremento
}

func (c *Coche) PincharRueda() {
	if c.NumeroRuedas == 1 {
		c.NumeroRuedas = 0
		c.roto = true
	} else {
		c.NumeroRuedas--
	}

}
