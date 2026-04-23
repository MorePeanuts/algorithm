package priorityqueue

import (
	"cmp"

	"github.com/MorePeanuts/algorithm/gods/heaps/binaryheap"
	"github.com/MorePeanuts/algorithm/gods/utils"
)

// Queue represents a priority queue data structure.
// Elements are ordered according to their natural ordering (for Ordered types)
// or a custom comparator provided during initialization.
type Queue[T comparable] struct {
	heap *binaryheap.Heap[T]
}

// New creates a new priority queue with the given Ordered type.
// The queue orders elements according to their natural ordering (ascending).
func New[T cmp.Ordered]() *Queue[T] {
	return &Queue[T]{binaryheap.New[T]()}
}

// NewWith creates a new priority queue with the given comparator.
// The queue orders elements according to the comparator provided.
func NewWith[T comparable](comparator utils.Comparator[T]) *Queue[T] {
	return &Queue[T]{binaryheap.NewWith(comparator)}
}

// Enqueue adds an element to the priority queue.
func (queue *Queue[T]) Enqueue(value T) {
	queue.heap.Insert(value)
}

// Dequeue removes and returns the element with the highest priority (minimum value).
// Returns ok=false if the queue is empty.
func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	return queue.heap.DeleteMin()
}

// Peek returns the element with the highest priority (minimum value) without removing it.
// Returns ok=false if the queue is empty.
func (queue *Queue[T]) Peek() (value T, ok bool) {
	return queue.heap.FindMin()
}

// IsEmpty returns true if the queue contains no elements.
func (queue *Queue[T]) IsEmpty() bool {
	return queue.heap.IsEmpty()
}

// Size returns the number of elements in the queue.
func (queue *Queue[T]) Size() int {
	return queue.heap.Size()
}

// Clear removes all elements from the queue.
func (queue *Queue[T]) Clear() {
	queue.heap.Clear()
}
