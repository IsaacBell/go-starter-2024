ifeq ($(OS), Windows_NT)
	BIN_FILENAME  := go-template-gateway.exe
else
	BIN_FILENAME  := go-template-gateway
endif

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "bin" rd /s /q bin	
else
	rm -fR ./bin
endif

.PHONY: build
build: clean
	go build -o ./bin/${BIN_FILENAME} ./cmd/api

.PHONY: docker-build
docker-build:
	docker build -t soapstone-image .

.PHONY: docker-compose
docker-compose:
	docker-compose --env-file .env stop
	docker-compose --env-file .env down
	docker-compose --env-file .env up -d

.PHONY: install-dependencies
install-dependencies:
	sudo apt install -y protobuf-compiler
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

.PHONY: setup-env
setup-env:
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin

.PHONY: execute
execute: docker-build docker-compose setup-env install-dependencies build
	sleep 2
	./bin/${BIN_FILENAME}