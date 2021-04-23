

all: run

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative shippingPortsProtocol/shippingPorts.proto

tidy: proto
	cd shippingPortsProtocol
	go mod tidy
	cd ..
	cd shippingPortsServer
	go mod tidy
	cd ..
	cd shippingPortsClient
	go mod tidy
	cd ..

test: tidy
	go test ./shippingPortsServer -v ./...
	go test ./shippingPortsClient -v ./...

build: test
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./shippingPortsServer/shippingPortsServerApp -v ./shippingPortsServer
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./shippingPortsClient/shippingPortsClientApp -v ./shippingPortsClient

docker: build
	docker build ./shippingPortsServer -t shippingPortsServer
	docker build ./shippingPortsClient -t shippingPortsClient

run: docker
	docker compose up -d

clean:
	go clean
	rm -f ./shippingPortsServer/shippingPortsServerApp
	rm -f ./shippingPortsClient/shippingPortsClientApp
