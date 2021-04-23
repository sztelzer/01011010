module github.com/sztelzer/01011010

replace (
	github.com/sztelzer/01011010/shippingportsclient => ./shippingportsclient
	github.com/sztelzer/01011010/shippingportsprotocol => ./shippingportsprotocol
	github.com/sztelzer/01011010/shippingportsserver => ./shippingportsserver
)

go 1.16
