package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/transactions", getTransactions)
	http.HandleFunc("/transactions/create", createATransactions)

	http.ListenAndServe(":8080", nil)
}

// Transaction : struct criada e quando for json altera o estilo
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"creadted_at"`
}

// Transactions :
type Transactions []Transaction

// getTransactions :
func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	/*
		layout : formato do pacote time:
			ano-mes-diaMarcadoHora:Minuto:Segundo

		Mais detalhes em: https://golang.org/pkg/time/
	*/
	var layout = "2006-01-02T15:04:05"

	salaryReceived, _ := time.Parse(layout, "2021-02-20T22:30:26")

	var transactions = Transactions{
		Transaction{
			Title:     "Salário",
			Amount:    1200.8,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

// createATransactions :
func createATransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
