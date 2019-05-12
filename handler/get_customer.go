package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"shop/contracts"
	"shop/logger"
	"shop/repository"
	"strconv"
	"strings"
)

func GetCustomerHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		customerRepository := &repository.Customers{DB: db}

		id := strings.TrimPrefix(r.URL.Path, "/customer/")

		response := &contracts.GetCustomerResponse{}
		customerID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			logger.Logger.Printf("[GetCustomerHandler]: %v \n", err)

			response.BadRequest("Bad Request", err.Error())
			writeGetCustomerResponse(w, response)
			return
		}

		customer, err := customerRepository.GetCustomer(customerID)
		if err == sql.ErrNoRows {
			logger.Logger.Printf("[GetCustomersHandler]: [GetCustomer]: %v for customerID: %d \n", err, customerID)

			response.BadRequest("Bad Request", fmt.Sprintf("No Customer found for given ID: %d", customerID))
			writeGetCustomerResponse(w, response)
			return
		}
		if err != nil {
			logger.Logger.Printf("[GetCustomersHandler]: %v \n", err)
			response := contracts.GetCustomerResponse{}
			response.ServerError(err.Error())
			writeGetCustomerResponse(w, &response)
			return
		}

		response.Success(customer)
		logger.Logger.Println("[GetCustomerHandler]: successfully served all customers.")
		writeGetCustomerResponse(w, response)
		return
	}
}

func writeGetCustomerResponse(w http.ResponseWriter, response *contracts.GetCustomerResponse) {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Printf("[GetCustomerHandler]: [WriteResponse]: %v \n", err)
		response.ServerError(err.Error())
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseBytes)
}
