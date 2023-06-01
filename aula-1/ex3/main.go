package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()

	router.GET("/products", func(c *gin.Context) {
		transactions, err := ReadTransactionFromJSON("ex3/transactions.json")
		if err != nil {
			c.Status(500)
			return
		}
		c.JSON(200, gin.H{"data": transactions})
	})

	router.Run()
}
