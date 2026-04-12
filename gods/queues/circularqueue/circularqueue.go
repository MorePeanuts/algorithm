// Package circularqueue implements a circular queue based on a slice.
// The circular queue uses a fixed-size slice with front and rear pointers
// to efficiently manage elements in FIFO (First-In-First-Out) order.
// When the queue reaches capacity, it automatically grows; when the queue
// size drops below a threshold, it automatically shrinks.
package circularqueue

// Queue represents a circular queue with elements of type T.
// T must be comparable to support equality checks.
type Queue[T comparable] struct {
	values   []T // Underlying slice to store queue elements
	front    int // Index of the front element
	rear     int // Index of the position after the end of the queue
	length   int // Current number of elements in the queue
	capacity int // Maximum capacity before resize
}

const (
	growthFactor    = float32(2.00) // Factor by which the queue grows when full
	shrinkFactor    = float32(0.50) // Factor by which the queue shrinks when below threshold
	shrinkThreshold = float32(0.20) // Shrink when length < threshold * capacity
	initialCap      = 16             // Initial capacity of a new queue
)

// New creates a new empty circular queue with initial capacity.
// The initial capacity is set to 16 elements.
func New[T comparable]() *Queue[T] {
	return &Queue[T]{make([]T, initialCap), 0, 0, 0, initialCap}
}

// queues.Queue interface implementation

// Enqueue adds an element to the end of the queue.
// If the queue is full, it automatically grows the underlying slice.
func (queue *Queue[T]) Enqueue(value T) {
	queue.values[queue.rear] = value
	queue.rear = queue.succ(queue.rear)
	queue.length++
	queue.grow()
}

// Dequeue removes and returns the element from the front of the queue.
// Returns the zero value of T and false if the queue is empty.
// After dequeue, it may shrink the queue if the size is below threshold.
func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	if queue.IsEmpty() {
		var t T
		return t, false
	}
	value, ok = queue.values[queue.front], true
	queue.front = queue.succ(queue.front)
	queue.length--
	queue.shrink()
	return
}

// Peek returns the element at the front of the queue without removing it.
// Returns the zero value of T and false if the queue is empty.
func (queue *Queue[T]) Peek() (value T, ok bool) {
	if queue.IsEmpty() {
		var t T
		return t, false
	}
	value, ok = queue.values[queue.front], true
	return
}

// containers.Container interface implementation

// IsEmpty returns true if the queue contains no elements.
func (queue *Queue[T]) IsEmpty() bool {
	return queue.length == 0
}

// Len returns the number of elements in the queue.
func (queue *Queue[T]) Len() int {
	return queue.length
}

// Clear removes all elements from the queue and resets it to initial capacity.
func (queue *Queue[T]) Clear() {
	clear(queue.values)
	queue.values = queue.values[:initialCap]
	queue.front = 0
	queue.rear = 0
	queue.length = 0
	queue.capacity = initialCap
}

// Private methods implementation of circular queue

// succ returns the next position in the circular buffer.
// If pos is at the end of the buffer, it wraps around to 0.
func (queue *Queue[T]) succ(pos int) int {
	pos++
	if pos == queue.capacity {
		pos = 0
	}
	return pos
}

// grow checks if the queue is full and grows it if necessary.
// When growing, the capacity is doubled and elements are rearranged
// to be contiguous starting from index 0.
func (queue *Queue[T]) grow() {
	if queue.length == queue.capacity {
		newCapacity := int(growthFactor * float32(queue.capacity))
		for range newCapacity {
			var t T
			queue.values = append(queue.values, t)
		}
		for i := range queue.rear {
			queue.values[queue.length+i] = queue.values[i]
		}
		copy(queue.values, queue.values[queue.front:queue.front+queue.length])
		queue.front = 0
		queue.rear = queue.length
		queue.capacity = newCapacity
	}
}

// shrink checks if the queue is below the shrink threshold and shrinks it if necessary.
// When shrinking, elements are rearranged to be contiguous starting from index 0.
// The capacity is halved when the queue length falls below 20% of capacity.
func (queue *Queue[T]) shrink() {
	if queue.length < int(shrinkThreshold*float32(queue.capacity)) {
		if queue.front < queue.rear {
			copy(queue.values, queue.values[queue.front:queue.rear])
		} else {
			temp := make([]T, queue.rear)
			copy(temp, queue.values[:queue.rear])
			n := copy(queue.values, queue.values[queue.front:])
			copy(queue.values[n:], temp)
		}
		newCapacity := int(shrinkFactor * float32(queue.capacity))
		queue.values = queue.values[:newCapacity]
		queue.capacity = newCapacity
		queue.front = 0
		queue.rear = queue.length
	}
}