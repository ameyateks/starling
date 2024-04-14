package main

import (
	"net/http"
	"starling/routes"
	"starling/services"

	_ "github.com/lib/pq"
)

func init() {
	services.RunningKnnOnTransactions()
}

func main() {
	mux := routes.CreateRouter()

	http.ListenAndServe(":8080", mux)
}
