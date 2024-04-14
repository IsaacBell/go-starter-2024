# Use the official Go image with version 1.20 as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Set the executable permissions for the binary
RUN chmod +x main

# Expose the necessary port (adjust if needed)
EXPOSE 8080

# Set the entry point command to run the application
CMD ["/app/main"]
