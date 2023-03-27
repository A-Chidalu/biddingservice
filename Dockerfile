# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o app cmd/server/main.go

# Expose port 8080 for the application
EXPOSE 5003

# Run the Go application
CMD ["./app"]
