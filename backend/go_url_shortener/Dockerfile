# Use official Golang image as base
FROM golang:1.22.1-alpine as build

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port
EXPOSE 8090

# Command to run the executable
CMD ["./main"]