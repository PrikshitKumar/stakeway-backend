# Use an official Golang image as a base
FROM golang:1.23-alpine AS builder

# Set working directory inside container
WORKDIR /Prikshit

# Copy Go modules & download dependencies
COPY go.mod go.sum ./

# Download dependencies (uses cache if go.mod/go.sum don’t change)
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main .

# Use a smaller base image for the final container
FROM alpine:latest

# Set working directory inside container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /Prikshit/main .

# Expose the API port
EXPOSE 8080

# Run the application
CMD ["./main"]
