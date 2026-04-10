// Package lists
package lists

import "github.com/MorePeanuts/algorithm/gods/containers"

type List[T comparable] interface {
	Get(idx int) (T, bool)
	Set(idx int, value T)
	Append(values ...T)
	Remove(idx int)
	Insert(idx int, values ...T)
	Contains(values ...T) bool
	Swap(i, j int)

	containers.Container[T]
	// IsEmpty() bool
	// Len() int
	// Clear()
}
