package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var token string

func TestLoginAndGetToken(t *testing.T) {
	// Simulate a login request
	loginData := map[string]string{
		"username": "admin",
		"password": "Kt3RickS0n",
	}
	loginBody, _ := json.Marshal(loginData)
	req, err := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewReader(loginBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the loginHandler function directly, passing in the ResponseRecorder and the Request
	http.HandlerFunc(loginHandler).ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("Login handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Extract token from the response body
	var tokenResponse map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &tokenResponse); err != nil {
		t.Fatalf("Error decoding token response: %v", err)
	}
	token = tokenResponse["token"]
	fmt.Printf("The token is %s", token)
}

func TestCreateOrderHandler(t *testing.T) {

	fmt.Printf("Order handler Start ")
	// Ensure token is obtained before running the test
	if token == "" {
		t.Fatal("Token not obtained from login")
	}

	// Simulate a request body
	requestBody := []byte(`{
        "customerId": "1",
        "products": [
            {
                "productId": "1",
                "quantity": 2
            }
        ]
    }`)

	// Create a new HTTP request with the request body
	req, err := http.NewRequest("POST", "/api/orders/create", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set the Authorization header with the obtained token
	req.Header.Set("Authorization", token)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the createOrderHandler function directly, passing in the ResponseRecorder and the Request
	http.HandlerFunc(createOrderHandler).ServeHTTP(rr, req)

	// Print status response
	fmt.Printf("Status Response: %v\n", rr.Code)

	// Print response body JSON
	fmt.Printf("Response Body JSON: %s\n", rr.Body.String())

	// Parse response body as JSON
	var responseBody map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	// Check if 'success' and 'message' fields exist in the response
	success, successExists := responseBody["success"].(bool)
	message, messageExists := responseBody["message"].(string)
	if !successExists || !messageExists {
		t.Errorf("response does not contain 'success' or 'message' fields")
	}

	fmt.Printf("Test successfully executed :: Parameters correctly appear in format response. \n  { success: %v, message: %s } \n", success, message)
	fmt.Printf("Order handler End")
}

func TestUpdateOrderHandler(t *testing.T) {
	// Simulate a request body
	requestBody := []byte(`{
		"id": "1",
		"customerId": "1",
		"products": [
			{
				"productId": "1",
				"quantity": 3
			}
		]
	}`)
	// Create a new HTTP request with the request body
	req, err := http.NewRequest("PUT", "/api/orders/update", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", ""+token)
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Call the updateOrderHandler function directly, passing in the ResponseRecorder and the Request
	http.HandlerFunc(updateOrderHandler).ServeHTTP(rr, req)
	// Parse response body as JSON
	var responseBody map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	// Check if 'success' and 'message' fields exist in the response
	success, successExists := responseBody["success"].(bool)
	message, messageExists := responseBody["message"].(string)
	if !successExists || !messageExists {
		t.Errorf("response does not contain 'success' or 'message' fields")
	}

	fmt.Printf("Test successfully executed :: Parameters correctly appear in format response. \n  { success: %v, message: %s } \n", success, message)
}

func TestDeleteOrderHandler(t *testing.T) {
	// Simulate a request with an order ID to delete
	req, err := http.NewRequest("DELETE", "/api/orders/delete?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", ""+token)
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Call the deleteOrderHandler function directly, passing in the ResponseRecorder and the Request
	http.HandlerFunc(deleteOrderHandler).ServeHTTP(rr, req)
	// Check the status code
	// Parse response body as JSON
	var responseBody map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	// Check if 'success' and 'message' fields exist in the response
	success, successExists := responseBody["success"].(bool)
	message, messageExists := responseBody["message"].(string)
	if !successExists || !messageExists {
		t.Errorf("response does not contain 'success' or 'message' fields")
	}

	fmt.Printf("Test successfully executed :: Parameters correctly appear in format response. \n  { success: %v, message: %s } \n", success, message)
}

func TestGetAllOrdersHandler(t *testing.T) {
	// Simulate a request to get all orders
	req, err := http.NewRequest("GET", "/api/orders", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", ""+token)
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Call the getAllOrdersHandler function directly, passing in the ResponseRecorder and the Request
	http.HandlerFunc(getAllOrdersHandler).ServeHTTP(rr, req)
	// Check the status code
	// Parse response body as JSON
	var responseBody map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	// Check if 'success' and 'message' fields exist in the response
	success, successExists := responseBody["success"].(bool)
	message, messageExists := responseBody["message"].(string)
	if !successExists || !messageExists {
		t.Errorf("response does not contain 'success' or 'message' fields")
	}

	fmt.Printf("Test successfully executed :: Parameters correctly appear in format response. \n  { success: %v, message: %s } \n", success, message)
}

func TestCreateProductHandler(t *testing.T) {
	// Simulate a request body to create a product
	requestBody := []byte(`{
		"name": "Test Product",
		"description": "Test Description",
		"price": 10.99,
		"quantity": 100
	}`)
	// Create a new HTTP request with the request body
	req, err := http.NewRequest("POST", "/api/products/create", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", ""+token)
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Call the createProductHandler function directly, passing in the ResponseRecorder and the Request
	http.HandlerFunc(createProductHandler).ServeHTTP(rr, req)
	// Check the status code
	// Parse response body as JSON
	var responseBody map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	// Check if 'success' and 'message' fields exist in the response
	success, successExists := responseBody["success"].(bool)
	message, messageExists := responseBody["message"].(string)
	if !successExists || !messageExists {
		t.Errorf("response does not contain 'success' or 'message' fields")
	}

	fmt.Printf("Test successfully executed :: Parameters correctly appear in format response. \n  { success: %v, message: %s } \n", success, message)
}

func TestUpdateProductHandler(t *testing.T) {
	// Simulate a request body to update a product
	requestBody := []byte(`{
		"id": "1",
		"name": "Updated Test Product",
		"description": "Updated Test Description",
		"price": 15.99,
		"quantity": 150
	}`)
	// Create a new HTTP request with the request body
	req, err := http.NewRequest("PUT", "/api/products/update", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", ""+token)
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Call the updateProductHandler function directly, passing in the ResponseRecorder and the Request
	http.HandlerFunc(updateProductHandler).ServeHTTP(rr, req)
	// Check the status code
	// Parse response body as JSON
	var responseBody map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	// Check if 'success' and 'message' fields exist in the response
	success, successExists := responseBody["success"].(bool)
	message, messageExists := responseBody["message"].(string)
	if !successExists || !messageExists {
		t.Errorf("response does not contain 'success' or 'message' fields")
	}

	fmt.Printf("Test successfully executed :: Parameters correctly appear in format response. \n  { success: %v, message: %s } \n", success, message)
}

func TestDeleteProductHandler(t *testing.T) {
	// Simulate a request with a product ID to delete
	req, err := http.NewRequest("DELETE", "/api/products/delete?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", ""+token)
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Call the deleteProductHandler function directly, passing in the ResponseRecorder and the Request
	http.HandlerFunc(deleteProductHandler).ServeHTTP(rr, req)
	// Check the status code
	// Parse response body as JSON
	var responseBody map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	// Check if 'success' and 'message' fields exist in the response
	success, successExists := responseBody["success"].(bool)
	message, messageExists := responseBody["message"].(string)
	if !successExists || !messageExists {
		t.Errorf("response does not contain 'success' or 'message' fields")
	}

	fmt.Printf("Test successfully executed :: Parameters correctly appear in format response. \n  { success: %v, message: %s } \n", success, message)
}

func TestGetAllProductsHandler(t *testing.T) {
	// Simulate a request to get all products
	req, err := http.NewRequest("GET", "/api/products", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", ""+token)
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Call the getAllProductsHandler function directly, passing in the ResponseRecorder and the Request
	http.HandlerFunc(getAllProductsHandler).ServeHTTP(rr, req)
	// Check the status code
	// Parse response body as JSON
	var responseBody map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	// Check if 'success' and 'message' fields exist in the response
	success, successExists := responseBody["success"].(bool)
	message, messageExists := responseBody["message"].(string)
	if !successExists || !messageExists {
		t.Errorf("response does not contain 'success' or 'message' fields")
	}

	fmt.Printf("Test successfully executed :: Parameters correctly appear in format response. \n  { success: %v, message: %s } \n", success, message)
}
