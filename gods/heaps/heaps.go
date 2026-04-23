// Package heaps
package heaps

import "github.com/MorePeanuts/algorithm/gods/containers"

type Heap[T comparable] interface {
	Insert(values ...T)
	DeleteMin() (value T, ok bool)
	FindMin() (value T, ok bool)

	containers.Container[T]
	// IsEmpty() bool
	// Size() int
	// Clear()
}
