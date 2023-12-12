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

type Order struct {
	orderRepo    model.Orders
	productRepo  model.Products
	clientRepo   model.B2BClients
	orderStatus  []model.OrderStatus
	procurements model.Procurements
}

// ProductRequest represents the product ID and quantity in the order request
type ProductRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

// ProductRequest represents the product ID and quantity in the order request
type ProductResponse struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}

// CreateOrderRequest struct for creating an order
type CreateOrderRequest struct {
	ClientID   uuid.UUID        `json:"client_id"`
	Products   []ProductRequest `json:"products"`
	TotalPrice float64          `json:"total_price"`
	CreatedAt  time.Time        `json:"created_at"`
}

// CreateOrderResponse struct for the created order
type CreateOrderResponse struct {
	ID         uuid.UUID         `json:"id"`
	ClientID   uuid.UUID         `json:"client_id"`
	Products   []ProductResponse `json:"products"`
	TotalPrice float64           `json:"total_price"`
	CreatedAt  time.Time         `json:"created_at"`
	Status     OrderStatus       `json:"status"`
}

type OrderStatus string

const (
	Ordered OrderStatus = "ordered"
	//Dispatched OrderStatus = "dispatched"
	//Delivered  OrderStatus = "delivered"
)

// NewOrder creates a new Order instance with the given OrderRepository.
func NewOrder(orderRepo model.Orders, productRepo model.Products, clientRepo model.B2BClients, procurements model.Procurements, orderStatuses []model.OrderStatus) *Order {
	return &Order{orderRepo: orderRepo, productRepo: productRepo, clientRepo: clientRepo, procurements: procurements, orderStatus: orderStatuses}
}

func (o *Order) areValidProductIDs(orderProducts []ProductRequest, products []model.Product) bool {
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
			return status.ID
		}
	}

	return uuid.Nil
}
func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("order create request")
	var request CreateOrderRequest

	// Parse the request body into the 'request' struct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("error while decoding request:", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get all products
	products, err := o.productRepo.GetAll()
	if err != nil {
		// Handle the error
		fmt.Println("error while getting all products:", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Check if the product IDs are valid (exist in the database)
	if !o.areValidProductIDs(request.Products, products) {
		fmt.Println("invalid product ID:", err.Error())
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	client, err := o.clientRepo.GetClientByID(request.ClientID)
	if err != nil {
		if _, ok := err.(model.NotFoundError); ok {
			fmt.Println("Client not found:", err.Error())
			http.Error(w, "Client not found", http.StatusNotFound)
			return
		}
		fmt.Println("Internal Server Error:", err.Error())
		// Handle other types of errors
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Client found, continue with order creation logic
	var orderProducts []model.OrderProduct

	for _, reqProduct := range request.Products {
		orderProduct := model.OrderProduct{
			ProductID: reqProduct.ProductID,
			Quantity:  reqProduct.Quantity,
		}
		orderProducts = append(orderProducts, orderProduct)
	}

	// Client found, continue with order creation logic
	order := model.Order{
		ID:         uuid.New(),
		ClientID:   client.ID,
		Products:   orderProducts,
		TotalPrice: request.TotalPrice,
		CreatedAt:  request.CreatedAt,
		Status:     fetchOrderStatusId(o.orderStatus, Ordered),
	}

	// Set the CreatedAt field to the current timestamp
	order.CreatedAt = time.Now()

	// Create the order in the database
	err = o.orderRepo.Create(&order)
	if err != nil {
		fmt.Println("Failed to create order: %s", err.Error())

		http.Error(w, fmt.Sprintf("Failed to create order: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	//  This could be done by a  worker via rmq/kafka messages.
	err = o.procurements.Create(order.Products, order.CreatedAt)
	if err != nil {
		fmt.Println("Failed to create procurements: %s", err.Error())
		http.Error(w, fmt.Sprintf("Failed to create procurement: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	// Create the response using CreateOrderResponse struct
	response := CreateOrderResponse{
		ID:         order.ID,
		ClientID:   order.ClientID,
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
		Status:     Ordered,
	}

	// Convert order.Products to []ProductRequest
	for _, product := range order.Products {
		requestProduct := ProductResponse{
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		}
		response.Products = append(response.Products, requestProduct)
	}

	// Respond with the created order
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List orders")

	clientIDString := r.URL.Query().Get("client_id")

	fmt.Printf("\n Failed to retrieve orders: %s", clientIDString)

	// Check if a client ID filter is provided
	if clientIDString != "" {
		clientID, err := uuid.Parse(clientIDString)
		if err != nil {
			http.Error(w, "Invalid client ID format", http.StatusBadRequest)
			return
		}

		// Use the client ID to filter orders
		orders, err := o.orderRepo.GetByClientID(clientID)
		if err != nil {
			http.Error(w, "Failed to retrieve orders by client ID", http.StatusInternalServerError)
			return
		}

		// Convert orders to CreateOrderResponse
		var response []CreateOrderResponse
		for _, order := range orders {
			createOrderResponse := CreateOrderResponse{
				ID:         order.ID,
				ClientID:   order.ClientID,
				TotalPrice: order.TotalPrice,
				CreatedAt:  order.CreatedAt,
				Status:     Ordered,
			}

			for _, product := range order.Products {
				requestProduct := ProductResponse{
					ProductID: product.ProductID,
					Quantity:  product.Quantity,
				}
				createOrderResponse.Products = append(createOrderResponse.Products, requestProduct)
			}

			response = append(response, createOrderResponse)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// If no filter is provided, return all orders
	orders, err := o.orderRepo.ListAll()
	if err != nil {
		fmt.Printf("\n Failed to retrieve orders: %s", err.Error())
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	// Convert orders to CreateOrderResponse
	var response []CreateOrderResponse
	for _, order := range orders {
		createOrderResponse := CreateOrderResponse{
			ID:         order.ID,
			ClientID:   order.ClientID,
			TotalPrice: order.TotalPrice,
			CreatedAt:  order.CreatedAt,
			Status:     Ordered,
		}

		for _, product := range order.Products {
			requestProduct := ProductResponse{
				ProductID: product.ProductID,
				Quantity:  product.Quantity,
			}
			createOrderResponse.Products = append(createOrderResponse.Products, requestProduct)
		}

		response = append(response, createOrderResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetByClientID retrieves orders specific to a given client.
func (o *Order) GetByClientID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("chec this ")

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
