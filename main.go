package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorld) // Rota principal
	fmt.Println("Servidor iniciado na porta 80...")
	http.ListenAndServe(":80", nil) // Escuta na porta 80
}
