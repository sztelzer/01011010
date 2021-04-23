package main

import (
	"context"
	"log"
	"net"
	"os"
	
	"github.com/sztelzer/01011010/shippingportsserver/database"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// shippingPortsServerAddress is the default portsProtoServer address, but it may be changed by SERVE_AT_ADDRESS env variable.
var shippingPortsServerAddress = ":50051"

// shippingPortsDatabase is the database instance that will be used during the life time of the server.
// It is a in memory database and it does not persist state.
var shippingPortsDatabase = database.New()

func init() {
	// just check if non default portsProtoServer port was given in the environment (usually Dockerfile or docker-compose.yaml)
	if envPortsServerAddress, ok := os.LookupEnv("SERVE_AT_ADDRESS"); ok {
		shippingPortsServerAddress = envPortsServerAddress
	}
}

// type portsProtoServer implements (embeds) the ports gRPC portsProtoServer interface.
// We need to give it the ports.UnimplementedPortsServer to implement the expected methods.
type portsProtoServer struct {
	shippingportsprotocol.UnimplementedShippingPortsServerServer
}

// Put to shippingPortsDatabase a port object, overwriting if exists. In this case we marshal to binary format before storing,
// so to disable lockers present in the object. It also allows us to recover the object state on Get.
// The actual shippingPortsDatabase is map[string][]byte and the Key is the port Id.
// In case we can't store the Port, we return the error causing it.
// TODO: implement context reactions on Put method.
func (s *portsProtoServer) Put(ctx context.Context, shippingPort *shippingportsprotocol.ShippingPort) (*shippingportsprotocol.Ok, error) {
	// byteEncoded is the []byte representation of the port.
	// in case of error, the server will respond with the Ok{} stub, but with error.
	byteEncodedShippingPort, err := proto.Marshal(shippingPort)
	if err != nil {
		return nil, err
	}
	
	// save it, overwrite if already exists.
	err = shippingPortsDatabase.Put(shippingPort.GetId(), &byteEncodedShippingPort)
	if err != nil {
		return nil, err
	}

	// very well, respond Ok, no errors.
	return &shippingportsprotocol.Ok{}, nil
}

// Get from database a port object by the Port Id. Unmarshal and respond.
// TODO: implement context reactions on Get method.
func (s *portsProtoServer) Get(ctx context.Context, shippingPortId *shippingportsprotocol.ShippingPortId) (*shippingportsprotocol.ShippingPort, error) {
	// retrieve the marshaled binary for the Port Id.
	// if not in the map/db, respond with nil and error
	byteEncodedShippingPort, err := shippingPortsDatabase.Get(shippingPortId.GetId())
	if err != nil {
		return nil, err
	}
	
	// lets unmarshal the []byte back to a port object.
	// in case of error, respond with the error.
	var shippingPort shippingportsprotocol.ShippingPort
	err = proto.Unmarshal(*byteEncodedShippingPort, &shippingPort)
	if err != nil {
		return nil, err
	}
	
	// great, respond with the retrieved port.
	return &shippingPort, nil
}

func main() {
	// open the shippingPortsServerAddress to the network. It will fail if the network port is already in use. It will not retry.
	listener, err := net.Listen("tcp", shippingPortsServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// success, log it the moment
	log.Printf("portsProtoServer listening on port: %v", shippingPortsServerAddress)
	
	// get a new generic grpc server that we will attach the protocol buffer of Ports
	grpcServer := grpc.NewServer()
	
	// register the portsProtoServer type/methods to the grpcServer
	// TODO: remove 'Server' from service definition? For Service?
	shippingportsprotocol.RegisterShippingPortsServerServer(grpcServer, &portsProtoServer{})
	
	// this will block the service open until some severe failure
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
