package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	pb "api-gateway/generated" // Import your compiled Protobuf package
)

const (
	productServiceAddress = "localhost:3007" // Address of Product Service gRPC server
	orderServiceAddress   = "localhost:3008" // Address of Order Service gRPC server
)

// Define a struct to represent the product data
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

// Define a struct to represent the order data
type Order struct {
	ID         string         `json:"id"`
	CustomerID string         `json:"customerId"`
	Products   []OrderProduct `json:"products"`
}

// Define a struct to represent the products within an order
type OrderProduct struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

// Sample product data
var products = []Product{}

// Create gRPC client connections
var productClient pb.ProductServiceClient
var orderClient pb.OrderServiceClient

func main() {
	// Create gRPC connections to Product Service and Order Service
	productConn, err := grpc.Dial(productServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Product Service: %v", err)
	}
	defer productConn.Close()
	productClient = pb.NewProductServiceClient(productConn)

	orderConn, err := grpc.Dial(orderServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Order Service: %v", err)
	}
	defer orderConn.Close()
	orderClient = pb.NewOrderServiceClient(orderConn)

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define endpoints
	router.HandleFunc("/products", getProducts).Methods("GET")
	router.HandleFunc("/orders", createOrder).Methods("POST")

	// Create a HTTP server
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the server
	log.Fatal(srv.ListenAndServe())
}

// Handler function to get all products
func getProducts(w http.ResponseWriter, r *http.Request) {
	// Call gRPC method to retrieve products from Product Service
	response, err := productClient.GetProducts(context.Background(), &pb.GetProductsRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert gRPC response to JSON
	products := []*Product{}
	for _, pbProduct := range response.Products {
		product := &Product{
			ID:          pbProduct.Id,
			Name:        pbProduct.Name,
			Description: pbProduct.Description,
			Price:       pbProduct.Price,
			Quantity:    int(pbProduct.Quantity),
		}
		products = append(products, product)
	}

	// Return products as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Handler function to create a new order
func createOrder(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into an Order struct
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert Order struct to gRPC request
	orderRequest := &pb.CreateOrderRequest{
		CustomerId: order.CustomerID,
		Products:   []*pb.OrderProduct{},
	}
	for _, p := range order.Products {
		orderRequest.Products = append(orderRequest.Products, &pb.OrderProduct{
			ProductId: p.ProductID,
			Quantity:  int32(p.Quantity),
		})
	}

	// Call gRPC method to create order in Order Service
	_, err = orderClient.CreateOrder(context.Background(), orderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
}
