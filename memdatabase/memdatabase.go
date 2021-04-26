// Package memdatabase implements a simple memory key value storage.
// Values must be encoded as slice of bytes. Keys must be string.
// One should get a new Memdatabase instance with memdatabase.New()
// It does not have any default persistence mechanism.
// Clients may use gobs of data to persist it.
package memdatabase

import (
	"errors"
	"fmt"
	"sync"
)

// Memdatabase is a map serving as stub for a key value pairs Memdatabase, such as Memcache.
// We use byte slice as the best general format for RPC struct types marshalling.
// The key string is the id of port.
type Memdatabase struct {
	// store have the data
	store map[string][]byte

	// index keys by position (order of arrival)
	index []string

	// reverseIndex have the key position in index to facilitate removal
	reverseIndex map[string]int

	// mu is a mutex to lock access to database while it is being accessed by one of the following functions
	mu sync.Mutex

	// max size is the maximum size in bytes of objects. it does not affect max quantity of elements.
	// you must be watching for overall index size also.
	max int

	// size is the actual size in bytes of objects.
	size int
}

// New returns a new instance of a Memdatabase
func New(maxsize int) *Memdatabase {
	return &Memdatabase{
		store:        make(map[string][]byte),
		index:        make([]string, 0, 1024),
		reverseIndex: make(map[string]int),
		max: maxsize,
	}
}

// Put store value under key index
func (db *Memdatabase) Put(key string, value []byte) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	// check if we have space for it
	dbsize := db.size

	// if key exists, discount it's present size
	if b, ok := db.store[key]; ok {
		dbsize = dbsize - len(b)
	}

	if dbsize+len(value) > db.max {
		return errors.New(fmt.Sprintf("storing this value will exceed max size. free space for this key is %d and for new values is %d", db.max-dbsize, db.max-db.size))
	}

	// store
	db.store[key] = value

	// index
	db.index = append(db.index, key)

	// reverseIndex
	db.reverseIndex[key] = len(db.index) - 1

	return nil
}

// Get retrieves the value under key, or returns error not nil if not found
func (db *Memdatabase) Get(key string) ([]byte, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	value, ok := db.store[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return value, nil
}

// Delete removes the key and value from Memdatabase, but does not verify the presence of the key before
// Returning an error could be misleading
// Delete don't remove key from index, only empty it's reference
// The reindexing of the database should be done eventually for consistency
func (db *Memdatabase) Delete(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	// delete form store
	delete(db.store, key)
	// get index position
	i := db.reverseIndex[key]
	// empty index position
	db.index[i] = ""
	// remove from reverseIndex also
	delete(db.reverseIndex, key)
}

func (db *Memdatabase) GetMany(offset int, count int) ([][]byte, int, bool, error) {
	var result = make([][]byte, 0)

	// check if offset is inside length of slice
	if offset > len(db.index)-1 {
		return result, 0, false, errors.New("offset out of database bounds")
	}

	if offset < 0 {
		offset = 0
	}

	if count < 1 {
		count = 100
	}

	var skips int
	for i := offset; i < offset+count+skips; i++ {
		// don't try to read past end
		if i >= len(db.index) {
			break
		}
		// if index position was deleted, skip it
		if db.index[i] == "" {
			skips++
			continue
		}
		// get the key from index
		key := db.index[i]

		// get the object from store
		if value, ok := db.store[key]; ok {
			result = append(result, value)
		} else {
			return result, len(result), offset+len(result) < len(db.store), errors.New("database index corrupted, please reindex")
		}
	}

	return result, len(result), offset+len(result) < len(db.store), nil
}
