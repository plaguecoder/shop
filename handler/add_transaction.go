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

func AddTransactionHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := &contracts.AddTransactionResponse{}
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Logger.Printf("[AddTransactionHandler]: %v \n", err)

			response.BadRequest("Bad Request", err.Error())
			writeAddTransactionResponse(w, response)
			return
		}

		transaction := &contracts.Transaction{}
		err = json.Unmarshal(bytes, transaction)
		if err != nil {
			logger.Logger.Printf("[AddTransactionHandler]: %v \n", err)

			response.BadRequest("Bad Request", err.Error())
			writeAddTransactionResponse(w, response)
			return
		}

		transactionRepository := &repository.Transactions{DB: db}
		err = transactionRepository.AddTransaction(transaction)
		if err != nil {
			logger.Logger.Printf("[AddTransactionHandler]: %v \n", err)

			response.ServerError(err.Error())
			writeAddTransactionResponse(w, response)
			return
		}

		logger.Logger.Printf("[AddTransactionHandler]: successfully added transaction for given customer: %d \n", transaction.CustomerID)

		response.Success()
		writeAddTransactionResponse(w, response)
		return
	}
}

func writeAddTransactionResponse(w http.ResponseWriter, response *contracts.AddTransactionResponse) {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Printf("[GetCustomersHandler]: [WriteResponse]: %v \n", err)
		response.ServerError(err.Error())
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseBytes)
}
