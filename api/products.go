package api

import (
	"encoding/json"
	"net/http"

	"github.com/bhuvana-chinnadurai/farmaze-backend/model"
)

type Products struct {
	products model.Products
}

func NewProducts(products model.Products) *Products {
	return &Products{
		products: products,
	}
}

func (pr *Products) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := pr.products.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
