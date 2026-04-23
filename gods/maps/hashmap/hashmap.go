// Package hashmap provides a generic hash map implementation backed by Go's built-in map.
package hashmap

import "github.com/MorePeanuts/algorithm/gods/maps"

// Interface guard
var _ maps.Map[string, int] = (*Map[string, int])(nil)

// Map is a generic hash map that stores key-value pairs.
// Keys must be comparable, and values can be any type.
type Map[K comparable, V any] struct {
	m map[K]V
}

// New creates and returns a new empty Map.
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{make(map[K]V)}
}

// maps.Map interface implementation

// Put inserts or updates a key-value pair in the map.
func (m *Map[K, V]) Put(key K, value V) {
	m.m[key] = value
}

// Get retrieves the value associated with the given key.
// It returns the value and a boolean indicating whether the key was found.
func (m *Map[K, V]) Get(key K) (value V, found bool) {
	value, found = m.m[key]
	return
}

// Remove deletes the key-value pair with the specified key from the map.
func (m *Map[K, V]) Remove(key K) {
	delete(m.m, key)
}

// Keys returns a slice containing all keys in the map.
// The order of keys is not specified.
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, m.Size())
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice containing all values in the map.
// The order of values is not specified.
func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, m.Size())
	for _, v := range m.m {
		values = append(values, v)
	}
	return values
}

// containers.Container interface implementation

// IsEmpty returns true if the map contains no key-value pairs.
func (m *Map[K, V]) IsEmpty() bool {
	return len(m.m) == 0
}

// Size returns the number of key-value pairs in the map.
func (m *Map[K, V]) Size() int {
	return len(m.m)
}

// Clear removes all key-value pairs from the map.
func (m *Map[K, V]) Clear() {
	clear(m.m)
}
