package main

import (
	"github.com/alexsukhrin/go-notes/controllers"
	"github.com/alexsukhrin/go-notes/models"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}
