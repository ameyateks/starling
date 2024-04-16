package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
}

func CreateRouter(db *sqlx.DB) *mux.Router{
	mux := mux.NewRouter()
	getStarlingTransaction := wrapStarlingTransactionsHandler(db)
	
	mux.HandleFunc("/api/accounts", starlingAccount).Methods(http.MethodGet)
	mux.HandleFunc("/api/transactions", getStarlingTransaction).Methods(http.MethodGet)
	mux.HandleFunc("/api/knn", classifyTransaction).Methods(http.MethodPost, http.MethodOptions)
	mux.HandleFunc("/api/category", updateCategoryForTransactionsHandler).Methods(http.MethodPost, http.MethodOptions)
	return mux
}

