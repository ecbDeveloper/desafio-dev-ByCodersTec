package handler

import (
	"desafio/database"
	"desafio/service"
	"encoding/json"
	"net/http"
)

func GetTransactionsDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não dísponível", http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Falha ao conectar com o banco de dados: "+err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	transactions, err := service.GetTransactionsDetails(db)
	if err != nil {
		http.Error(w, "Falha ao recuperar as informações das transações: "+err.Error(), http.StatusInternalServerError)
	}

	arrByteTransaction, err := json.Marshal(*transactions)
	if err != nil {
		http.Error(w, "Falha ao serializar transações", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(arrByteTransaction))
}
