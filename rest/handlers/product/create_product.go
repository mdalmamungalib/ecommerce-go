package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	//Parse jwt
	//parse header and payload or claims
	//hmac-sha-256 algorithm -> has hmac(header, payload, secret key) -> hash hmac(header, payload, secret key)
	//parse signature part from the jwt
	//if the signature and hash is same => forward to create products
	// otherwise 401 status code with Unauthorized

	if r.Method != "POST" {
		http.Error(w, "Plz give me POST request", 400)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	util.SendData(w, createdProduct, 201)
}
