package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/sztelzer/01011010/shippingportsprotocol"
	"google.golang.org/protobuf/encoding/protojson"
)

// mainHandler checks the path requested, extracts the Shipping Port Id and get it from the server.
func mainHandler(shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			log.Println("GET", r.URL.Path)

			var content []byte
			var err error

			urlEndpointFilter := regexp.MustCompile(`^\/shippingports$`)
			if urlEndpointFilter.MatchString(r.URL.Path) {
				content, err = getManyShippingPorts(r.URL, shippingPortsServerClient)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Add("Content-Type", "application/json")
				_, err = w.Write(content)
				log.Println(err)
				return
			}

			urlEndpointFilter = regexp.MustCompile(`^\/shippingports\/`)
			if urlEndpointFilter.MatchString(r.URL.Path) {
				content, err = getShippingPort(r.URL, shippingPortsServerClient)
				if err != nil {
					if err.Error() == "rpc error: code = Unknown desc = not found" {
						http.Error(w, "not found", http.StatusNotFound)
						return
					}

					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Header().Add("Content-Type", "application/json")
				_, err = w.Write(content)
				log.Println(err)
				return

			}

			http.Error(w, "not found", http.StatusNotFound)

		} else {
			http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		}
	}
}

// extractShippingPortId from endpoint matching /shippingports/IDXX
func extractShippingPortId(path string) string {
	// extract something resembling a shippingPortId from the url path under endpoint
	urlEndpointFilter := regexp.MustCompile(`(?m)^\/shippingports\/`)
	shippingPortId := urlEndpointFilter.ReplaceAllString(path, "")
	shippingPortIdFilter := regexp.MustCompile(`[^a-zA-Z0-9]`)
	shippingPortId = shippingPortIdFilter.ReplaceAllString(shippingPortId, "")
	shippingPortId = strings.ToUpper(shippingPortId)

	return shippingPortId
}

func getShippingPort(url *url.URL, shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient) ([]byte, error) {
	id := extractShippingPortId(url.Path)
	// get the shippingPort from server
	shippingPort, err := shippingPortsServerClient.Get(context.Background(), &shippingportsprotocol.ShippingPortId{Id: id})
	if err != nil {
		return []byte(""), err
	}

	// marshal shippingPort to JSON
	shippingPortJsonBytes, err := protojson.Marshal(shippingPort)
	if err != nil {
		return []byte(""), err
	}

	return shippingPortJsonBytes, nil
}

func getManyShippingPorts(url *url.URL, shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient) ([]byte, error) {
	var size int32 = 100
	var page int32 = 1

	q := url.Query()
	qSize := q.Get("size")
	qPage := q.Get("page")

	if qSize != "" {
		qSizeInt, err := strconv.Atoi(qSize)
		if err == nil {
			size = int32(qSizeInt)
			if size < 1 {
				size = 1
			}
		}
	}

	if qPage != "" {
		qPageInt, err := strconv.Atoi(qPage)
		if err == nil {
			page = int32(qPageInt)
			if page < 1 {
				page = 1
			}
		}
	}

	// get the shippingPort from server
	shippingPorts, err := shippingPortsServerClient.GetMany(context.Background(), &shippingportsprotocol.Pagination{
		Offset: (page - 1) * size,
		Size:   size,
	})
	if err != nil {
		return []byte(""), err
	}

	// marshal shippingPort to JSON
	shippingPortsJsonBytes, err := protojson.Marshal(shippingPorts)
	if err != nil {
		return []byte(""), err
	}

	return shippingPortsJsonBytes, nil
}
