package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/sztelzer/01011010/shippingportsprotocol"
	"google.golang.org/protobuf/encoding/protojson"
)

// mainHandler checks the path requested, extracts the Shipping Port Id and get it from the server.
func mainHandler(shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			log.Println("GET", r.URL.Path)
			
			shippingPortId, err := extractShippingPortId(r.URL.Path)
			if err != nil {
				log.Println(err)
				http.NotFound(w, r)
				return
			}
			
			// get the shippingPort from server
			shippingPort, err := shippingPortsServerClient.Get(context.Background(), &shippingportsprotocol.ShippingPortId{Id: shippingPortId})
			if err != nil {
				log.Printf("error getting %s %+v", shippingPortId, err)
				http.NotFound(w, r)
				return
			}
			
			// marshal shippingPort to JSON
			shippingPortJsonBytes, err := protojson.Marshal(shippingPort)
			if err != nil {
				http.Error(w, "error marshalling response json", http.StatusInternalServerError)
				return
			}
			
			w.Header().Add("Content-Type", "application/json")
			writtenCount, err := w.Write(shippingPortJsonBytes)
			if err != nil {
				log.Printf("error writing %d bytes %+v", writtenCount, err)
			}
		} else {
			http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

// extractShippingPortId from endpoint matching /shippingports/IDXX
func extractShippingPortId (path string) (string, error) {
	// extract something resembling a shippingPortId from the url path under endpoint
	urlEndpointFilter := regexp.MustCompile(`(?m)^\/shippingports\/`)
	if !urlEndpointFilter.MatchString(path) {
		return "", errors.New("not found")
	}
	shippingPortId := urlEndpointFilter.ReplaceAllString(path, "")
	shippingPortIdFilter := regexp.MustCompile(`[^a-zA-Z0-9]`)
	shippingPortId = shippingPortIdFilter.ReplaceAllString(shippingPortId, "")
	shippingPortId = strings.ToUpper(shippingPortId)
	
	return shippingPortId, nil
}