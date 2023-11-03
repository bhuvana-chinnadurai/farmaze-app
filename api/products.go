package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/bhuvana-chinnadurai/farmaze-backend/model"
)

type Products struct {
	products     model.Products
	orders       model.Orders
	procurements model.Procurements
}

// ProcurementResponse represents the response for procured products.
type ProcurementResponse struct {
	ProductID   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Unit        string    `json:"unit"`
}

func NewProducts(products model.Products, procurements model.Procurements) *Products {
	return &Products{
		products:     products,
		procurements: procurements,
	}
}

// Add a new API endpoint to create a product
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	// Parse the request body into the 'product' struct
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the corresponding method in model to add the product
	err = p.products.Create(&product)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create product: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Respond with the created product
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (pr *Products) Edit(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	// Parse the request body into the 'product' struct
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the product ID is valid (exists in the database)
	_, err = pr.products.GetByID(product.ID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Do not allow changing the ID
	// Reset the ID to the one sent in the request, to avoid accidental updates to other products
	product.ID = uuid.MustParse(mux.Vars(r)["id"])

	// Update the product in the database
	err = pr.products.Edit(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (pr *Products) Delete(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["id"]

	id, err := uuid.Parse(productID)
	if err != nil {
		http.Error(w, "Invalid product ID format", http.StatusBadRequest)
		return
	}

	err = pr.products.Delete(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete product: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (pr *Products) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := pr.products.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (pr *Products) GetByID(w http.ResponseWriter, r *http.Request) {
	productIDString := mux.Vars(r)["product_id"]

	productID, err := uuid.Parse(productIDString)
	if err != nil {
		http.Error(w, "Invalid product ID format", http.StatusBadRequest)
		return
	}

	product, err := pr.products.GetByID(productID)
	if err != nil {
		http.Error(w, "Failed to retrieve product by ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
func (pr *Products) ListProcurement(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")

	if dateStr == "" {
		http.Error(w, "Date parameter is required", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	procurements, err := pr.procurements.ListByDate(date)
	if err != nil {
		fmt.Println("error while listing procurements ", err)
		http.Error(w, "Failed to retrieve procurements", http.StatusInternalServerError)
		return
	}

	productQuantities := make(map[uuid.UUID]int)

	fmt.Println("procurement.ProductID", procurements)

	for _, procurement := range procurements {
		fmt.Println("procurement.ProductID", procurement.ProductID)
		productQuantities[procurement.ProductID] += procurement.Quantity
	}

	products, err := pr.products.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}
	productMap := make(map[uuid.UUID]model.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	fmt.Println("productQuantities,", productQuantities)

	// Convert productQuantities map to ProcurementResponse slice
	var response []ProcurementResponse
	for _, procurement := range procurements {
		product, ok := productMap[procurement.ProductID]
		if !ok {
			// Product not found in the map, handle error or skip
			continue
		}

		response = append(response, ProcurementResponse{
			ProductID:   procurement.ProductID,
			ProductName: product.Name,
			Quantity:    procurement.Quantity,
			Unit:        product.Unit,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
