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
		w.Header().Set("Content-Type", "application/json")
		response := &contracts.AddCustomerResponse{}
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Logger.Printf("[AddCustomerHandler]: %v \n", err)

			response.BadRequest("Bad Request", err.Error())
			writeAddCustomerResponse(w, response)
			return
		}

		customer := &contracts.Customer{}
		err = json.Unmarshal(bytes, customer)
		if err != nil {
			logger.Logger.Printf("[AddCustomerHandler]: %v \n", err)

			response.BadRequest("Bad Request", err.Error())
			writeAddCustomerResponse(w, response)
			return
		}

		customerRepository := &repository.Customers{DB: db}
		err = customerRepository.AddCustomer(customer)
		if err != nil {
			logger.Logger.Printf("[AddCustomerHandler]: %v \n", err)

			response.ServerError(err.Error())
			writeAddCustomerResponse(w, response)
			return
		}

		logger.Logger.Println("[AddCustomerHandler]: successfully added a customers.")

		response.Success()
		writeAddCustomerResponse(w, response)
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
