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

func AddAreaHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := &contracts.AddAreaResponse{}
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Logger.Printf("[AddAreaHandler]: %v \n", err)

			response.BadRequest("Bad Request", err.Error())
			writeAddAreaResponse(w, response)
			return
		}

		area := &contracts.Area{}
		err = json.Unmarshal(bytes, area)
		if err != nil {
			logger.Logger.Printf("[AddAreaHandler]: %v \n", err)

			response.BadRequest("Bad Request", err.Error())
			writeAddAreaResponse(w, response)
			return
		}

		areaRepository := &repository.Areas{DB: db}
		err = areaRepository.AddArea(area.Name)
		if err != nil {
			logger.Logger.Printf("[AddAreaHandler]: %v \n", err)

			response.ServerError(err.Error())
			writeAddAreaResponse(w, response)
			return
		}

		logger.Logger.Println("[AddAreaHandler]: successfully added given area.")

		response.Success()
		writeAddAreaResponse(w, response)
		return
	}
}

func writeAddAreaResponse(w http.ResponseWriter, response *contracts.AddAreaResponse) {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Printf("[GetAreasHandler]: [WriteResponse]: %v \n", err)
		response.ServerError(err.Error())
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseBytes)
}
