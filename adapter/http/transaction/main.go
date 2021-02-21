package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/treinalinux/dio-expert-session-finance/model/transaction"
)

// GetTransactions :
func GetTransactions(w http.ResponseWriter, r *http.Request) {
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

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.8,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

// CreateATransactions :
func CreateATransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
