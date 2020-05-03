package main

import "fmt"

func main() {

	//var asigArray []string
	var asignaturas map[string]string
	fmt.Println(asignaturas)
	//asigArray := []string{}
	//asigArray = append(asigArray, "futbol", "golf")
	//fmt.Println(asigArray)

	deportes := map[string]int{
		"Tiro con arco": 20,
		"Fútbol":        22,
	}
	deportes["Golf"] = 40
	deportes["Bádminton"] = 11

	resultado, ok := deportes["Golf"]
	fmt.Println(resultado, ok)

	delete(deportes, "Bádminton")
	fmt.Println(deportes)

	// recorremos el mapa
	for clave /*valor*/, _ := range deportes {
		fmt.Println(clave)
	}

}
