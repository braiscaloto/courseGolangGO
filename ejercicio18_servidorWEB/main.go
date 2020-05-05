package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error creando el servidor")
		fmt.Println(err)
	}
}
