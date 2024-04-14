$env:GOPATH = "$env:USERPROFILE\go"
$env:PATH += ";$env:GOPATH\bin"

# Install protobuf-compiler
# Note: There's no direct equivalent of `apt` in PowerShell.
# You can download the protobuf-compiler binaries from the official website
# and add the path to the binaries to your system's PATH environment variable.

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

protoc --go_out=. proto/basic/*.proto
