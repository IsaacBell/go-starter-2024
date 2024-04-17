#!/bin/bash

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export PATH="$PATH:$(go env GOPATH)/bin"

OS=$(uname)

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

protoc --go_out=. \
  proto/basic/*.proto \
  proto/api/v1/*.proto
