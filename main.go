package main

import (
	"desafio/handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Servidor iniciado em http://localhost:3000 �")

	http.HandleFunc("/transactions/upload", handler.InsertFileInDatabaseHandler)
	http.HandleFunc("/transactions/details", handler.GetTransactionsDetailsHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
