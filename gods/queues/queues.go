// Package queues
package queues

import "github.com/MorePeanuts/algorithm/gods/containers"

type Queue[T comparable] interface {
	Enqueue(value T)
	Dequeue() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
	// IsEmpty() bool
	// Len() int
	// Clear()
}
