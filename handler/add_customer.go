package handler

import (
	"database/sql"
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

		areaRepository := &repository.Areas{DB: db}
		area, err := areaRepository.GetArea(customer.Area)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Logger.Printf("[AddCustomerHandler]: [GetArea]: %v \n", err)
				response.BadRequest("Bad Request", "Area not Found please add it first")
				writeAddCustomerResponse(w, response)
				return
			}

			logger.Logger.Printf("[AddCustomerHandler]: %v \n", err)
			response.ServerError(err.Error())
			writeAddCustomerResponse(w, response)
			return
		}

		customer.AreaID = area.ID

		customerRepository := &repository.Customers{DB: db}
		err = customerRepository.AddCustomer(customer)
		if err != nil {
			logger.Logger.Printf("[AddCustomerHandler]: %v \n", err)

			response.ServerError(err.Error())
			writeAddCustomerResponse(w, response)
			return
		}

		logger.Logger.Println("[AddCustomerHandler]: successfully added given customer.")

		response.Success()
		writeAddCustomerResponse(w, response)
		return
	}
}

func writeAddCustomerResponse(w http.ResponseWriter, response *contracts.AddCustomerResponse) {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Printf("[GetCustomersHandler]: [WriteResponse]: %v \n", err)
		response.ServerError(err.Error())
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseBytes)
}
