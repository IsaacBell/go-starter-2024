#!/usr/bin/env bash

docker build -t soapstone-image .

soapstone_compose="docker-compose --env-file .env"

s_compose="docker-compose --env-file .env"
$s_compose stop && $s_compose down && $s_compose up

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

sudo apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go clean && go build
go run ./cmd/api/main.go
