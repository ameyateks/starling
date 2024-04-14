package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
}

func CreateRouter() *mux.Router{
	mux := mux.NewRouter()
	// mux.HandleFunc("/api/user", starlingUser).Methods("GET")
	mux.HandleFunc("/api/accounts", starlingAccount).Methods("GET")
	mux.HandleFunc("/api/transactions", starlingTransactions).Methods("GET")
	mux.HandleFunc("/api/knn", classifyTransaction).Methods("POST", "OPTIONS")
	mux.HandleFunc("/api/category", updateCategoryForTransactionsHandler).Methods("POST", "OPTIONS")
	return mux
}

