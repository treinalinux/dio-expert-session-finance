package http

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/treinalinux/dio-expert-session-finance/adapter/http/actuator"
	"github.com/treinalinux/dio-expert-session-finance/adapter/http/transaction"
)

// Init : Inicia o http
func Init() {

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransactions)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
