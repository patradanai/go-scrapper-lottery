package main

import (
	"fmt"
	"lottery-web-scrapping/api/routes"
	"lottery-web-scrapping/driver"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	// Init Mongo Connection
	if err := driver.ConnectionMongo(); err != nil {
		panic(err)
	}

	if err := driver.ConnectionRedis(); err != nil {
		panic(err)
	}

	fmt.Println("Connection to Mongo")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routes.InitialRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sig

	os.Exit(0)
}
