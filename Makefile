BIND_PATH=~/shipping
JSON_FILENAME=ports.json

all: run

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative shippingportsprotocol/shippingports.proto

tidy: proto
	cd memdatabase; go mod tidy
	cd shippingportsprotocol; go mod tidy
	cd shippingportsserver; go mod tidy
	cd shippingportsclient;	go mod tidy

test: tidy
	cd memdatabase; go test ./... -v
	cd shippingportsserver; go test ./... -v
	cd shippingportsclient; go test ./... -v

docker: test
	docker build .

run:
	env BIND_PATH=$(BIND_PATH) JSON_FILENAME=$(JSON_FILENAME) docker compose up