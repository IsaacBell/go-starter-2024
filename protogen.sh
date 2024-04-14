export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

sudo apt install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

protoc --go_out=. proto/basic/*.proto