// Package memdatabase implements a simple memory key value storage.
// One should get a new memdatabase instance with shippingPortsMemDatabase.New()
package shippingportsmemdatabase

import (
	"errors"
	"sync"
)

// ShippingPortsDatabase is a map serving as stub for a key value pairs memdatabase, such as Memcache.
// We use byte slice as the best general format for RPC struct types marshalling.
// The key string is the id of port.
type ShippingPortsDatabase struct {
	// store have the data
	store map[string][]byte
	// index keys by position (order of arrival)
	index []string
	// reverseIndex have the key position in index to facilitate removal
	reverseIndex map[string]int
	// mu is a mutex to lock access to database while it is being accessed by one of the following functions
	mu sync.Mutex
}

// New returns a new instance of a shippingPorts memdatabase
func New() *ShippingPortsDatabase {
	return &ShippingPortsDatabase{
		store: make(map[string][]byte),
		index: make([]string, 0, 1024),
		reverseIndex: make(map[string]int),
	}
}

// Put store value under key index
func (spd *ShippingPortsDatabase) Put(key string, value *[]byte) {
	spd.mu.Lock()
	defer spd.mu.Unlock()

	// store
	spd.store[key] = *value

	// index
	spd.index = append(spd.index, key)

	// reverseIndex
	spd.reverseIndex[key] = len(spd.index)
}

// Get retrieves the value under key, or returns error not nil if not found
func (spd *ShippingPortsDatabase) Get(key string) (*[]byte, error) {
	spd.mu.Lock()
	defer spd.mu.Unlock()
	value, ok := spd.store[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return &value, nil
}

// Delete removes the key and value from memdatabase, but does not verify the presence of the key before
// Returning an error could be misleading
func (spd *ShippingPortsDatabase) Delete(key string) {
	spd.mu.Lock()
	defer spd.mu.Unlock()
	delete(spd.store, key)
}


