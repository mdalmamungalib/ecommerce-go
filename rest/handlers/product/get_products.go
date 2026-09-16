package handlers

import (
	"net/http"
	"ecommerce/database"
	"ecommerce/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Plz give me GET request", 400)
		return
	}

	util.SendData(w, database.List(), 200)
}
