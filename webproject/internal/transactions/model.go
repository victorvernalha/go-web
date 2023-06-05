package transactions

import "time"

type Transaction struct {
	ID       int
	Code     string    `json:"transactionCode"`
	Currency string    `json:"currency"`
	Amount   float64   `json:"amount"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	Date     time.Time `json:"date"`
}
