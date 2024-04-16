package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"starling/routes"
	"starling/services"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	dbenv := os.Getenv("DATABASE_URL")
	fmt.Println(dbenv)

	dba, err := sql.Open("postgres", dbenv)
	driver, err := postgres.WithInstance(dba, &postgres.Config{})
    m, err := migrate.NewWithDatabaseInstance(
        "file:///migrations",
        "postgres", driver)

	if err != nil {
		fmt.Print("failed to run migrations")
		return
	}
	m.Up()
	

	dbHost := os.Getenv("PG_HOST")
	dbUser := os.Getenv("PG_USER")
	dbPass := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_NAME")
	dbPort := os.Getenv("PG_PORT")


	
	db, err := sqlx.Connect("postgres", "user="+dbUser+" dbname="+dbName+" sslmode=disable"+" password="+dbPass+" host="+dbHost+" port="+dbPort)

    m.Up()
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
