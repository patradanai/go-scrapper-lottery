package main

import (
	"fmt"
	"lottery-web-scrapping/configs"
	"lottery-web-scrapping/routes"
	"net/http"
	"time"
)

func main() {

	// Init Mongo Connection
	if err := configs.ConnectionMongo(); err != nil {
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
}
