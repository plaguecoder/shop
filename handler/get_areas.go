package handler

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"net/http"
	"shop/contracts"
	"shop/logger"
	"shop/repository"
)

func GetAllAreasHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := &contracts.GetAllAreasResponse{}

		areaRepository := &repository.Areas{DB: db}
		names, err := areaRepository.GetAllAreas()
		if err != nil {
			logger.Logger.Printf("[AddAreaHandler]: %v \n", err)

			response.ServerError(err.Error())
			writeGetAllAreasResponse(w, response)
			return
		}

		logger.Logger.Println("[AddAreaHandler]: successfully returned all area.")

		response.Success(names)
		writeGetAllAreasResponse(w, response)
		return
	}
}

func writeGetAllAreasResponse(w http.ResponseWriter, response *contracts.GetAllAreasResponse) {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Printf("[GetAreasHandler]: [WriteResponse]: %v \n", err)
		response.ServerError(err.Error())
		return
	}

	w.WriteHeader(response.StatusCode)
	w.Write(responseBytes)
}
