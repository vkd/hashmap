package hashmap

import "errors"

// Key - type of key of hashmap
type Key interface{}

// HashMaper - interface for hashmap struct
type HashMaper interface {
	Set(key Key, value interface{}) error
	Get(key Key) (value interface{}, err error)
	Unset(key Key) error
	Count() int
}

// HashFunc - hash function
type HashFunc func(blockSize uint, key Key) (uint, error)

// ErrKeyNotFound - error if not found key on hashmap
var ErrKeyNotFound = errors.New("key not found")
