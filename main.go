package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World! Teste")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Servidor HTTP rodando na porta 80")
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
    }
}
