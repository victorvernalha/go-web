package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victorvernalha/go-web/webproject/internal/transactions"
)

type Transactions struct {
	Service transactions.Service
}

type addTransactionRequest struct {
	Code     string    `json:"transactionCode"`
	Currency string    `json:"currency"`
	Amount   float64   `json:"amount"`
	Sender   string    `json:"sender"`
	Receiver string    `json:"receiver"`
	Date     time.Time `json:"date"`
}

func (h *Transactions) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req addTransactionRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		t, err := h.Service.Create(req.Code, req.Currency, req.Sender, req.Receiver, req.Amount, req.Date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, t)
	}
}

func (h *Transactions) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ts, err := h.Service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusCreated, ts)
	}
}
