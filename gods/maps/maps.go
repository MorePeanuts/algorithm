// Package maps
package maps

import "github.com/MorePeanuts/algorithm/gods/containers"

type Map[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) (value V, found bool)
	Remove(key K)
	Keys() []K
	Values() []V

	containers.Container[V]
	// IsEmpty() bool
	// Size() int
	// Clear()
}
