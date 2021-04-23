package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
	"sync/atomic"
	"time"
	
	"github.com/sztelzer/01011010/shippingportsprotocol"
)

// savePortsFromFile reads a file of objects and saves each to server
// TODO: segment file and run multiple readers
func saveShippingPortsFromFile(filename string, shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient) {
	startTime := time.Now()
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
		return
	}
	// we will use only one reader, but could be many
	reader := bufio.NewReader(f)
	
	// discard first line
	_, err = reader.ReadBytes('\n')
	if err != nil {
		log.Fatalln(err)
	}
	
	var putCount int32
	
	var waitGroup sync.WaitGroup
	var shippingPortsChan = make(chan *shippingportsprotocol.ShippingPort, 256)
	
	// spawn multiple loaders
	for loaders := 0; loaders < 8; loaders++ {
		go func() {
			for {
				select {
				case shippingPort, more := <-shippingPortsChan:
					_, err = shippingPortsServerClient.Put(context.Background(), shippingPort)
					if err != nil {
						log.Println(err)
						continue
					}
					atomic.AddInt32(&putCount, 1)
					waitGroup.Done()
					if !more {
						return
					}
				}
			}
		}()
	}
	
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
		
		waitGroup.Add(1)
		shippingPortsChan <- nextShippingPort
	}
	
	waitGroup.Wait()
	
	log.Printf("successfully loaded %d shippingPorts from file to server in %v", putCount, time.Now().Sub(startTime))
}

// ReadNextShippingPort advances one block of lines that represents each shippingPort and returns it
// As reader is a pointer, the position is stateful
// TODO: better handling of block closures
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
