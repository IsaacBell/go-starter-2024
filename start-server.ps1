# Build the Docker image
docker build -t soapstone-image .

# Set the Docker Compose command with the .env file
$soapstone_compose = "docker-compose --env-file .env"

# Set the short version of the Docker Compose command
$s_compose = "docker-compose --env-file .env"

# Stop and remove the containers, then start them up again
& $s_compose stop
& $s_compose down
& $s_compose up

# Set the GOPATH and update the PATH
$env:GOPATH = "$env:USERPROFILE\go"
$env:PATH += ";$env:GOPATH\bin"

# Install protobuf-compiler
# Note: There's no direct equivalent of `apt` in PowerShell.
# You can download the protobuf-compiler binaries from the official website
# and add the path to the binaries to your system's PATH environment variable.

# Install protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Clean, build, and run the Go application
go clean
go build
go run ./cmd/api/main.go
