package handler

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shop/contracts"
	"shop/repository"
)

func AddMerchantHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		amrb := &contracts.AddMerchantRequest{}
		err = json.Unmarshal(bytes, amrb)
		if err != nil {
			panic(err)
		}

		merchantRepository := &repository.Merchants{DB: db}
		err = merchantRepository.AddMerchant(amrb.Area, amrb.Name, amrb.Phone)
		if err != nil {
			panic(err)
		}
	}
}
