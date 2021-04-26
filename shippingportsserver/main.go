package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	"github.com/sztelzer/01011010/memdatabase"
	"google.golang.org/grpc"

	"github.com/sztelzer/01011010/shippingportsprotocol"
)

// shippingPortsServerAddress is the default shippingPortsProtocolServer address, but it may be changed by SERVE_AT_ADDRESS env variable.
var shippingPortsServerAddress = ":50051"

// default maxsize of 100 Mebibytes in memory.
var maxDatabaseSize = 1024 * 1024 * 100

// shippingPortsDatabase is the memdatabase instance that will be used during the life time of the server.
// It is a in memory memdatabase and it does not persist state.
var shippingPortsDatabase *memdatabase.Memdatabase

func init() {
	// just check if non default shippingPortsProtocolServer port was given in the environment (usually Dockerfile or docker-compose.yaml)
	if envPortsServerAddress, ok := os.LookupEnv("SERVE_AT_ADDRESS"); ok {
		shippingPortsServerAddress = envPortsServerAddress
	}

	// check if we should use a different database size
	if maxSizeMbStr, ok := os.LookupEnv("DATABASE_MAX_SIZE_MB"); ok {
		maxSizeMbInt, err := strconv.Atoi(maxSizeMbStr)
		if err != nil {
			log.Println("invalid value for DATABASE_MAX_SIZE_MB, please use integer, defaulting to 100 Mebibytes")
		} else {
			maxDatabaseSize = maxSizeMbInt * 1024 * 1024
		}
	}

	shippingPortsDatabase = memdatabase.New(maxDatabaseSize)
}


func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// open the shippingPortsServerAddress to the network. It will fail if the network port is already in use. It will not retry.
	listener, err := net.Listen("tcp", shippingPortsServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	// success, log the moment
	log.Printf("shippingPortsProtocolServer started")

	// get a new generic grpc server that we will attach the protocol buffer of Ports
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()
	// register the shippingPortsProtocolServer type/methods to the grpcServer
	shippingportsprotocol.RegisterShippingPortsServerServer(grpcServer, &shippingPortsProtocolServer{})

	go func() {
		// this will block the service open until some severe failure
		if err = grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return
	}
}



