package main

import (
	"log"
	"net/http"
	"starling/routes"
	"starling/services"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	services.RunningKnnOnTransactions()
}

func main() {
	mux := routes.CreateRouter()

	http.ListenAndServe(":8080", mux)
}
