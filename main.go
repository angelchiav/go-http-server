package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Si ves esto el servidor funciona.")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Escuchando en localhost:8080/")
	http.ListenAndServe(":8080", nil)
}