package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const TRANSACTIONS_FILE = "ex3/transactions.json"

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

func FindByID(c *gin.Context) {
	transactions, err := ReadTransactionFromJSON(TRANSACTIONS_FILE)
	if err != nil {
		c.Status(500)
		return
	}

	target := c.Param("id")
	for _, transaction := range transactions {
		if transaction.ID == target {
			c.JSON(200, transaction)
			return
		}
	}
	c.Status(404)
}

func main() {
	router := gin.Default()
	router.GET("/products/:id", FindByID)
	router.Run()
}
