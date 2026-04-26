package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me a valid id", 400)
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "data not found", 200)
}
