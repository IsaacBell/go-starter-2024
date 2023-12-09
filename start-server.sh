#!/usr/bin/env bash

docker start postgresDB15

docker-compose --env-file dev_influxdb.env up

go clean && go build && go run ./cmd/api/main.go