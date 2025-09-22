package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server working properly.")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("🚀 Listening on localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
