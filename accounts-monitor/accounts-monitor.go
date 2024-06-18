package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const csvFile = "transactions.csv"

var mu sync.Mutex

func getBalance() (int, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		return 0, err
	}

	balance := 0
	for _, record := range records {
		amount, err := strconv.Atoi(record[1])
		if err == nil {
			balance += amount
		}
	}

	return balance, nil
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	balance, err := getBalance()
	if err != nil {
		http.Error(w, "Error getting balance", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Current balance: %d", balance)
}

func transactionLogHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open(csvFile)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	r2 := csv.NewReader(file)
	records, err := r2.ReadAll()
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	for _, record := range records {
		fmt.Fprintf(w, "Time: %s, Amount: %s\n", record[0], record[1])
	}
}

func main() {
	http.HandleFunc("/summary", summaryHandler)
	http.HandleFunc("/transactions", transactionLogHandler)
	http.ListenAndServe(":8082", nil)
}
