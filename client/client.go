// client.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Transaction struct {
	Amount float64 `json:"amount"`
}

func main() {
	accountManagerURL := "http://account-manager-service:8080/transaction"
	for {
		amount := rand.Float64()*200 - 100 // Random amount between -100 and +100
		transaction := Transaction{Amount: amount}
		transactionJSON, _ := json.Marshal(transaction)

		resp, err := http.Post(accountManagerURL, "application/json", bytes.NewBuffer(transactionJSON))
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		} else {
			fmt.Println("Transaction sent:", transaction, "Response status:", resp.Status)
			resp.Body.Close()
		}

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}
}
