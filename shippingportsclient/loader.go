package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sztelzer/01011010/shippingportsclient/blockreader"
	"github.com/sztelzer/01011010/shippingportsprotocol"
	"google.golang.org/protobuf/encoding/protojson"
)

// savePortsFromFile reads a file of objects and saves each to server
// TODO: segment file and run multiple readers
func loadShippingPortsFromFileToServer(ctx context.Context, shippingPortsServerClient shippingportsprotocol.ShippingPortsServerClient) {
	// will use to time the process
	startTime := time.Now()

	// open the file (set by env)
	f, err := os.Open(shippingPortsOriginJsonFile)
	if err != nil {
		log.Println(err)
		// if we can't open the file, don't continue
		return
	}
	defer f.Close()

	// we will use only one reader, but could be many
	reader := bufio.NewReader(f)

	// var counter will have the position of reader to use as update order
	var counter int

	// jump to inside of external block
	discard, err := reader.ReadBytes('{')
	if err != nil {
		log.Println(err)
		// if we can't go to inside, don't continue
		return
	}
	counter = len(discard)

	// now we can set some vars

	// count items saved
	var putCount int32

	// use waiting to guarantee putting every shippingPort on channel
	var waitGroup sync.WaitGroup

	// use buffered channel to control different speeds in reading and putting
	var shippingPortsChan = make(chan *shippingportsprotocol.ShippingPort, 512)

	// spawn multiple putters
	for putters := 0; putters < shippingPortsPutters; putters++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					// stop this putter if context is canceled
					return
				case shippingPort, more := <-shippingPortsChan:
					if shippingPort != nil {
						// put to server
						_, err = shippingPortsServerClient.Put(ctx, shippingPort)
						if err != nil {
							log.Println(err)
							waitGroup.Done()
							continue
						}
						// increment counter properly
						atomic.AddInt32(&putCount, 1)
						// remove 1 from waiting
						waitGroup.Done()
					}
					if !more {
						// stop this putter if channel is closed
						return
					}
				}
			}
		}()
	}

	// lets range over blocks that represents each shippingPort in the json
	for {
		// stop everything if we are
		if ctx.Err() != nil {
			return
		}

		// get next valid block.
		// it will continue retrying fit an object right if some error happens in reading.
		// if we reach end of file, stop readings
		var nextShippingPort *shippingportsprotocol.ShippingPort
		nextShippingPort, counter, err = readNextShippingPort(reader, counter)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("bytes read: %d", counter)
			log.Println(err)
			continue
		}

		// add the next shippingPort to the channel and wait group
		waitGroup.Add(1)
		shippingPortsChan <- nextShippingPort
	}

	// wait all complete
	waitGroup.Wait()
	// close the channel properly, this signal putters to stop
	close(shippingPortsChan)

	log.Printf("successfully loaded %d shippingPorts from file to server in %v", putCount, time.Now().Sub(startTime))
}

// readNextShippingPort advances one block of lines that represents each shippingPort and returns it
// As reader is a pointer, the position is stateful
func readNextShippingPort(reader *bufio.Reader, counter int) (*shippingportsprotocol.ShippingPort, int, error) {
	// we should be near next property as shippingPort id
	var id []byte
	var err error

	id, counter, err = blockreader.NextBlock(reader, '"', 16, counter)
	if err != nil {
		return nil, counter, err
	}

	updateOrder := counter - len(id)

	// after id must be the property block
	var block []byte
	block, counter, err = blockreader.NextBlock(reader, '{', 512, counter)
	if err != nil {
		return nil, counter, err
	}

	// unmarshal to protocol ShippingPort using the protojson unmarshal function
	var shippingPort shippingportsprotocol.ShippingPort
	err = protojson.Unmarshal(block, &shippingPort)
	if err != nil {
		return nil, counter, err
	}
	// set the outside id on the type
	shippingPort.Id = string(id[1 : len(id)-1])

	// set update position
	shippingPort.Order = int32(updateOrder)

	return &shippingPort, counter, nil
}
