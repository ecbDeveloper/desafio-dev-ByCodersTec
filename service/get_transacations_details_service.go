package service

import (
	"database/sql"
	"log"
)

type Transaction struct {
	ShopName              string `json:"shop_name"`
	TransactionTotalValue string `json:"transaction_total_value"`
}

func GetTransactionsDetails(db *sql.DB) (*[]Transaction, error) {
	rows, err := db.Query(`
		SELECT 
			t.nome_loja,
			SUM(
				CASE
					WHEN tt.tipo_operacao = '+' THEN t.valor_transacao
					WHEN tt.tipo_operacao = '-' THEN -t.valor_transacao
				END
			) AS total_valor
		FROM transacoes t
		LEFT JOIN tipo_transacao tt
			ON tt.id = t.tipo_transacao
		GROUP BY t.nome_loja;
	`)
	if err != nil {
		log.Println("Erro ao selecionar dados das transações:", err)
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction

		if err := rows.Scan(&transaction.ShopName, &transaction.TransactionTotalValue); err != nil {
			log.Println("Erro ao escanear linhas da tabela transacoes:", err)
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return &transactions, nil
}
