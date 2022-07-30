.PHONY: proto build

build: proto
	go build -o client ./cmd/client/main.go
	go build -o server ./cmd/server/main.go

proto:
	protoc  --go_out=./ --go-grpc_out=./ -I ./ proto/controlplane.proto --validate_out="lang=go:./"
