package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/bhuvana-chinnadurai/farmaze-backend/api"
	"github.com/bhuvana-chinnadurai/farmaze-backend/config" // Replace with the actual module path
	"github.com/bhuvana-chinnadurai/farmaze-backend/model"
)

func InitDB(dbConfig config.DBConfig) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func SetupRoutes(conn *sql.DB) *mux.Router {
	fmt.Println("setting up routes")
	router := mux.NewRouter()

	clientRepo := model.NewB2BClientRepository(conn)
	productRepo := model.NewProductRepository(conn)
	orderProductRepo := model.NewOrderProductRepository(conn)
	orderRepo := model.NewOrderRepository(conn, orderProductRepo)
	orderStatusRepo := model.NewOrderStatusRepository(conn)
	orderStatuses, err := orderStatusRepo.GetAll()
	if err != nil {
		fmt.Println("error occurred while fetching orderStatus: ", err.Error())
	}

	b2bClient := api.New(clientRepo)
	products := api.NewProducts(productRepo)
	orderAPI := api.NewOrder(orderRepo, productRepo, clientRepo, orderStatuses)

	// Define routes with corrected paths
	router.HandleFunc("/b2bclients/summary", b2bClient.GetSummary).Methods("GET")
	router.HandleFunc("/b2bclients/{client_id}", b2bClient.GetDetailsById).Methods("GET")
	router.HandleFunc("/products", products.GetAll).Methods("GET")
	router.HandleFunc("/orders", orderAPI.Create).Methods("POST")
	router.HandleFunc("/orders/{client_id}", orderAPI.GetByClientID).Methods("GET")

	return router
}

func main() {

	fmt.Println("welcome to farmaze backend")

	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	conn := InitDB(appConfig.DB)
	defer conn.Close()
	dbConfig := appConfig.DB
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Name,
	)
	fmt.Println("check,", connStr)

	fmt.Println("Connected to PostgreSQL")

	router := SetupRoutes(conn)
	fmt.Println("running on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
