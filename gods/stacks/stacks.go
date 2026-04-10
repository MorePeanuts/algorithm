// Package stacks
package stacks

import "github.com/MorePeanuts/algorithm/gods/containers"

type Stack[T any] interface {
	Push(value T)
	Pop() (value T, ok bool)
	Peek() (value T, ok bool)

	containers.Container[T]
	// IsEmpty() bool
	// Len() int
	// Clear()
}
