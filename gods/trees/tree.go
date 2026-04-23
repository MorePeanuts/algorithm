// Package trees
package trees

import (
	"github.com/MorePeanuts/algorithm/gods/containers"
)

type Tree[T any] interface {
	containers.Container[T]
	// IsEmpty() bool
	// Size() int
	// Clear()
}

// TODO: Interfaces need to be designed for trees with different purposes.
