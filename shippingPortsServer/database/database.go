// Package database implements a simple memory key value storage.
// One should get a new database instance with database.New()
package database

import (
	"errors"
	"fmt"
)

// ShippingPortsDatabase is a map serving as stub for a key value pairs database, such as Memcache.
// We use byte slice as the best general format for RPC struct types marshalling.
// The key string is the id of port.
type ShippingPortsDatabase struct {
	store map[string][]byte
}

// New returns a new instance of a shippingPorts database
func New() *ShippingPortsDatabase {
	return &ShippingPortsDatabase{store: make(map[string][]byte)}
}

// Put store value under key index
func (spd *ShippingPortsDatabase) Put(key string, value *[]byte) error {
	spd.store[key] = *value
	// this is just a self check, should be impossible have this error
	if _, ok := spd.store[key]; !ok {
		return errors.New(fmt.Sprintf("could not store value under key %s", key))
	}
	return nil
}

// Get retrieves the value under key, or returns error not nil if not found
func (spd *ShippingPortsDatabase) Get(key string) (*[]byte, error) {
	value, ok := spd.store[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return &value, nil
}

// Delete removes the key and value from database, but does not verify the presence of the key before
func (spd *ShippingPortsDatabase) Delete(key string) {
	delete(spd.store, key)
}
