package server

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"shop/handler"
)

func NewRouter(db *sqlx.DB) *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/homepage").Handler(http.StripPrefix("/homepage", http.FileServer(http.Dir("./frontend/dist"))))

	router.HandleFunc("/area", handler.AddAreaHandler(db)).Methods(http.MethodPut)
	router.HandleFunc("/areas", handler.GetAllAreasHandler(db)).Methods(http.MethodGet)

	router.HandleFunc("/customer", handler.AddCustomerHandler(db)).Methods(http.MethodPut)
	router.HandleFunc("/customer/{id}", handler.GetCustomerHandler(db)).Methods(http.MethodGet)
	router.HandleFunc("/customers", handler.GetCustomersHandler(db)).Methods(http.MethodGet)

	router.HandleFunc("/transaction", handler.AddTransactionHandler(db)).Methods(http.MethodPut)

	router.HandleFunc("/ping", pingHandler).Methods("GET")
	return router
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": \"pong\"}"))
}
