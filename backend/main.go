package main

import (
	"log"
	"net/http"
	"os"
	"starling/routes"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

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

	knnErr := routes.RunningKnnOnTransactions(db)
	if knnErr != nil {
		panic(err)
	}

	mux := routes.CreateRouter(db)

	http.ListenAndServe(":8080", mux)
}
