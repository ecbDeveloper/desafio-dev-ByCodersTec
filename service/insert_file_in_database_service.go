package service

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func InsertFileInDatabase(db *sql.DB, fileLine string) error {
	transactionType := fileLine[0:1]
	transactionDate := fileLine[1:9]
	transactionValue := fileLine[9:19]
	CPF := fileLine[19:30]
	card := fileLine[30:42]
	transactionHour := fileLine[42:48]
	shopOwner := strings.TrimSpace(fileLine[48:62])
	shopName := strings.TrimSpace(fileLine[62:])

	transactionDateTime, err := time.Parse("20060102", transactionDate)
	if err != nil {
		fmt.Println("Erro ao converter string para data:", err)
		return err
	}

	transactionHourTime, err := time.Parse("150405", transactionHour)
	if err != nil {
		fmt.Println("Erro ao converter string para hora:", err)
		return err
	}

	transactionValueFloat, err := strconv.ParseFloat(transactionValue, 64)
	if err != nil {
		fmt.Println("Erro ao converter:", err)
		return err
	}

	_, err = db.Exec(`
		INSERT INTO transacoes
		(tipo_transacao, data_transacao, valor_transacao, 
		cpf, cartao, hora_transacao, dono_loja, nome_loja)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, transactionType, transactionDateTime, transactionValueFloat/100, CPF, card, transactionHourTime, shopOwner, shopName)
	if err != nil {
		log.Println("Failed to insert line on database:", err)
		return err
	}

	return nil
}
