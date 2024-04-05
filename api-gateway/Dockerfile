# Use the official Go image
FROM golang:1.22-alpine

ENV ENVIRONMENT=Staging
ENV SECRET_KEY=87Pdmo76EFv_LGb7jTnunHniYD!90svCEC
ENV AUTH_USER=admin
ENV AUTH_PASSWORD=Kt3RickS0n


# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Command to run the application
CMD ["./main"]


