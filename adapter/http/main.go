package http

import (
	"net/http"

	"github.com/treinalinux/dio-expert-session-finance/adapter/http/transaction"
)

// Init : Inicia o http
func Init() {

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransactions)

	http.ListenAndServe(":8080", nil)
}