package transaction

import (
	"time"
)

// Transaction : struct criada e quando for json altera o estilo
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"creadted_at"`
}

// Transactions : Transaction Sendo chamada.
type Transactions []Transaction
