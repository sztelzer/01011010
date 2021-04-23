package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
	"google.golang.org/grpc"
	
	"google.golang.org/protobuf/encoding/protojson"
)


// shippingPortsServerAddress is the default connection string to the gRPC shippingPortsServer
// It may be changed by the environment variable SERVER_ADDRESS
var shippingPortsServerAddress = "shippingportsserver:50051"

// shippingPortsClientAddress is the default port to serve the REST API
var shippingPortsClientAddress = ":8080"

func init() {
	// just check if non default shippingPortsServerAddress address was given in the environment (usually Dockerfile or docker-compose.yaml)
	if envPortsServerAddress, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		shippingPortsServerAddress = envPortsServerAddress
	}
	
	if envPortsClientAddress, ok := os.LookupEnv("SERVE_CLIENT_AT_ADDRESS"); ok {
		shippingPortsClientAddress = envPortsClientAddress
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
	go saveShippingPortsFromFile("dropbox/ports.json", shippingPortsServerClient)
	
	// Serve the REST API
	// Register a handler for requests
	http.HandleFunc("/", mainHandler(shippingPortsServerClient))

	// Start server
	log.Fatal(http.ListenAndServe(shippingPortsClientAddress, nil))
}


// mainHandler checks the path requested, extracts the Shipping Port Id and get it from the server.
func mainHandler (shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			
			// extract something resembling a shippingPortId from the url path under endpoint
			urlEndpointFilter := regexp.MustCompile("^/shippingports/")
			if !urlEndpointFilter.MatchString(r.URL.Path) {
				http.NotFound(w, r)
				return
			}
			
			shippingPortId := urlEndpointFilter.ReplaceAllString(r.URL.Path, "")
			shippingPortIdFilter := regexp.MustCompile("[^a-zA-Z0-9]")
			shippingPortId = shippingPortIdFilter.ReplaceAllString(r.URL.Path, "")
			shippingPortId = strings.ToUpper(shippingPortId)
			
			// try to get the shippingPort from server
			shippingPort, err := shippingPortsServerClient.Get(context.Background(), &shippingportsprotocol.ShippingPortId{Id: shippingPortId})
			if err != nil {
				http.NotFound(w, r)
				return
			}
			
			shippingPortJsonBytes, err := protojson.Marshal(shippingPort)
			if err != nil {
				http.Error(w, "error marshalling response json", http.StatusInternalServerError)
				return
			}
			
			w.Header().Add("Content-Type", "application/json")
			w.Write(shippingPortJsonBytes)
		} else {
			http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		}
	}
}


// savePortsFromFile reads a file of objects and saves each to server
func saveShippingPortsFromFile(filename string, shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
		return
	}
	
	reader := bufio.NewReader(f)
	_, err = reader.ReadBytes('\n') // discard first line with the opening {
	if err != nil {
		log.Fatalln(err)
	}
	
	var putCount int
	
	// lets range over blocks of lines that represents each shippingPort in the json
	for {
		nextShippingPort, err := ReadNextShippingPort(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			continue
		}
		_, err = shippingPortsServerClient.Put(context.Background(), nextShippingPort)
		if err != nil {
			log.Println(err)
			continue
		}
		putCount++
	}
	
	log.Println(putCount)
}


// ReadNextShippingPort advances one block of lines that represents each shippingPort and returns it
// As reader is a pointer, the position is stateful
func ReadNextShippingPort(reader *bufio.Reader) (*shippingportsprotocol.ShippingPort, error) {
	
	firstLine, err := reader.ReadBytes(':')
	if err != nil {
		return nil, err
	}
	shippingPortCodeFilter := regexp.MustCompile("[^A-Z0-9]")
	id := shippingPortCodeFilter.ReplaceAllString(string(firstLine), "")
	
	block, err := reader.ReadBytes('}')
	if err != nil {
		return nil, err
	}
	
	var shippingPort shippingportsprotocol.ShippingPort
	err = json.Unmarshal(block, &shippingPort)
	if err != nil {
		return nil, err
	}
	shippingPort.Id = id
	
	return &shippingPort, nil
}
