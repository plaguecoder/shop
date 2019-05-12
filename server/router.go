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

	addCustomerHandler := handler.AddCustomerHandler(db)
	getCustomerHandler := handler.GetCustomerHandler(db)
	getCustomersHandler := handler.GetCustomersHandler(db)

	router.HandleFunc("/customer", addCustomerHandler).Methods(http.MethodPut)
	router.HandleFunc("/customer/{id}", getCustomerHandler).Methods(http.MethodGet)
	router.HandleFunc("/customers", getCustomersHandler).Methods(http.MethodGet)

	router.HandleFunc("/ping", pingHandler).Methods("GET")
	return router
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"success\": \"pong\"}"))
}
