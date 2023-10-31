package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/bhuvana-chinnadurai/farmaze-backend/model"
)

type Order struct {
	orderRepo   model.Orders
	productRepo model.Products
	clientRepo  model.B2BClients
	orderStatus []model.OrderStatus
}

type OrderStatus string

const (
	Ordered OrderStatus = "ordered"
	//Dispatched OrderStatus = "dispatched"
	//Delivered  OrderStatus = "delivered"
)

// NewOrder creates a new Order instance with the given OrderRepository.
func NewOrder(orderRepo model.Orders, productRepo model.Products, clientRepo model.B2BClients, orderStatuses []model.OrderStatus) *Order {
	return &Order{orderRepo: orderRepo, productRepo: productRepo, clientRepo: clientRepo, orderStatus: orderStatuses}
}

func (o *Order) areValidProductIDs(orderProducts []model.OrderProduct, products []model.Product) bool {
	for _, orderProduct := range orderProducts {
		found := false
		for _, product := range products {
			if product.ID.String() == orderProduct.ProductID.String() {
				fmt.Println("found the product id", orderProduct.ProductID)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("did not found the product id: ", orderProduct.ProductID.String())
			return false
		}
	}
	return true
}

func fetchOrderStatusId(orderStatuses []model.OrderStatus, desiredStatus OrderStatus) uuid.UUID {
	for _, status := range orderStatuses {
		if status.Name == string(desiredStatus) {
			fmt.Println("check order status", status.ID)
			return status.ID
		}
	}

	return uuid.Nil
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	var order model.Order

	// Parse the request body into the 'order' struct
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get all products
	products, err := o.productRepo.GetAllProducts()
	if err != nil {
		// Handle the error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Check if the product ID is valid (exists in the database)
	if !o.areValidProductIDs(order.Products, products) {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	client, err := o.clientRepo.GetClientByID(order.ClientID)
	if err != nil {
		if _, ok := err.(model.NotFoundError); ok {
			http.Error(w, "Client not found", http.StatusNotFound)
			return
		}
		// Handle other types of errors
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("inputs are correct")
	// Client found, continue with order creation logic
	order.ClientID = client.ID

	// Generate a new order ID
	orderID := uuid.New()

	fmt.Println("o.orderStatus", o.orderStatus)
	// Set the order ID and status (defaulted to "ordered")
	order.ID = orderID
	orderStatusId := fetchOrderStatusId(o.orderStatus, Ordered)
	if orderStatusId == uuid.Nil {
		http.Error(w, "Invalid order status ID", http.StatusBadRequest)
		return
	}
	fmt.Println("orderStatusId is ", orderStatusId)
	order.Status = orderStatusId

	// Create the order in the database
	err = o.orderRepo.Create(&order)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create order: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// Respond with the created order
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// List retrieves all orders from the database.
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	orders, err := o.orderRepo.List()
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// GetByClientID retrieves orders specific to a given client.
func (o *Order) GetByClientID(w http.ResponseWriter, r *http.Request) {
	clientIDString := mux.Vars(r)["client_id"]

	clientID, err := uuid.Parse(clientIDString)
	if err != nil {
		http.Error(w, "Invalid client ID format", http.StatusBadRequest)
		return
	}

	orders, err := o.orderRepo.GetByClientID(clientID)
	if err != nil {
		http.Error(w, "Failed to retrieve orders by client ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
