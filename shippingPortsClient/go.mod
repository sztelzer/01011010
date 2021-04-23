module github.com/sztelzer/01011010/shippingportsclient

replace github.com/sztelzer/01011010/shippingportsprotocol => ../shippingportsprotocol

go 1.16

require (
	github.com/sztelzer/01011010/shippingportsprotocol v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
)
