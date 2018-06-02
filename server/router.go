package server

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"shop/handler"
)

func NewRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/homepage").Handler(http.FileServer(http.Dir("./frontend/dist")))

	addMerchantHandler := handler.AddMerchantHandler(db)

	router.HandleFunc("/merchants", addMerchantHandler).Methods(http.MethodPut)

	router.HandleFunc("/ping", pingHandler).Methods("GET")
	return router
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": \"pong\"}"))
}
