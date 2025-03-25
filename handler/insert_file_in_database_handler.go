package handler

import (
	"bufio"
	"desafio/database"
	"desafio/service"
	"fmt"
	"net/http"
)

func InsertFileInDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não dísponível", http.StatusBadRequest)
		return
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Falha ao conectar com o banco de dados: "+err.Error(), http.StatusInternalServerError)
	}
	defer db.Close()

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Erro ao criar form-data"+err.Error(), http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("cnab")
	if err != nil {
		http.Error(w, "Erro ao obter arquivo"+err.Error(), http.StatusBadRequest)
		return
	}

	if fileHeader.Header == nil || file == nil {
		http.Error(w, "Arquivo não inserido", http.StatusBadRequest)
		return
	}

	fileScanner := bufio.NewScanner(file)

	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
		fileLine := fileScanner.Text()

		err := service.InsertFileInDatabase(db, fileLine)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao inserir dados da linha %d: %s", lineCount, err.Error()), http.StatusBadRequest)
			return
		}
	}
	if err = fileScanner.Err(); err != nil {
		http.Error(w, "Erro ao inserir dados da linha: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Todos as transações forma inseridas no banco com sucesso"}`))
}
