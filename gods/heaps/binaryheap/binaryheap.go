// Package binaryheap provides a generic min-heap implementation.
package binaryheap

import (
	"cmp"

	"github.com/MorePeanuts/algorithm/gods/heaps"
	"github.com/MorePeanuts/algorithm/gods/utils"
)

// Ensure that Heap implements the heaps.Heap interface.
var _ heaps.Heap[int] = (*Heap[int])(nil)

// Heap represents a min-heap data structure.
// It is implemented using a slice, where for any element at index i:
//   - Left child is at index 2*i + 1
//   - Right child is at index 2*i + 2
//   - Parent is at index (i-1)/2
type Heap[T comparable] struct {
	list       []T
	Comparator utils.Comparator[T]
}

// New creates a new min-heap for ordered types.
// It uses the natural ordering of the type.
func New[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{[]T{}, cmp.Compare[T]}
}

// NewWith creates a new heap with a custom comparator.
// The comparator determines the heap ordering (min or max).
func NewWith[T comparable](comparator utils.Comparator[T]) *Heap[T] {
	return &Heap[T]{[]T{}, comparator}
}

// heaps.Heap interface implementation

// Insert adds one or more values to the heap.
// For a single value, it appends and percolates up.
// For multiple values, it uses Floyd's algorithm for O(n) heap construction.
func (heap *Heap[T]) Insert(values ...T) {
	if len(values) == 1 {
		heap.list = append(heap.list, values[0])
		heap.percolateUp()
	} else {
		// Floyd's algorithm - heapify from the bottom up
		heap.list = append(heap.list, values...)
		for i := len(heap.list)/2 - 1; i >= 0; i-- {
			heap.percolateDownAt(i)
		}
	}
}

// DeleteMin removes and returns the minimum element from the heap.
// Returns (zero value, false) if the heap is empty.
func (heap *Heap[T]) DeleteMin() (value T, ok bool) {
	if len(heap.list) == 0 {
		var zero T
		return zero, false
	}
	value, ok = heap.list[0], true
	heap.list[0] = heap.list[len(heap.list)-1]
	heap.list = heap.list[:len(heap.list)-1]
	heap.percolateDown()
	return
}

// FindMin returns the minimum element without removing it.
// Returns (zero value, false) if the heap is empty.
func (heap *Heap[T]) FindMin() (value T, ok bool) {
	if len(heap.list) == 0 {
		var zero T
		return zero, false
	}
	return heap.list[0], true
}

// Private methods implementation of binaryheap.

// percolateDownAt moves the element at the given index down the heap
// until the heap property is restored.
func (heap *Heap[T]) percolateDownAt(idx int) {
	size := len(heap.list)
	if size == 0 {
		return
	}
	idxValue := heap.list[idx]
	for childIdx := 2*idx + 1; childIdx < size; childIdx = 2*idx + 1 {
		// Find the smaller child (or larger if it's a max-heap)
		if childIdx+1 < size && heap.Comparator(heap.list[childIdx], heap.list[childIdx+1]) > 0 {
			childIdx++
		}
		// Swap with child if necessary
		if heap.Comparator(idxValue, heap.list[childIdx]) > 0 {
			heap.list[idx] = heap.list[childIdx]
		} else {
			heap.list[idx] = idxValue
			return
		}
		idx = childIdx
	}
	heap.list[idx] = idxValue
}

// percolateDown moves the root element down the heap to restore the heap property.
func (heap *Heap[T]) percolateDown() {
	heap.percolateDownAt(0)
}

// percolateUp moves the last element up the heap until the heap property is restored.
func (heap *Heap[T]) percolateUp() {
	idx := len(heap.list) - 1
	lastValue := heap.list[idx]
	for parentIdx := (idx - 1) / 2; idx > 0; parentIdx = (idx - 1) / 2 {
		parentValue := heap.list[parentIdx]
		if heap.Comparator(parentValue, lastValue) <= 0 {
			heap.list[idx] = lastValue
			return
		}
		heap.list[idx] = parentValue
		idx = parentIdx
	}
	heap.list[idx] = lastValue
}

// containers.Container interface implementation

// IsEmpty returns true if the heap contains no elements.
func (heap *Heap[T]) IsEmpty() bool {
	return len(heap.list) == 0
}

// Size returns the number of elements in the heap.
func (heap *Heap[T]) Size() int {
	return len(heap.list)
}

// Clear removes all elements from the heap.
func (heap *Heap[T]) Clear() {
	heap.list = heap.list[:0]
}
