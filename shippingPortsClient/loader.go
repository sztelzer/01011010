package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
	"sync/atomic"
	"time"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
	"google.golang.org/protobuf/encoding/protojson"
)

// savePortsFromFile reads a file of objects and saves each to server
// TODO: segment file and run multiple readers
func loadShippingPortsFromFileToServer(ctx context.Context, shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient) {
	startTime := time.Now()
	
	f, err := os.Open(shippingPortsOriginJsonFile)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer f.Close()
	
	// we will use only one reader, but could be many
	reader := bufio.NewReader(f)
	
	// discard first line
	_, err = reader.ReadBytes('\n')
	if err != nil {
		log.Fatalln(err)
	}
	
	var putCount int32
	
	var waitGroup sync.WaitGroup
	var shippingPortsChan = make(chan *shippingportsprotocol.ShippingPort, 512)
	
	// spawn multiple loaders
	for putters := 0; putters < shippingPortsPutters; putters++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case shippingPort, more := <-shippingPortsChan:
					if shippingPort != nil {
						_, err = shippingPortsServerClient.Put(ctx, shippingPort)
						if err != nil {
							log.Println(err)
							waitGroup.Done()
							continue
						}
						atomic.AddInt32(&putCount, 1)
						waitGroup.Done()
					}
					if !more {
						return
					}
				}
			}
		}()
	}
	
	// lets range over blocks of lines that represents each shippingPort in the json
	for {
		if ctx.Err() != nil {
			return
		}
		nextShippingPort, err := readNextShippingPort(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			continue
		}
		
		waitGroup.Add(1)
		shippingPortsChan <- nextShippingPort
	}
	
	waitGroup.Wait()
	close(shippingPortsChan)
	
	log.Printf("successfully loaded %d shippingPorts from file to server in %v", putCount, time.Now().Sub(startTime))
}

// readNextShippingPort advances one block of lines that represents each shippingPort and returns it
// As reader is a pointer, the position is stateful
// TODO: better handling of block closures
func readNextShippingPort(reader *bufio.Reader) (*shippingportsprotocol.ShippingPort, error) {
	
	firstLine, err := reader.ReadBytes(':')
	if err != nil {
		return nil, err
	}
	shippingPortCodeFilter := regexp.MustCompile("[^A-Z0-9]")
	id := shippingPortCodeFilter.ReplaceAllString(string(firstLine), "")
	
	// read each rune and control if is inside string and same level

	
	
	
	
	
	
	block, err := reader.ReadBytes('}')
	if err != nil {
		return nil, err
	}
	
	var shippingPort shippingportsprotocol.ShippingPort
	
	err = protojson.Unmarshal(block, &shippingPort)
	if err != nil {
		return nil, err
	}
	shippingPort.Id = id
	
	return &shippingPort, nil
}
