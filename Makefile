BIND_PATH=~/shipping
JSON_FILENAME=ports.json

all: run

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative shippingPortsProtocol/shippingPorts.proto

tidy: proto
	cd shippingPortsProtocol; go mod tidy
	cd shippingPortsServer; go mod tidy
	cd shippingPortsClient;	go mod tidy
	cd memdatabase;	go mod tidy

test: tidy
	cd shippingPortsServer; go test ./... -v
	cd shippingPortsClient; go test ./... -v
	cd memdatabase; go test ./... -v

docker: test
	docker build .

run:
	env BIND_PATH=$(BIND_PATH) JSON_FILENAME=$(JSON_FILENAME) docker compose up