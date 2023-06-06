package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/victorvernalha/go-web/webproject/cmd/server/handler"
	"github.com/victorvernalha/go-web/webproject/cmd/server/middleware"
	"github.com/victorvernalha/go-web/webproject/internal/transactions"
)

const API_KEY_NAME = "API_KEY"
const API_KEY_HEADER = "authorization"

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Could not load .env file")
	}

	tRepo := transactions.CreateInMemoryRepo()
	tService := transactions.DefaultService{Repo: &tRepo}
	tHandler := handler.Transactions{Service: &tService}

	tokenValidator := middleware.TokenValidator(os.Getenv(API_KEY_NAME), API_KEY_HEADER)

	router := gin.Default()
	group := router.Group("/transactions")
	{
		group.Use(tokenValidator)
		group.GET("/", tHandler.GetAll())
		group.POST("/", tHandler.Add())
		group.PUT("/:id", tHandler.Replace())
		group.DELETE("/:id", tHandler.Delete())
		group.PATCH("/:id", tHandler.Update())
	}

	router.Run()
}
