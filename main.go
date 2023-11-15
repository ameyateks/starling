package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")

	if exists {
		fmt.Println(accessToken)
	}

	mux := mux.NewRouter()
	fmt.Println("accessToken:", os.Getenv("ACCESS_TOKEN"))
	mux.HandleFunc("/api/case/{caseId}/outgoingsByRegion", outgoingsByCaseId).Methods("GET")

	http.ListenAndServe(":8080", mux)
}
