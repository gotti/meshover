.PHONY: proto build deploy

build: proto build-ci

build-ci:
	CGO_ENABLED=0 go build -o client ./cmd/client/main.go
	CGO_ENABLED=0 go build -o meshserver ./cmd/server/main.go
	CGO_ENABLED=0 go build -o exporter ./cmd/exporter/main.go

proto:
	protoc  --go_out=./ --go-grpc_out=./ -I ./ proto/statusmanager.proto proto/controlplane.proto proto/ip.proto --validate_out="lang=go:./"

deploy: build
	scp ./meshserver meshover:~/
	ssh meshover2 "sudo pkill client" || echo "$?"
	ssh meshover3 "sudo pkill client" || echo "$?"
	scp ./client meshover2:~/
	scp ./client meshover3:~/
	ssh meshover2 "nohup sudo ./client -controlserver 192.168.129.66:12384 -statusserver 192.168.129.66:12385 -frr nerdctl &" &
	ssh meshover3 "nohup sudo ./client -controlserver 192.168.129.66:12384 -statusserver 192.168.129.66:12385 -frr nerdctl &" &

test:
	go mod tidy
	go test ./...
