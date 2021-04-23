module github.com/sztelzer/01011010/shippingportsserver

go 1.16

replace github.com/sztelzer/01011010/shippingportsprotocol => ../shippingPortsProtocol

replace github.com/sztelzer/01011010/memdatabase => ../memdatabase

require (
	github.com/sztelzer/01011010/memdatabase v0.0.0-00010101000000-000000000000
	github.com/sztelzer/01011010/shippingportsprotocol v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
)
