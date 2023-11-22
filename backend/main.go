package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/api/user", starlingUser).Methods("GET")
	mux.HandleFunc("/api/accounts", starlingAccount).Methods("GET")

	http.ListenAndServe(":8080", mux)
}
