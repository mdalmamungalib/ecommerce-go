package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

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

	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", 400)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdProduct, http.StatusCreated)
}
