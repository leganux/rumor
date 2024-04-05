## API Gateway

This repository contains the code for an API Gateway that acts as an entry point to interact with microservices managing
orders and products. The API Gateway handles authentication, routing, and forwards requests to the appropriate
microservices. This README file provides instructions on how to execute and use the API Gateway.

### Prerequisites:

Docker installed on your machine.
Docker Compose installed on your machine.
Understanding of RESTful APIs and gRPC communication.
Execution:

1. Clone the Repository:
   Clone this repository to your local machine using the following command:

```bash
 git clone https://github.com/leganux/rumor
 ```

2. Start Docker Compose:
   Run the following command to start Docker Compose, which will set up the API Gateway along with the required
   microservices:

```bash 
docker-compose up
```

3. Access the API:
   Once Docker Compose has successfully started, you can access the API using the base URL http://localhost:8080.

### Files and Folders

```text

 -rumor
    -api-gateway
    -grpc-protocols
    -order-service
    -product-service
    RADME.md
    docker-compose.yml
    rumor.postman_collection.json
    
 ```

*Description Folders*

* api-gateway: Includes files from GO project, gateway api client is connected to product and order services.
* grpc-protocols: Includes generated files from protoc for both projects client service and servers in nestjs, includes
  proto files too.
* order-service: Includes order micro-service nestjs files to control orders.
* product-service: Includes product micro-service nestjs files to control product.
* docker-compose.yml: The compose file to run all infrastructure of this project
* README.md: This file documentation
* rumor.postman_collection.json: Its file of postman documentation

### Endpoints with examples:

* POST /login:
  Endpoint to authenticate and receive a JWT token. Requires a JSON body with username and password fields.

```javascript 
const myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");

const raw = JSON.stringify({
    "Username": "admin",
    "Password": "Kt3RickS0n"
});

const requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow"
};

fetch("http://localhost:8080/login", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));
```

* GET /api/orders:
  Retrieve all orders.

```Javascript
const myHeaders = new Headers();
myHeaders.append("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTIzNTEyMTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.nApWupO7nUbIEstbtQF5crzPh5ugE7V1LB_23j1g0qg");

const raw = "";

const requestOptions = {
    method: "GET",
    headers: myHeaders,
    body: raw,
    redirect: "follow"
};

fetch("http://localhost:8080/api/orders", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));
```

* POST /api/orders/create:
  Create a new order. Requires a JSON body with customerId and products fields.

```javascript
 const myHeaders = new Headers();
myHeaders.append("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTIzNTEyMTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.nApWupO7nUbIEstbtQF5crzPh5ugE7V1LB_23j1g0qg");
myHeaders.append("Content-Type", "application/json");

const raw = JSON.stringify({
    "customerId": "1",
    "products": [
        {
            "productId": "1",
            "quantity": 2
        },
        {
            "productId": "2",
            "quantity": 1
        }
    ]
});

const requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow"
};

fetch("http://localhost:8080/api/orders/create", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));
```

* PUT /api/orders/update:
  Update an existing order. Requires a JSON body with id, customerId, and products fields.

```javascript
const myHeaders = new Headers();
myHeaders.append("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTIzNTEyMTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.nApWupO7nUbIEstbtQF5crzPh5ugE7V1LB_23j1g0qg");
myHeaders.append("Content-Type", "application/json");

const raw = JSON.stringify({
    "id": "1",
    "customerId": "3",
    "products": [
        {
            "productId": "1",
            "quantity": 2
        },
        {
            "productId": "2",
            "quantity": 1
        },
        {
            "productId": "3",
            "quantity": 10
        }
    ]
});

const requestOptions = {
    method: "PUT",
    headers: myHeaders,
    body: raw,
    redirect: "follow"
};

fetch("http://localhost:8080/api/orders/update", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));
```

* DELETE /api/orders/delete?id={orderId}:
  Delete an existing order. Requires id parameter in the URL.

```javascript
const myHeaders = new Headers();
myHeaders.append("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTIzNTEyMTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.nApWupO7nUbIEstbtQF5crzPh5ugE7V1LB_23j1g0qg");

const raw = "";

const requestOptions = {
    method: "DELETE",
    headers: myHeaders,
    body: raw,
    redirect: "follow"
};

fetch("http://localhost:8080/api/orders/delete?id=1", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));
```

* GET /api/products:
  Retrieve all products.

```javascript
const myHeaders = new Headers();
myHeaders.append("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTIzNTEyMTcsInVzZXJuYW1lIjoiYWRtaW4ifQ.nApWupO7nUbIEstbtQF5crzPh5ugE7V1LB_23j1g0qg");

const raw = "";

const requestOptions = {
    method: "GET",
    headers: myHeaders,
    body: raw,
    redirect: "follow"
};

fetch("http://localhost:8080/api/products", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));
```

