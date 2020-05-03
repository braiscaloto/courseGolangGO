package main

import (
	"fmt"
	"math"
)

type FuncionesCalculadora func(float64, float64) float64

func computar(funcion FuncionesCalculadora, a float64, b float64) float64 {

	return funcion(a, b)

}
func IntroSigno(signo string) FuncionesCalculadora {
	switch signo {
	case "+":
		return func(a, b float64) float64 {
			return a + b
		}
	case "-":
		return func(a, b float64) float64 {
			return a - b
		}
	case "*":
		return func(a, b float64) float64 {
			return a * b
		}
	case "/":
		return func(a, b float64) float64 {
			return a / b
		}
	case "^":
		return math.Pow

	default:
		return func(a, b float64) float64 {
			return a
		}
	}

}

func main() {

	fmt.Println(IntroSigno("+")(10, 10))
	fmt.Println(IntroSigno("-")(10, 10))
	fmt.Println(IntroSigno("/")(10, 10))
	fmt.Println(IntroSigno("*")(10, 10))
	fmt.Println(IntroSigno("")(10, 10))

}
