package main

import (
	"encoding/json"
	"net/url"
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

func ToMap[T any](s T) (mapped map[string]any) {
	serialized, _ := json.Marshal(s)
	json.Unmarshal(serialized, &mapped)
	return
}

func GetFilterPredicate(params *url.Values) func(Transaction) bool {
	return func(t Transaction) bool {
		tMap := ToMap[Transaction](t)

		matches := true
		for key, value := range *params {
			tVal, ok := tMap[key]
			if ok {
				matches = matches && (tVal == value[0])
			}
		}
		return matches
	}
}

func FilterTransactions(c *gin.Context) {
	transactions, err := ReadTransactionFromJSON(TRANSACTIONS_FILE)
	if err != nil {
		c.Status(500)
		return
	}

	params := c.Request.URL.Query()
	filterPredicate := GetFilterPredicate(&params)
	transactions = Filter[Transaction](transactions, filterPredicate)

	c.JSON(200, gin.H{"data": transactions})
}

func main() {
	router := gin.Default()
	router.GET("/transactions", FilterTransactions)
	router.Run()
}
