package main

import (
	"context"
	"log"

	"github.com/sztelzer/01011010/shippingportsprotocol"
	"google.golang.org/protobuf/proto"
)

// type shippingPortsProtocolServer implements (embeds) the ports gRPC shippingPortsProtocolServer interface.
// We need to give it the ports.UnimplementedPortsServer to implement the expected methods.
type shippingPortsProtocolServer struct {
	shippingportsprotocol.UnimplementedShippingPortsServerServer
}

// Put to shippingPortsDatabase a port object, overwriting if exists. In this case we marshal to binary format before storing,
// so to disable lockers present in the object. It also allows us to recover the object state on Get.
// The actual shippingPortsDatabase is map[string][]byte and the Key is the port Id.
// In case we can't store the Port, we return the error causing it.
func (s *shippingPortsProtocolServer) Put(ctx context.Context, shippingPort *shippingportsprotocol.ShippingPort) (*shippingportsprotocol.Ok, error) {
	// byteEncoded is the []byte representation of the port.
	// in case of error, the server will respond with the Ok{} stub, but with error.
	byteEncodedShippingPort, err := proto.Marshal(shippingPort)
	if err != nil {
		return nil, err
	}

	// we cannot overwrite with older versions of the object
	// so we must get it first and compare order
	// TODO: this should relaid to the database capabilities of update if
	ExistentByteEncodedShippingPort, err := shippingPortsDatabase.Get(shippingPort.GetId())
	var ExistentShippingPort shippingportsprotocol.ShippingPort
	if err == nil {
		err = proto.Unmarshal(ExistentByteEncodedShippingPort, &ExistentShippingPort)
		if err != nil {
			return nil, err
		}
		// if we already have the newer version return
		if ExistentShippingPort.GetOrder() > shippingPort.GetOrder() {
			return &shippingportsprotocol.Ok{}, nil
		}
	}

	// save it, overwrite.
	err = shippingPortsDatabase.Put(shippingPort.GetId(), byteEncodedShippingPort)
	if err != nil {
		return nil, err
	}

	// very well, respond Ok, no errors.
	return &shippingportsprotocol.Ok{}, nil
}

// Get from memdatabase a shippingPort object by the Port Id. Unmarshal and respond.
func (s *shippingPortsProtocolServer) Get(ctx context.Context, shippingPortId *shippingportsprotocol.ShippingPortId) (*shippingportsprotocol.ShippingPort, error) {
	// retrieve the marshaled binary for the Port Id.
	// if not in the map/db, respond with nil and error
	byteEncodedShippingPort, err := shippingPortsDatabase.Get(shippingPortId.GetId())
	if err != nil {
		return nil, err
	}

	// lets unmarshal the []byte back to a port object.
	// in case of error, respond with the error.
	var shippingPort shippingportsprotocol.ShippingPort
	err = proto.Unmarshal(byteEncodedShippingPort, &shippingPort)
	if err != nil {
		return nil, err
	}

	// great, respond with the retrieved port.
	return &shippingPort, nil
}

// GetMany items from database with offset and size
func (s *shippingPortsProtocolServer) GetMany(ctx context.Context, pagination *shippingportsprotocol.Pagination) (*shippingportsprotocol.ManyShippingPorts, error) {
	// retrieve the marshaled binary for the Port Id.
	// if not in the map/db, respond with nil and error
	byteEncodedShippingPorts, n, more, err := shippingPortsDatabase.GetMany(int(pagination.GetOffset()), int(pagination.GetSize()))
	if err != nil {
		return nil, err
	}

	// result decoded array
	shippingPorts := make([]*shippingportsprotocol.ShippingPort, 0, pagination.GetSize())
	for _, byteEncodedShippingPort := range byteEncodedShippingPorts {
		// lets unmarshal the []byte back to a port object.
		var shippingPort shippingportsprotocol.ShippingPort
		err = proto.Unmarshal(byteEncodedShippingPort, &shippingPort)
		if err != nil {
			// in case of error, just log and continue
			log.Println(err)
			continue
		}

		shippingPorts = append(shippingPorts, &shippingPort)
	}

	manyShippingPorts := shippingportsprotocol.ManyShippingPorts{
		Shippingports: shippingPorts,
		Count:         int32(n),
		More:          more,
	}

	// great, respond with the retrieved port.
	return &manyShippingPorts, nil
}

