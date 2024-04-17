#!/usr/bin/env bash

# Build Docker image
docker build -t soapstone-image .

# Define Docker Compose command
s_compose="docker-compose --env-file .env"

# Stop and remove existing containers, then start new containers
$s_compose stop
$s_compose down
$s_compose up -d

# Set GOPATH and update PATH
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Detect the operating system
OS=$(uname)

# Install protobuf-compiler based on the operating system
if [ "$OS" = "Darwin" ]; then
  # macOS
  brew install protobuf
elif [ "$OS" = "Linux" ]; then
  # Linux
  sudo apt install -y protobuf-compiler
else
  echo "Unsupported operating system: $OS"
  exit 1
fi

# Install protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Clean and build the Go project
go clean
go build

# Run the Go application
go run ./cmd/api/main.go