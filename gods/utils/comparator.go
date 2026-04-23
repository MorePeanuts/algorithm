// Package utils provides utility functions and types for the gods package.
package utils

// Comparator is a function type that compares two values of type T.
// It returns a negative number if x < y, zero if x == y, and a positive number if x > y.
type Comparator[T any] func(x, y T) int
