package main

import (
	"net/http"
	"shop/server"
)

func main() {
	router := server.NewRouter()
	http.ListenAndServe(":8080", router)
}