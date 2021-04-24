ARCH=amd64
BIND_PATH=~/shipping
JSON_FILENAME=ports.json

all: run

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative shippingPortsProtocol/shippingPorts.proto

tidy: proto
	cd shippingPortsProtocol; go mod tidy
	cd shippingPortsServer; go mod tidy
	cd shippingPortsClient;	go mod tidy

test: tidy
	cd shippingPortsServer; go test .
	cd shippingPortsClient; go test .

build: test
	cd shippingPortsServer; env CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build -o shippingPortsServerApp
	cd shippingPortsClient; env CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build -o shippingPortsClientApp

docker: build
	docker build ./shippingPortsServer -t shippingportsserver
	docker build ./shippingPortsClient -t shippingportsclient

run: docker
	env BIND_PATH=$(BIND_PATH) JSON_FILENAME=$(JSON_FILENAME) docker compose up

clean:
	go clean
	rm -f ./shippingPortsServer/shippingPortsServerApp
	rm -f ./shippingPortsClient/shippingPortsClientApp
