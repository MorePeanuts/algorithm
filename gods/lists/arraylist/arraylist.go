// Package arraylist provides a generic dynamic array implementation.
package arraylist

import "slices"

// List is a generic dynamic array that stores elements of type T.
// It implements the list.List interface and containers.Container interface.
type List[T comparable] struct {
	values []T
}

const (
	// growthFactor is the factor by which the array grows when resizing
	growthFactor = float32(2.00)
	// shrinkFactor is the threshold ratio (length/capacity) below which the array shrinks
	shrinkFactor = float32(0.25)
)

// New creates a new array list with the given values (optional).
// If no values are provided, an empty list is returned.
func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Append(values...)
	}
	return list
}

// lists.List interface implementation

// Get returns the element at the specified index.
// It returns the element and true if the index is valid,
// otherwise it returns the zero value of T and false.
func (list *List[T]) Get(idx int) (T, bool) {
	if !list.withinRange(idx) {
		var t T
		return t, false
	}
	return list.values[idx], true
}

// Set updates the element at the specified index.
// If the index is equal to the list length, the value is appended.
// If the index is out of range, no action is taken.
func (list *List[T]) Set(idx int, value T) {
	if !list.withinRange(idx) {
		if idx == len(list.values) {
			list.Append(value)
		}
		return
	}
	list.values[idx] = value
}

// Append adds one or more values to the end of the list.
func (list *List[T]) Append(values ...T) {
	length := len(list.values)
	list.growBy(len(values))
	for i, value := range values {
		list.values[length+i] = value
	}
}

// Remove removes the element at the specified index.
// If the index is out of range, no action is taken.
func (list *List[T]) Remove(idx int) {
	if !list.withinRange(idx) {
		return
	}
	list.values = slices.Delete(list.values, idx, idx+1)
	list.shrink()
}

// Insert inserts one or more values at the specified index.
// If the index is equal to the list length, the values are appended.
// If the index is out of range, no action is taken.
func (list *List[T]) Insert(idx int, values ...T) {
	if !list.withinRange(idx) {
		if idx == len(list.values) {
			list.Append(values...)
		}
		return
	}
	length := len(list.values)
	list.growBy(len(values))
	list.values = slices.Insert(list.values[:length], idx, values...)
}

// Contains checks if all the specified values exist in the list.
// Returns true if all values are found, false otherwise.
// If no values are provided, returns true.
func (list *List[T]) Contains(values ...T) bool {
	for _, value := range values {
		if !slices.Contains(list.values, value) {
			return false
		}
	}
	return true
}

// Swap swaps the elements at the specified indices i and j.
// If either index is out of range or i == j, no action is taken.
func (list *List[T]) Swap(i, j int) {
	if i == j || !list.withinRange(i) || !list.withinRange(j) {
		return
	}

	list.values[i], list.values[j] = list.values[j], list.values[i]
}

// containers.Container interface implementation

// IsEmpty returns true if the list contains no elements.
func (list *List[T]) IsEmpty() bool {
	return len(list.values) == 0
}

// Len returns the number of elements in the list.
func (list *List[T]) Len() int {
	return len(list.values)
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	clear(list.values[:cap(list.values)])
	list.values = list.values[:0]
}

// withinRange checks if the given index is valid for the list.
// Returns true if idx is between 0 (inclusive) and list length (exclusive).
func (list *List[T]) withinRange(idx int) bool {
	return idx >= 0 && idx < len(list.values)
}

// resize creates a new underlying array with the specified length and capacity.
func (list *List[T]) resize(length, capacity int) {
	newValues := make([]T, length, capacity)
	copy(newValues, list.values)
	list.values = newValues
}

// growBy increases the list's length by n elements, resizing if necessary.
// Special note: This method will change the length of the array.
func (list *List[T]) growBy(n int) {
	currentCapacity := cap(list.values)

	if newLength := len(list.values) + n; newLength >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newLength, newCapacity)
	} else {
		list.values = list.values[:newLength]
	}
}

// shrink reduces the underlying array capacity if the length is below the shrink threshold.
func (list *List[T]) shrink() {
	if shrinkFactor > 0 {
		currentCapacity := cap(list.values)
		if len(list.values) <= int(float32(currentCapacity)*shrinkFactor) {
			list.resize(len(list.values), len(list.values))
		}
	}
}
