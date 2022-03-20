package main

import (
	"fmt"
	"lottery-web-scrapping/api/handler"
	"lottery-web-scrapping/api/router"
	"lottery-web-scrapping/driver"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	schedule "lottery-web-scrapping/internal/scheduler"
)

func main() {
	gin.SetMode(gin.DebugMode)

	// Init Mongo Connection
	c := driver.ConnectionMongo()
	h := handler.NewHanlder(c)
	r := router.InitialRouter(h)

	if err := driver.ConnectionRedis(); err != nil {
		panic(err)
	}

	fmt.Println("Connection to Mongo")

	// Scheduler
	schedule.Handler()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
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
