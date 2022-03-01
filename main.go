package main

import (
	"fmt"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/repositories"
	"lottery-web-scrapping/routes"
	"lottery-web-scrapping/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	// Init Mongo Connection
	if err := configs.ConnectionMongo(); err != nil {
		panic(err)
	}

	if err := configs.ConnectionRedis(); err != nil {
		panic(err)
	}

	fmt.Println("Connection to Mongo")

	LotteryRepo := repositories.NewWebScrappingRepository(configs.ClientMongo)
	LotteryService := services.NewWebScrappingService(LotteryRepo)

	LotteryService.FindByDate("https://lottery.kapook.com/check/010365")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routes.InitialRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
