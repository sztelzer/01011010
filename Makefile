BIND_PATH=~/shipping
JSON_FILENAME=ports.json

all: run

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative shippingportsprotocol/shippingports.proto

tidy: proto
	cd shippingportsprotocol; go mod tidy
	cd shippingportsserver; go mod tidy
	cd shippingportsclient;	go mod tidy
	cd shippingportsmemdatabase; go mod tidy

test: tidy
	cd shippingPortsServer; go test ./... -v
	cd shippingPortsClient; go test ./... -v
	cd shippingportsmemdatabase; go test ./... -v

docker: test
	docker build .

run:
	env BIND_PATH=$(BIND_PATH) JSON_FILENAME=$(JSON_FILENAME) docker compose up