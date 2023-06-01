package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Transaction struct {
	ID       string
	Code     string `json:"transactionCode"`
	Currency string
	Amount   float64
	Sender   string
	Receiver string
	Date     time.Time
}

func ReadTransactionFromJSON(file string) (transactions []Transaction, err error) {
	data, err := os.ReadFile("ex1/transactions.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &transactions)
	if err != nil {
		return
	}
	return
}

func main() {
	transactions, err := ReadTransactionFromJSON("ex1/transactions.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(transactions)
}
