package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// portsServerAddress is the default portsProtoServer address, but it may be changed by SERVE_AT_ADDRESS env variable.
var portsServerAddress = ":50051"

// portsDatabase is a map serving as stub for a key value pairs database, such as Memcache.
// We use byte slice as the best general format for RPC struct types marshalling.
// The key string is the id of port.
var portsDatabase = make(map[string][]byte)

func init() {
	// just check if non default portsProtoServer port was given in the environment (usually Dockerfile or docker-compose.yaml)
	if envPortsServerAddress, ok := os.LookupEnv("SERVE_AT_ADDRESS"); ok {
		portsServerAddress = envPortsServerAddress
	}
}

// type portsProtoServer implements (embeds) the ports gRPC portsProtoServer interface.
// We need to give it the ports.UnimplementedPortsServer to implement the expected methods.
type portsProtoServer struct {
	ports.UnimplementedPortsServer
}

// Put to portsDatabase a port object, overwriting if exists. In this case we marshal to binary format before storing,
// so to disable lockers present in the object. It also allows us to recover the object state on Get.
// The actual portsDatabase is map[string][]byte and the Key is the port Id.
// In case we can't store the Port, we return the error causing it.
// TODO: implement context reactions on Put method.
func (s *portsProtoServer) Put(ctx context.Context, port *ports.Port) (*ports.Ok, error) {
	// byteEncoded is the []byte representation of the port.
	// in case of error, the server will respond with the Ok{} stub, but with error.
	byteEncodedPort, err := proto.Marshal(port)
	if err != nil {
		return nil, err
	}

	// save it, overwrite if already exists.
	portsDatabase[port.GetId()] = byteEncodedPort

	// very well, respond Ok, no errors.
	return &ports.Ok{}, nil
}

// Get from database a port object by the Port Id. Unmarshal and respond.
// TODO: implement context reactions on Get method.
func (s *portsProtoServer) Get(ctx context.Context, id *ports.Id) (*ports.Port, error) {
	// retrienve the marshaled binary for the Port Id.
	// if not in the map/db, respond with nil and error
	byteEncodedPort, ok := portsDatabase[id.GetId()]
	if !ok {
		return nil, errors.New("not found")
	}

	// lets unmarshal the []byte back to a port object.
	// in case of error, respond with the error.
	var port ports.Port
	err := proto.Unmarshal(byteEncodedPort, &port)
	if err != nil {
		return nil, err
	}

	// great, respond with the retrieved port.
	return &port, nil
}



func main() {
	// open the portsServerAddress to the network. It will fail if the network port is already in use. It will not retry.
	listener, err := net.Listen("tcp", portsServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// success, log it the moment
	log.Printf("portsProtoServer listening on port: %v", portsServerAddress)
	
	// get a new generic grpc server that we will attach the protocol buffer of Ports
	grpcServer := grpc.NewServer()
	
	// register the portsProtoServer type/methods to the grpcServer
	ports.RegisterPortsServer(grpcServer, &portsProtoServer{})

	// this will block the service open until some severe failure
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
