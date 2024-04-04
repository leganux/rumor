package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"

	orderpb "api-gateway/generated/order"
	productpb "api-gateway/generated/product"

	gcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// Client connections to gRPC services
var orderClient orderpb.OrderServiceClient
var productClient productpb.ProductServiceClient

const (
	//orderServiceAddress   = "127.0.0.1:50052" // Dirección del servicio de órdenes
	//productServiceAddress = "127.0.0.1:50057" // Dirección del servicio de productos
	orderServiceAddress   = "order_service:50052"   // Dirección del servicio de órdenes
	productServiceAddress = "product_service:50057" // Dirección del servicio de productos
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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
			return []byte("your-secret-key"), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Si el token es válido, almacenar el usuario en el contexto
		gcontext.Set(r, "user", token.Claims.(jwt.MapClaims)["username"])

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
	if user.Username == "admin" && user.Password == "Kt3RickS0n" {
		// Crea el token JWT
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expira en 24 horas

		// Firma el token JWT
		tokenString, err := token.SignedString([]byte("your-secret-key"))
		if err != nil {
			http.Error(w, "Error signing token", http.StatusInternalServerError)
			return
		}

		// Devuelve el token JWT al cliente
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
}

func main() {
	// Create HTTP server
	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/api/orders", createOrderHandler).Methods("POST")
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
