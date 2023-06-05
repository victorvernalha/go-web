package main

import (
	"github.com/gin-gonic/gin"
	"github.com/victorvernalha/go-web/webproject/cmd/server/handler"
	"github.com/victorvernalha/go-web/webproject/internal/transactions"
)

func main() {
	tRepo := transactions.CreateInMemoryRepo()
	tService := transactions.DefaultService{Repo: &tRepo}
	tHandler := handler.Transactions{Service: &tService}

	router := gin.Default()
	group := router.Group("/transactions")
	{
		group.GET("/", tHandler.GetAll())
		group.POST("/", tHandler.Add())
		group.PUT("/:id", tHandler.Replace())
	}

	router.Run()
}
