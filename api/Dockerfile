# Use official Golang image as base
FROM golang:1.24.2

# Set working directory inside container
WORKDIR /usr/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go app
RUN go build -o server .

# Expose the port the app runs on
EXPOSE 8080

# Run the built binary
CMD ["./server"]
