package main

import (
	"log"
	"net/http"
	"os"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
	"google.golang.org/grpc"
)

// shippingPortsServerAddress is the default connection string to the gRPC shippingPortsServer
// It may be changed by the environment variable SERVER_ADDRESS
var shippingPortsServerAddress = "shippingportsserver:50051"

// shippingPortsClientAddress is the default port to serve the REST API
var shippingPortsClientAddress = ":8080"

// default json file to load from
var shippingPortsOriginJsonFile = "./ports.json"

func init() {
	// load env files to substitute defaults
	if envServerAddress, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		shippingPortsServerAddress = envServerAddress
	}
	
	if envClientAddress, ok := os.LookupEnv("SERVE_CLIENT_AT_ADDRESS"); ok {
		shippingPortsClientAddress = envClientAddress
	}
	
	if envLoadFilename, ok := os.LookupEnv("LOAD_SHIPPING_PORTS_JSON_FILENAME"); ok {
		shippingPortsOriginJsonFile = "dropbox/" + envLoadFilename
	}
	
}

func main() {
	
	// Set up a connection to the ports server
	// It will block if can't connect, like waiting for the server go online
	// It will exit if occur an explicit error connecting
	serverConnection, err := grpc.Dial(shippingPortsServerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// Connection closes if exiting.
	defer serverConnection.Close()
	
	shippingPortsServerClient := shippingportsprotocol.NewShippingPortsServerClient(serverConnection)
	
	// load ports from file to shippingPortsServerClient, don't wait the loading
	go saveShippingPortsFromFile(shippingPortsOriginJsonFile, shippingPortsServerClient)
	
	// Serve the REST API
	// Register a handler for requests
	http.HandleFunc("/", mainHandler(shippingPortsServerClient))
	
	// Start server
	log.Fatal(http.ListenAndServe(shippingPortsClientAddress, nil))
}
