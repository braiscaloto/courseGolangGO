package model2

/* Para declarar una sctruct necesitamos la palabra reservada type,
nombre de la estructura y el tipo.
Para declarar las variables que la componen hay
que poner el nombre y el tipo de variable. */

type Motor struct {
	NumeroCilindros int
	Cilindrada      int
}

type Coche struct {
	MotorCoche   Motor
	NumeroRuedas int
	NumeroSerie  string
}
type Coche2 struct {
	MotorCoche   Motor
	NumeroRuedas int
	NumeroSerie  string
}
