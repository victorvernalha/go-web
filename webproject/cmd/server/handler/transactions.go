package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/victorvernalha/go-web/pkg/middleware"
	"github.com/victorvernalha/go-web/pkg/responses"
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

// Add godoc
//
//	@Summary	Add new transaction
//	@Tags		Transactions
//	@Accept		json
//	@Produce	json
//	@Param		authorization	header		string				true	"Authentication token"
//	@Param		transaction		body		AddRequest			true	"Transaction to be added"
//	@Success	200				{object}	responses.Response	"Returns updated transaction"
//	@Failure	500				{object}	responses.Response	"Could not save transaction"
//	@Router		/transactions [post]
func (h *Transactions) Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := middleware.ParsedRequest[AddRequest](c)

		t, err := h.Service.Create(req.Code, req.Currency, req.Sender, req.Receiver, req.Amount, req.Date)
		if err != nil {
			responses.Error(c, http.StatusInternalServerError, err)
			return
		}

		responses.Success(c, http.StatusCreated, t)
	}
}

// GetAll godoc
//
//	@Summary	Get all transactions
//	@Tags		Transactions
//	@Accept		json
//	@Produce	json
//	@Param		authorization	header		string	true	"Authentication token"
//	@Success	200				{object}	responses.Response
//	@Failure	500				{object}	responses.Response	"Could not fetch transactions"
//	@Router		/transactions [get]
func (h *Transactions) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ts, err := h.Service.GetAll()
		if err != nil {
			responses.Error(c, http.StatusInternalServerError, err)
			return
		}
		responses.Success(c, http.StatusCreated, ts)
	}
}

// Replace godoc
//
//	@Summary	Replaces given transaction
//	@Tags		Transactions
//	@Accept		json
//	@Produce	json
//	@Param		authorization	header		string				true	"Authentication token"
//	@Param		id				path		int					true	"Transaction ID"
//	@Param		trnsaction		body		AddRequest			true	"Updated transaction"
//	@Success	200				{object}	responses.Response	"Returns updated transaction"
//	@Failure	400				{object}	responses.Response	"Path parameter is not an int"
//	@Failure	400				{object}	responses.Response	"Missing or invalid transaction parameters"
//	@Failure	500				{object}	responses.Response	"Given ID does not exist"
//	@Failure	500				{object}	responses.Response	"Could not replace transaction"
//	@Router		/transactions/:id [put]
func (h *Transactions) Replace() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := validateIntPathParam(c, "id")
		if err != nil {
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
			responses.Error(c, http.StatusInternalServerError, err)
			return
		}
		responses.Success(c, http.StatusOK, t)
	}
}

// Delete godoc
//
//	@Summary	Get all transactions
//	@Tags		Transactions
//	@Accept		json
//	@Produce	json
//	@Param		authorization	header		string	true	"Authentication token"
//	@Param		id				path		int		true	"Transaction ID"
//	@Success	200				{object}	responses.Response
//	@Failure	400				{object}	responses.Response	"Path parameter is not an int"
//	@Failure	404				{object}	responses.Response	"Could not find given transaction"
//	@Router		/transactions [delete]
func (h *Transactions) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := validateIntPathParam(c, "id")
		if err != nil {
			return
		}

		err = h.Service.Delete(int(id))
		if err != nil {
			responses.Error(c, http.StatusNotFound, err)
			return
		}
		responses.SuccessNoData(c, http.StatusOK)
	}
}

// Update godoc
//
//	@Summary	Updates given transaction
//	@Tags		Transactions
//	@Accept		json
//	@Produce	json
//	@Param		authorization	header		string				true	"Authentication token"
//	@Param		id				path		int					true	"Transaction ID"
//	@Param		transactionCode	body		string				true	"New transaction code"
//	@Param		amount			body		number				true	"New transaction amount"
//	@Success	200				{object}	responses.Response	"Returns updated transaction"
//	@Failure	400				{object}	responses.Response	"Path parameter is not an int"
//	@Failure	400				{object}	responses.Response	"Missing or invalid transaction parameters"
//	@Failure	400				{object}	responses.Response	"Given ID does not exist"
//	@Failure	400				{object}	responses.Response	"Could not replace transaction"
//	@Router		/transactions/:id [patch]
func (h *Transactions) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := validateIntPathParam(c, "id")
		if err != nil {
			return
		}

		req := middleware.ParsedRequest[UpdateRequest](c)

		t, err := h.Service.Update(int(id), req.Code, req.Amount)
		if err != nil {
			responses.Error(c, http.StatusBadRequest, err)
			return
		}

		responses.Success(c, http.StatusOK, t)
	}
}

func validateIntPathParam(c *gin.Context, param string) (val int, err error) {
	strParam, _ := c.Params.Get(param)
	val64, err := strconv.ParseInt(strParam, 10, 0)
	if err != nil {
		err = fmt.Errorf("invalid path parameter %s; expected int", strParam)
		responses.Error(c, http.StatusBadRequest, err)
		return
	}
	val = int(val64)
	return
}
