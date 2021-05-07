package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
	"google.golang.org/grpc"
)

// shippingPortsServerAddress is the default connection string to the gRPC shippingportsserver
// It may be changed by the environment variable SERVER_ADDRESS
var shippingPortsServerAddress = ":50051"

// shippingPortsClientAddress is the default port to serve the REST API
var shippingPortsClientAddress = ":8080"

// default json file to load from
var shippingPortsOriginJsonFile = "./ports.json"

// default number of shippingPorts putters routines
var shippingPortsPutters = 8

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

	if envLoadPutters, ok := os.LookupEnv("LOAD_SHIPPING_PORTS_PUTTERS"); ok {
		n, err := strconv.Atoi(envLoadPutters)
		if err != nil {
			log.Printf("could not parse LOAD_SHIPPING_PORTS_PUTTERS=%s, using default", envLoadPutters)
		} else {
			shippingPortsPutters = n
		}
	}

}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	
	// Set up a connection to the ports server
	// It will block if can't connect, like waiting for the server go online
	// It will exit if occur an explicit error connecting
	serverConnection, err := grpc.Dial(shippingPortsServerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer serverConnection.Close()

	// create a client for this connection
	shippingPortsServerClient := shippingportsprotocol.NewShippingPortsServerClient(serverConnection)

	// load ports from file to shippingPortsServerClient, don't wait the loading
	go loadShippingPortsFromFileToServer(ctx, shippingPortsServerClient)

	// Serve the REST API
	// Register a handler for requests
	http.HandleFunc("/", mainHandler(shippingPortsServerClient))

	srv := &http.Server{Addr: shippingPortsClientAddress}

	// Start server
	go func() {
		err = srv.ListenAndServe()
		log.Println(err)
		cancel()
	}()
	// success, log the moment
	log.Printf("shippingPortsClient started")

	// Graceful shutdown
	select {
	case <-ctx.Done():
		srv.Shutdown(ctx)
		log.Println(ctx.Err())
		return
	}

}
