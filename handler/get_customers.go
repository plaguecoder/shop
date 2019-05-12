package handler

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"net/http"
	"shop/contracts"
	"shop/logger"
	"shop/repository"
)

func GetCustomersHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		customerRepository := &repository.Customers{DB: db}
		customers, err := customerRepository.GetAllCustomers()
		if err != nil {
			logger.Logger.Printf("[GetCustomersHandler]: %v \n", err)
			response := contracts.GetCustomersResponse{
				StatusCode: http.StatusInternalServerError,
				Error: &contracts.Error{
					Title:   "Internal Server Error",
					Message: err.Error(),
				},
			}

			writeResponse(w, &response)
			return
		}

		response := contracts.GetCustomersResponse{
			StatusCode: http.StatusOK,
			Data:       customers,
		}

		logger.Logger.Println("[GetCustomersHandler]: successfully served all customers.")
		writeResponse(w, &response)
		return
	}
}

func writeResponse(w http.ResponseWriter, response *contracts.GetCustomersResponse) {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Printf("[GetCustomersHandler]: [WriteResponse]: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseBytes)
}
