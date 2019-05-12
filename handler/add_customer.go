package handler

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
	"shop/contracts"
	"shop/logger"
	"shop/repository"
)

func AddCustomerHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		customer := &contracts.Customer{}
		err = json.Unmarshal(bytes, customer)
		if err != nil {
			logger.Logger.Printf("[AddCustomerHandler]: %v \n", err)
			response := contracts.AddCustomerResponse{
				StatusCode: http.StatusBadRequest,
				Error: &contracts.Error{
					Title:   "Bad Request",
					Message: err.Error(),
				},
			}

			writeAddCustomerResponse(w, &response)
			return
		}

		customerRepository := &repository.Customers{DB: db}
		err = customerRepository.AddCustomer(customer)
		if err != nil {
			logger.Logger.Printf("[AddCustomerHandler]: %v \n", err)
			response := contracts.AddCustomerResponse{
				StatusCode: http.StatusInternalServerError,
				Error: &contracts.Error{
					Title:   "Internal Server Error",
					Message: err.Error(),
				},
			}

			writeAddCustomerResponse(w, &response)
			return
		}

		response := contracts.AddCustomerResponse{
			StatusCode: http.StatusOK,
			Data:       "Successfully added a customers.",
		}

		logger.Logger.Println("[AddCustomerHandler]: successfully added a customers.")
		writeAddCustomerResponse(w, &response)
		return
	}
}

func writeAddCustomerResponse(w http.ResponseWriter, response *contracts.AddCustomerResponse) {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Printf("[GetCustomersHandler]: [WriteResponse]: %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseBytes)
}