* POST /api/products/create:
  Create a new product. Requires a JSON body with name, description, price, and quantity fields.

```javascript
const myHeaders = new Headers();
myHeaders.append("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTIzNTc3MDEsInVzZXJuYW1lIjoiYWRtaW4ifQ.0W9DBHdr0vB370jdj-Lx1SAaFvlwL2S1YU9D-_U4BQM");
myHeaders.append("Content-Type", "application/json");

const raw = JSON.stringify({
    "name": "Coke 100 Ml2",
    "description": "Soda 100 ML cola flavor2",
    "price": 34.2,
    "quantity": 5
});

const requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow"
};

fetch("http://localhost:8080/api/products/create", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));
```

* GET /api/products/get?id={productId}:
  Retrieve a product by its ID. Requires id parameter in the URL.

```javascript
const myHeaders = new Headers();
myHeaders.append("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTIzNTc3MDEsInVzZXJuYW1lIjoiYWRtaW4ifQ.0W9DBHdr0vB370jdj-Lx1SAaFvlwL2S1YU9D-_U4BQM");

const raw = "";

const requestOptions = {
    method: "GET",
    headers: myHeaders,
    body: raw,
    redirect: "follow"
};

fetch("http://localhost:8080/api/products/get?id=1", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));
```

* PUT /api/products/update:
  Update an existing product. Requires a JSON body with id, name, description, price, and quantity fields.

```javascript
const myHeaders = new Headers();
myHeaders.append("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTIzNTc3MDEsInVzZXJuYW1lIjoiYWRtaW4ifQ.0W9DBHdr0vB370jdj-Lx1SAaFvlwL2S1YU9D-_U4BQM");
myHeaders.append("Content-Type", "application/json");

const raw = JSON.stringify({
    "id": "1",
    "name": "Cokessseeeeeeeeeeees 100 Ml2",
    "description": "Soda eeeeeeee100 ML cola flavor2",
    "price": 354.2,
    "quantity": 5
});

const requestOptions = {
    method: "PUT",
    headers: myHeaders,
    body: raw,
    redirect: "follow"
};

fetch("http://localhost:8080/api/products/update", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));
```

* DELETE /api/products/delete?id={productId}:
  Delete an existing product. Requires id parameter in the URL.

```javascript
const myHeaders = new Headers();
myHeaders.append("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTIzNTc3MDEsInVzZXJuYW1lIjoiYWRtaW4ifQ.0W9DBHdr0vB370jdj-Lx1SAaFvlwL2S1YU9D-_U4BQM");
myHeaders.append("Content-Type", "application/json");

const raw = JSON.stringify({
    "id": "2",
    "name": "Cokessss 100 Ml2",
    "description": "Soda 100 ML cola flavor2",
    "price": 34.2,
    "quantity": 5
});

const requestOptions = {
    method: "DELETE",
    headers: myHeaders,
    body: raw,
    redirect: "follow"
};

fetch("http://localhost:8080/api/products/delete?id=1", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));
```

### Response Standardized example

* status: The code status
* data: The information tha user request for
* error: The internal error message
* success: If request was correct or not
* Message: A message to show to user

*OK Response Example*

```json
{
  "status": 200,
  "data": {
    "products": [
      {
        "id": "1",
        "name": "Coke 100 Ml2",
        "description": "Soda 100 ML cola flavor2",
        "price": 34.2,
        "quantity": 5
      }
    ]
  },
  "error": "",
  "success": true,
  "message": "OK"
}

```

*Error Response Example*

```json
{
  "status": 500,
  "data": null,
  "error": "Internal Server Error: rpc error: code = Unknown desc = Internal server error",
  "success": false,
  "message": "Error retrieving product"
}
```

### Test

To execute testing

```bash
go test
```

<small>*note: Testing evaluates only mock structure of response, cause go cannot instantiate GRPC communication </small>

### Authentication:

To access any endpoint other than /login, you need to include the JWT token received from the /login endpoint in the
request headers under the key Authorization.

### Conclusion:

The API Gateway provides a centralized entry point to interact with microservices managing orders and products. By
following the provided instructions, you can easily set up and utilize the API Gateway to perform various operations
related to orders and products.

<hr>


<p align="center">
    <img src="https://www.kingtide.com/_next/static/media/logo.1fa4aea8.svg" width="100" title="hover text">
    <br>
  This project is a test for rumor made by <a href="https://www.kingtide.com">https://www.kingtide.com</a>  and Angel Erick Cruz O. &copy; 2024 rights reserved.
    <br>
   This project is distributed under the MIT license. 
    <br>
<br>
    The project was made with ♥️ by Angel Erick Cruz Olivera and KingTide team
<br>
<br>
This project was built with docker, nestjs, go, nodejs, grpc and other tools. Their logos and base software for their implementation reserve their rights to their own creators.
<br>