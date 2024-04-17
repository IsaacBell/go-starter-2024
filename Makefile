CONFIG_PATH=${HOME}/.config/

ifeq ($(OS), Windows_NT)
    BIN_FILENAME  := go-template-gateway.exe
else
    BIN_FILENAME  := go-template-gateway
endif

.PHONY: init
init:
	mkdir -p ${CONFIG_PATH}
.PHONY: gencert
gencert:
	cfssl gencert \
		-initca test/ca-csr.json | cfssljson -bare ca
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=test/ca-config.json \
		-profile=server \
		test/server-csr.json | cfssljson -bare server
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=test/ca-config.json \
		-profile=client \
		test/client-csr.json | cfssljson -bare client
	mv *.pem *.csr ${CONFIG_PATH}

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

.PHONY: test
test:
	go test -race ./...

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
	export PATH="$PATH:$(go env GOPATH)/bin"
ifeq ($(shell uname), Darwin)
	brew install protobuf
	brew install go
else
	sudo apt install -y protobuf-compiler
	sudo apt install -y golang-go
endif
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	
.PHONY: compile
compile:
	protoc api/v1/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.

.PHONY: execute
execute: docker-build docker-compose install-dependencies build
	sleep 1
	./bin/${BIN_FILENAME}