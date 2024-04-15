package utils

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func ConnectToDB() *sqlx.DB {
	dbHost := os.Getenv("PG_HOST")
	dbUser := os.Getenv("PG_USER")
	dbPass := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_NAME")

	db, err := sqlx.Connect("postgres", "user="+dbUser+" dbname="+dbName+" sslmode=disable"+" password="+dbPass+" host="+dbHost+" port=5433")
	if err != nil {
		log.Fatalln(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	return db
}