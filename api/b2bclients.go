package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/bhuvana-chinnadurai/farmaze-backend/model"
)

type B2B struct {
	clients model.B2BClients // Add ClientRepository as a field
}

func New(clients model.B2BClients) *B2B {
	return &B2B{
		clients: clients,
	}
}

// GetSummary godoc
// @Summary Get client summaries
// @Description Retrieve a list of all B2B clients' summaries
// @Tags B2B
// @Accept  json
// @Produce  json
// @Success 200 {array} model.B2BClient "List of B2B client summaries"
// @Router /api/v1/b2bclients [get]
func (b *B2B) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("check")
	// Use the ClientRepository to get the client summaries
	clients, err := b.clients.GetAllClients()
	if err != nil {
		// Handle the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Send the list of clients as a response
	//You may choose to serialize the clients to JSON or any other format you prefer
	//For example, you can use encoding/json package to marshal clients to JSON
	//Then write the JSON response to the http.ResponseWriter
	//Example:
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

// GetDetailsById godoc
// @Summary Get client details by ID
// @Description Retrieve details of a specific B2B client by their ID
// @Tags B2B
// @Accept  json
// @Produce  json
// @Param client_id path string true "Client ID"
// @Success 200 {object} model.B2BClient "B2B client details"
// @Failure 400 {object} map[string]string "Invalid client ID format"
// @Failure 404 {object} map[string]string "Client not found"
// @Router /api/v1/b2bclients/{client_id} [get]
func (b *B2B) GetDetailsById(w http.ResponseWriter, r *http.Request) {
	clientIDString := mux.Vars(r)["client_id"]

	clientID, err := uuid.Parse(clientIDString)
	if err != nil {
		// Respond with Bad Request and error message
		errorMessage := "Invalid client ID format"
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}
	client, err := b.clients.GetClientByID(clientID)
	if err != nil {
		// Handle error (client not found)
		// Respond with appropriate status code and error message
	}
	//Send the client details as a response
	//You may choose to serialize the client to JSON or any other format you prefer
	//For example, you can use encoding/json package to marshal client to JSON
	//Then write the JSON response to the http.ResponseWriter
	//Example:
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}
