package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	orderpb "api-gateway/generated/order"
	productpb "api-gateway/generated/product"
	jwt "github.com/dgrijalva/jwt-go"
	gcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// Client connections to gRPC services
var orderClient orderpb.OrderServiceClient
var productClient productpb.ProductServiceClient
var orderServiceAddress = "order_service:50052"
var productServiceAddress = "product_service:50057"
var rumorSecretKey = "rumor_secret_key"
var rumorUser = "admin"
var rumorPassword = "Kt3RickS0n"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Initialize gRPC clients
func init() {
	// Create gRPC client connections with insecure credentials
	orderConn, err := grpc.NewClient(orderServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create client connection for order server: %v", err)
	}
	//defer orderConn.Close()

	productConn, err := grpc.NewClient(productServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create client connection for product server: %v", err)
	}
	//defer productConn.Close()

	// Create clients from the connections
	orderClient = orderpb.NewOrderServiceClient(orderConn)
	productClient = productpb.NewProductServiceClient(productConn)
}

func sendResponse(w http.ResponseWriter, status int, data interface{}, err error, message string) {
	response := struct {
		Status  int         `json:"status"`
		Data    interface{} `json:"data"`
		Error   string      `json:"error"`
		Success bool        `json:"success"`
		Message string      `json:"message"`
	}{
		Status:  status,
		Data:    data,
		Success: err == nil,
		Message: message,
	}

	if err != nil {
		response.Error = "Internal Server Error: " + err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// Handlers for REST endpoints
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendResponse(w, http.StatusMethodNotAllowed, nil, nil, "Method not allowed")
		return
	}

	var request productpb.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendResponse(w, http.StatusBadRequest, nil, err, "Invalid request body")
		return
	}

	response, err := productClient.CreateProduct(context.Background(), &request)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil, err, "Error creating product")
		return
	}

	sendResponse(w, http.StatusOK, response, nil, "OK")
}
func getProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendResponse(w, http.StatusMethodNotAllowed, nil, nil, "Method not allowed")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		sendResponse(w, http.StatusBadRequest, nil, nil, "Missing product ID")
		return
	}

	request := &productpb.FindProductByIdRequest{
		Id: id,
	}

	response, err := productClient.FindProductById(context.Background(), request)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil, err, "Error retrieving product")
		return
	}

	sendResponse(w, http.StatusOK, response, nil, "OK")
}
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		sendResponse(w, http.StatusMethodNotAllowed, nil, nil, "Method not allowed")
		return
	}

	var request productpb.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendResponse(w, http.StatusBadRequest, nil, err, "Invalid request body")
		return
	}

	response, err := productClient.UpdateProduct(context.Background(), &request)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil, err, "Error updating product")
		return
	}

	sendResponse(w, http.StatusOK, response, nil, "OK")
}
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		sendResponse(w, http.StatusMethodNotAllowed, nil, nil, "Method not allowed")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		sendResponse(w, http.StatusBadRequest, nil, nil, "Missing product ID")
		return
	}

	request := &productpb.DeleteProductRequest{
		Id: id,
	}

	response, err := productClient.DeleteProduct(context.Background(), request)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil, err, "Error deleting product")
		return
	}

	sendResponse(w, http.StatusOK, response, nil, "OK")
}
func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendResponse(w, http.StatusMethodNotAllowed, nil, nil, "Method not allowed")
		return
	}

	request := &productpb.FindAllProductsRequest{}

	response, err := productClient.FindAllProducts(context.Background(), request)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil, err, "Error retrieving products")
		return
	}

	sendResponse(w, http.StatusOK, response, nil, "OK")
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendResponse(w, http.StatusMethodNotAllowed, nil, nil, "Method not allowed")
		return
	}

	var request orderpb.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendResponse(w, http.StatusBadRequest, nil, err, "Invalid request body")
		return
	}

	response, err := orderClient.CreateOrder(context.Background(), &request)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil, err, "Error creating order")
		return
	}

	requestJSON, err := json.Marshal(request)
	if err != nil {
		// Handle error if unable to marshal request to JSON
		sendResponse(w, http.StatusInternalServerError, nil, err, "Error marshaling request JSON")
		return
	}
	fmt.Println("Request body JSON:", string(requestJSON))

	sendResponse(w, http.StatusOK, response, nil, "OK")
}
func updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		sendResponse(w, http.StatusMethodNotAllowed, nil, nil, "Method not allowed")
		return
	}

	var request orderpb.UpdateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		sendResponse(w, http.StatusBadRequest, nil, err, "Invalid request body")
		return
	}

	response, err := orderClient.UpdateOrder(context.Background(), &request)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil, err, "Error updating order")
		return
	}

	sendResponse(w, http.StatusOK, response, nil, "OK")
}
func deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		sendResponse(w, http.StatusMethodNotAllowed, nil, nil, "Method not allowed")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		sendResponse(w, http.StatusBadRequest, nil, nil, "Missing order ID")
		return
	}

	request := &orderpb.DeleteOrderRequest{
		Id: id,
	}

	response, err := orderClient.DeleteOrder(context.Background(), request)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil, err, "Error deleting order")
		return
	}

	sendResponse(w, http.StatusOK, response, nil, "OK")
}
func getAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendResponse(w, http.StatusMethodNotAllowed, nil, nil, "Method not allowed")
		return
	}

	request := &orderpb.FindAllOrdersRequest{}

	response, err := orderClient.FindAllOrders(context.Background(), request)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil, err, "Error retrieving orders")
		return
	}

	sendResponse(w, http.StatusOK, response, nil, "OK")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Construye el JSON de respuesta
	jsonResponse := map[string]interface{}{
		"message": "started correctly",
		"success": true,
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
func validateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verifica si la ruta es "/" o "/login"
		if r.URL.Path == "/" || r.URL.Path == "/login" {
			// Si es así, permite el acceso sin autenticación
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verifica el método de firma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(rumorSecretKey), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Si el token es válido, almacenar el usuario en el contexto
		gcontext.Set(r, "user", token.Claims.(jwt.MapClaims)["username"])
		w.Header().Set("Content-Type", "application/json")
		// Continuar con la siguiente manejador
		next.ServeHTTP(w, r)
	})
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Imprime la estructura User como JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding User to JSON", http.StatusInternalServerError)
		return
	}
	fmt.Println(string(userJSON))

	// Verifica las credenciales del usuario
	if user.Username == rumorUser && user.Password == rumorPassword {
		// Crea el token JWT
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expira en 24 horas

		// Firma el token JWT
		tokenString, err := token.SignedString([]byte(rumorSecretKey))
		if err != nil {
			http.Error(w, "Error signing token", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		// Devuelve el token JWT al cliente
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
}

func main() {
	// Create HTTP server

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	env := os.Getenv("ENVIRONMENT")
	secret_key := os.Getenv("SECRET_KEY")
	user := os.Getenv("AUTH_USER")
	password := os.Getenv("AUTH_PASSWORD")

	if env == "develop" {
		orderServiceAddress = "localhost:50052"
		productServiceAddress = "localhost:50057"
	}
	if secret_key != "" {
		rumorSecretKey = secret_key
	}
	if user != "" {
		rumorUser = user
	}
	if password != "" {
		rumorPassword = password
	}
	fmt.Printf("orderservice =  %s  productService = %s ", orderServiceAddress, productServiceAddress)

	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("POST")

	router.HandleFunc("/api/orders", getAllOrdersHandler).Methods("GET")
	router.HandleFunc("/api/orders/create", createOrderHandler).Methods("POST")
	router.HandleFunc("/api/orders/update", updateOrderHandler).Methods("PUT")
	router.HandleFunc("/api/orders/delete", deleteOrderHandler).Methods("DELETE")

	router.HandleFunc("/api/products", getAllProductsHandler).Methods("GET")
	router.HandleFunc("/api/products/create", createProductHandler).Methods("POST")
	router.HandleFunc("/api/products/get", getProductByIdHandler).Methods("GET")
	router.HandleFunc("/api/products/update", updateProductHandler).Methods("PUT")
	router.HandleFunc("/api/products/delete", deleteProductHandler).Methods("DELETE")
	router.HandleFunc("/", indexHandler).Methods("GET")

	// Use el middleware para validar el token JWT
	router.Use(validateTokenMiddleware)

	fmt.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
