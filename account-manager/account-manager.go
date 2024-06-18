package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const csvFile = "transactions.csv"

var balance int

func init() {
	file, err := os.OpenFile(csvFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize balance from file
	r := csv.NewReader(file)
	records, _ := r.ReadAll()
	for _, record := range records {
		amount, _ := strconv.Atoi(record[1])
		balance += amount
	}
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	ttype := r.URL.Query().Get("type")
	amountStr := r.URL.Query().Get("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	if ttype == "withdrawal" {
		amount = -amount
	}

	balance += amount

	file, err := os.OpenFile(csvFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	wtr := csv.NewWriter(file)
	defer wtr.Flush()
	wtr.Write([]string{time.Now().Format(time.RFC3339), strconv.Itoa(amount)})

	fmt.Fprintf(w, "Transaction successful, new balance: %d", balance)
}

func main() {
	http.HandleFunc("/transaction", transactionHandler)
	http.ListenAndServe(":8081", nil)
}
