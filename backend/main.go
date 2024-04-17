package main

import (
	"log"
	"net/http"
	"os"
	"starling/routes"
	"starling/services"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
func main() {
	dbHost := os.Getenv("PG_HOST")
	dbUser := os.Getenv("PG_USER")
	dbPass := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_NAME")
	dbPort := os.Getenv("PG_PORT")
	db, err := sqlx.Connect("postgres", "user="+dbUser+" dbname="+dbName+" sslmode=disable"+" password="+dbPass+" host="+dbHost+" port="+dbPort)
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	knnErr := services.RunningKnnOnTransactions(db)
	if knnErr != nil {
		panic(err)
	}

	router := routes.CreateRouter(db)
	router.Use(routes.CorsMiddleware)
	router.Use(routes.JsonResponseMiddleware)

	http.ListenAndServe(":8080", router)
}