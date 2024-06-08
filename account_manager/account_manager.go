// account_manager.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Transaction struct {
	Amount float64 `json:"amount"`
}

type Account struct {
	sync.Mutex
	Balance      float64
	Transactions []string
}

var account = Account{}

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	var t Transaction
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account.Lock()
	account.Balance += t.Amount
	account.Transactions = append(account.Transactions, fmt.Sprintf("Time: %s, Amount: %.2f", time.Now().Format(time.RFC3339), t.Amount))
	account.Unlock()

	log.Printf("Transaction processed: %+v", t)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/transaction", handleTransaction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
