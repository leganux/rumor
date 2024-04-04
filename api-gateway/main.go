package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	orderpb "api-gateway/generated/order"
	productpb "api-gateway/generated/product"
)

// Client connections to gRPC services
var orderClient orderpb.OrderServiceClient
var productClient productpb.ProductServiceClient

const (
	orderServiceAddress   = "localhost:50052" // Dirección del servicio de órdenes
	productServiceAddress = "localhost:50057" // Dirección del servicio de productos
)

func checkServiceStatus(address string) string {
	// Create a new gRPC client connection with insecure credentials
	conn, err := grpc.NewClient(address, grpc.WithInsecure())
	if err != nil {
		return "error"
	}
	//defer conn.Close()

	// Create a client from the connection
	client := orderpb.NewOrderServiceClient(conn)

	// Attempt to perform a simple RPC request to the service
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = client.FindAllOrders(ctx, &orderpb.FindAllOrdersRequest{})
	fmt.Printf("%s", err)
	if err != nil {
		return "error"
	}

	return "ok"
}

// Initialize gRPC clients
func init() {
	// Create gRPC client connections with insecure credentials
	orderConn, err := grpc.NewClient(orderServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create client connection for order server: %v", err)
	}
	defer orderConn.Close()

	productConn, err := grpc.NewClient(productServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create client connection for product server: %v", err)
	}
	//defer productConn.Close()

	// Create clients from the connections
	orderClient = orderpb.NewOrderServiceClient(orderConn)
	productClient = productpb.NewProductServiceClient(productConn)
}

// Handlers for REST endpoints
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request productpb.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := productClient.CreateProduct(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating product: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func getProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	request := &productpb.FindProductByIdRequest{
		Id: id,
	}

	response, err := productClient.FindProductById(context.Background(), request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving product: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request productpb.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := productClient.UpdateProduct(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating product: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	request := &productpb.DeleteProductRequest{
		Id: id,
	}

	response, err := productClient.DeleteProduct(context.Background(), request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting product: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request orderpb.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := orderClient.CreateOrder(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating order: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request orderpb.UpdateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := orderClient.UpdateOrder(context.Background(), &request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating order: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing order ID", http.StatusBadRequest)
		return
	}

	request := &orderpb.DeleteOrderRequest{
		Id: id,
	}

	response, err := orderClient.DeleteOrder(context.Background(), request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting order: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	request := &productpb.FindAllProductsRequest{}

	response, err := productClient.FindAllProducts(context.Background(), request)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving products: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Verifica el estado del servicio de órdenes
	orderStatus := checkServiceStatus(orderServiceAddress)

	// Verifica el estado del servicio de productos
	productStatus := checkServiceStatus(productServiceAddress)

	// Construye el JSON de respuesta
	jsonResponse := map[string]interface{}{
		"message": "started correctly",
		"success": true,
		"connection": map[string]string{
			"orders":   orderStatus,
			"products": productStatus,
		},
	}

	// Convierte el JSON a bytes
	responseData, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Establece el tipo de contenido de la respuesta como application/json
	w.Header().Set("Content-Type", "application/json")

	// Escribe la respuesta JSON
	_, err = w.Write(responseData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Create HTTP server
	http.HandleFunc("/api/orders", createOrderHandler)
	http.HandleFunc("/api/orders/update", updateOrderHandler)
	http.HandleFunc("/api/orders/delete", deleteOrderHandler)
	http.HandleFunc("/api/products", getAllProductsHandler)
	http.HandleFunc("/api/products/create", createProductHandler)
	http.HandleFunc("/api/products/get", getProductByIdHandler)
	http.HandleFunc("/api/products/update", updateProductHandler)
	http.HandleFunc("/api/products/delete", deleteProductHandler)

	http.HandleFunc("/", indexHandler)

	fmt.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
