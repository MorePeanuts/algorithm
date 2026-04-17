// Package arrayqueue provides a generic queue implementation using a dynamic array.
package arrayqueue

import "github.com/MorePeanuts/algorithm/gods/lists/arraylist"

// Queue is a generic FIFO (First In, First Out) queue that stores elements of type T.
// It implements the queues.Queue interface and containers.Container interface.
type Queue[T comparable] struct {
	list *arraylist.List[T]
}

// New creates a new empty queue.
func New[T comparable]() *Queue[T] {
	return &Queue[T]{arraylist.New[T]()}
}

// queues.Queue interface implementation

// Enqueue adds an element to the end of the queue.
func (queue *Queue[T]) Enqueue(value T) {
	queue.list.Append(value)
}

// Dequeue removes and returns the element from the front of the queue.
// It returns the element and true if the queue is not empty,
// otherwise it returns the zero value of T and false.
// Note that this implementation requires linear time.
func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	value, ok = queue.list.Get(0)
	if ok {
		queue.list.Remove(0)
	}
	return
}

// Peek returns the element from the front of the queue without removing it.
// It returns the element and true if the queue is not empty,
// otherwise it returns the zero value of T and false.
func (queue *Queue[T]) Peek() (value T, ok bool) {
	return queue.list.Get(0)
}

// containers.Container interface implementation

// IsEmpty returns true if the queue contains no elements.
func (queue *Queue[T]) IsEmpty() bool {
	return queue.list.IsEmpty()
}

// Size returns the number of elements in the queue.
func (queue *Queue[T]) Size() int {
	return queue.list.Size()
}

// Clear removes all elements from the queue.
func (queue *Queue[T]) Clear() {
	queue.list.Clear()
}
