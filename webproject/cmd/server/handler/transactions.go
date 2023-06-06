package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victorvernalha/go-web/webproject/cmd/server/middleware"
	"github.com/victorvernalha/go-web/webproject/internal/transactions"
)

type Transactions struct {
	Service transactions.Service
}

type AddRequest struct {
	Code     string    `binding:"required" json:"transactionCode"`
	Currency string    `binding:"required" json:"currency"`
	Amount   float64   `binding:"required" json:"amount"`
	Sender   string    `binding:"required" json:"sender"`
	Receiver string    `binding:"required" json:"receiver"`
	Date     time.Time `binding:"required" json:"date"`
}

type UpdateRequest struct {
	Code   string  `binding:"required" json:"transactionCode"`
	Amount float64 `binding:"required" json:"amount"`
}

func (h *Transactions) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := middleware.ParsedRequest[AddRequest](c)

		t, err := h.Service.Create(req.Code, req.Currency, req.Sender, req.Receiver, req.Amount, req.Date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func (h *Transactions) Replace() gin.HandlerFunc {
	return func(c *gin.Context) {
		strId, _ := c.Params.Get("id")
		id, err := strconv.ParseInt(strId, 10, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Path parameter must be an integer"})
			return
		}

		req := middleware.ParsedRequest[AddRequest](c)
		t := transactions.Transaction{
			ID:       int(id),
			Code:     req.Code,
			Currency: req.Currency,
			Amount:   req.Amount,
			Sender:   req.Sender,
			Receiver: req.Receiver,
			Date:     req.Date,
		}

		t, err = h.Service.Replace(t)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": strings.Split(err.Error(), "\n")})
			return
		}
		c.JSON(http.StatusOK, t)
	}
}

func (h *Transactions) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		strId, _ := c.Params.Get("id")
		id, err := strconv.ParseInt(strId, 10, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Path parameter must be an integer"})
			return
		}
		err = h.Service.Delete(int(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transaction does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}

func (h *Transactions) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		strId, _ := c.Params.Get("id")
		id, err := strconv.ParseInt(strId, 10, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Path parameter must be an integer"})
			return
		}

		req := middleware.ParsedRequest[UpdateRequest](c)

		t, err := h.Service.Update(int(id), req.Code, req.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, t)
	}
}
