#!/usr/bin/env bash

soapstone_compose="docker-compose --env-file .env"
$soapstone_compose rm && $soapstone_compose up

go clean && go build && go run ./cmd/api/main.go