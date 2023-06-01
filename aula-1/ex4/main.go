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

func Filter[T any](arr []T, predicate func(T) bool) (result []T) {
	for _, x := range arr {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return
}

func MatchTransactionWithParameters(c *gin.Context) func(Transaction) bool {
	// TODO...
	return nil
}

func SliceID(c *gin.Context) {
	transactions, err := ReadTransactionFromJSON(TRANSACTIONS_FILE)
	if err != nil {
		c.Status(500)
		return
	}
	filtered := []string{}
	for _, transaction := range transactions {
		filtered = append(filtered, transaction.ID)
	}
	c.JSON(200, filtered)
}

func SliceCurrency(c *gin.Context) {
	transactions, err := ReadTransactionFromJSON(TRANSACTIONS_FILE)
	if err != nil {
		c.Status(500)
		return
	}
	filtered := []string{}
	for _, transaction := range transactions {
		filtered = append(filtered, transaction.Currency)
	}
	c.JSON(200, filtered)
}
func SliceCode(c *gin.Context) {
	transactions, err := ReadTransactionFromJSON(TRANSACTIONS_FILE)
	if err != nil {
		c.Status(500)
		return
	}
	filtered := []string{}
	for _, transaction := range transactions {
		filtered = append(filtered, transaction.Code)
	}
	c.JSON(200, filtered)
}

func main() {
	router := gin.Default()
	group := router.Group("/products")
	{
		group.GET("/id", SliceID)
		group.GET("/code", SliceCode)
		group.GET("/currency", SliceCurrency)
	}
	router.Run()
}
