package circularqueue

import (
	"testing"
)

func TestNew(t *testing.T) {
	// Test empty queue
	queue := New[int]()
	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if queue.Size() != 0 {
		t.Errorf("New queue length should be 0, got %d", queue.Size())
	}
}

func TestEnqueue(t *testing.T) {
	// Test enqueue to empty queue
	queue := New[int]()
	queue.Enqueue(1)
	if queue.Size() != 1 {
		t.Errorf("Enqueue to empty queue: length should be 1, got %d", queue.Size())
	}
	if val, ok := queue.Peek(); !ok || val != 1 {
		t.Errorf("Enqueue to empty queue: front element should be 1, got %v", val)
	}

	// Test enqueue multiple values
	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.Size() != 3 {
		t.Errorf("Enqueue multiple: length should be 3, got %d", queue.Size())
	}
	if val, ok := queue.Peek(); !ok || val != 1 {
		t.Errorf("Enqueue multiple: front element should be 1, got %v", val)
	}

	// Test enqueue many values (exceed initial capacity to trigger grow)
	largeQueue := New[int]()
	for i := 0; i < 100; i++ {
		largeQueue.Enqueue(i)
	}
	if largeQueue.Size() != 100 {
		t.Errorf("Enqueue 100 values: length should be 100, got %d", largeQueue.Size())
	}
	if val, ok := largeQueue.Peek(); !ok || val != 0 {
		t.Errorf("Enqueue 100 values: front element should be 0, got %v", val)
	}
}

func TestDequeue(t *testing.T) {
	// Test dequeue from queue with elements
	queue := New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val, ok := queue.Dequeue()
	if !ok || val != 1 {
		t.Errorf("Dequeue should return 1, got %v (ok: %v)", val, ok)
	}
	if queue.Size() != 2 {
		t.Errorf("After dequeue: length should be 2, got %d", queue.Size())
	}

	// Test dequeue second element
	val, ok = queue.Dequeue()
	if !ok || val != 2 {
		t.Errorf("Dequeue should return 2, got %v (ok: %v)", val, ok)
	}
	if queue.Size() != 1 {
		t.Errorf("After dequeue: length should be 1, got %d", queue.Size())
	}

	// Test dequeue last element
	val, ok = queue.Dequeue()
	if !ok || val != 3 {
		t.Errorf("Dequeue should return 3, got %v (ok: %v)", val, ok)
	}
	if !queue.IsEmpty() {
		t.Error("After dequeuing all elements, queue should be empty")
	}

	// Test dequeue from empty queue
	emptyQueue := New[int]()
	val, ok = emptyQueue.Dequeue()
	if ok {
		t.Errorf("Dequeue from empty queue should not return ok, got value %v", val)
	}

	// Test with string type
	strQueue := New[string]()
	strQueue.Enqueue("a")
	strQueue.Enqueue("b")
	strVal, ok := strQueue.Dequeue()
	if !ok || strVal != "a" {
		t.Errorf("Dequeue string queue should return 'a', got %v", strVal)
	}
}

func TestPeek(t *testing.T) {
	// Test peek on queue with elements
	queue := New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)

	val, ok := queue.Peek()
	if !ok || val != 1 {
		t.Errorf("Peek should return 1, got %v (ok: %v)", val, ok)
	}
	if queue.Size() != 2 {
		t.Errorf("Peek should not modify length, got %d", queue.Size())
	}

	// Test peek after enqueue
	queue.Enqueue(3)
	val, ok = queue.Peek()
	if !ok || val != 1 {
		t.Errorf("Peek should return 1, got %v (ok: %v)", val, ok)
	}

	// Test peek after dequeue
	queue.Dequeue()
	val, ok = queue.Peek()
	if !ok || val != 2 {
		t.Errorf("Peek should return 2, got %v (ok: %v)", val, ok)
	}

	// Test peek on empty queue
	emptyQueue := New[int]()
	val, ok = emptyQueue.Peek()
	if ok {
		t.Errorf("Peek on empty queue should not return ok, got value %v", val)
	}

	// Test peek on single-element queue
	singleQueue := New[int]()
	singleQueue.Enqueue(42)
	val, ok = singleQueue.Peek()
	if !ok || val != 42 {
		t.Errorf("Peek on single-element queue should return 42, got %v", val)
	}
}

func TestIsEmpty(t *testing.T) {
	// Test empty queue
	queue := New[int]()
	if !queue.IsEmpty() {
		t.Error("IsEmpty() on new queue should return true")
	}

	// Test non-empty queue
	queue.Enqueue(1)
	if queue.IsEmpty() {
		t.Error("IsEmpty() on queue with element should return false")
	}

	// Test queue cleared
	queue.Clear()
	if !queue.IsEmpty() {
		t.Error("IsEmpty() on cleared queue should return true")
	}

	// Test queue with multiple elements
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.IsEmpty() {
		t.Error("IsEmpty() on queue with multiple elements should return false")
	}

	// Test after dequeuing all elements
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	if !queue.IsEmpty() {
		t.Error("IsEmpty() after dequeuing all elements should return true")
	}
}

func TestSize(t *testing.T) {
	// Test empty queue
	queue := New[int]()
	if queue.Size() != 0 {
		t.Errorf("Len() on new queue should be 0, got %d", queue.Size())
	}

	// Test after enqueue
	queue.Enqueue(1)
	if queue.Size() != 1 {
		t.Errorf("Len() after enqueue should be 1, got %d", queue.Size())
	}

	// Test after multiple enqueues
	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.Size() != 3 {
		t.Errorf("Len() after multiple enqueues should be 3, got %d", queue.Size())
	}

	// Test after dequeue
	queue.Dequeue()
	if queue.Size() != 2 {
		t.Errorf("Len() after dequeue should be 2, got %d", queue.Size())
	}

	// Test after clear
	queue.Clear()
	if queue.Size() != 0 {
		t.Errorf("Len() after clear should be 0, got %d", queue.Size())
	}
}

func TestClear(t *testing.T) {
	// Test clear non-empty queue
	queue := New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Clear()
	if !queue.IsEmpty() {
		t.Error("Clear() should leave queue empty")
	}
	if queue.Size() != 0 {
		t.Errorf("Clear() should set length to 0, got %d", queue.Size())
	}

	// Test clear empty queue (should not panic)
	queue = New[int]()
	queue.Clear()
	if !queue.IsEmpty() {
		t.Error("Clear() on empty queue should leave it empty")
	}

	// Test clear and then use
	queue = New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Clear()
	queue.Enqueue(10)
	queue.Enqueue(20)
	if queue.Size() != 2 {
		t.Errorf("Clear() then enqueue: length should be 2, got %d", queue.Size())
	}
	if val, ok := queue.Peek(); !ok || val != 10 {
		t.Errorf("Clear() then enqueue: front element should be 10, got %v", val)
	}

	// Test clear single-element queue
	queue = New[int]()
	queue.Enqueue(42)
	queue.Clear()
	if !queue.IsEmpty() {
		t.Error("Clear() on single-element queue should leave it empty")
	}
}

// Test circular queue wrap-around behavior
func TestWrapAround(t *testing.T) {
	queue := New[int]()

	// Fill queue partially, dequeue, then enqueue to trigger wrap-around
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
	}
	for i := 0; i < 5; i++ {
		val, ok := queue.Dequeue()
		if !ok || val != i {
			t.Errorf("Dequeue %d should return %d, got %v", i, i, val)
		}
	}

	// Now enqueue more elements to trigger wrap-around
	for i := 10; i < 20; i++ {
		queue.Enqueue(i)
	}

	// Verify all elements are in order
	expected := 5
	for !queue.IsEmpty() {
		val, ok := queue.Dequeue()
		if !ok || val != expected {
			t.Errorf("Dequeue should return %d, got %v", expected, val)
		}
		expected++
	}
	if expected != 20 {
		t.Errorf("Expected to dequeue up to 19, stopped at %d", expected-1)
	}
}

// Test grow when queue is full and wrap-around has occurred
func TestGrowWithWrapAround(t *testing.T) {
	// Use smaller capacity for easier testing
	queue := New[int]()

	// Enqueue initial capacity elements
	for i := 0; i < initialCap; i++ {
		queue.Enqueue(i)
	}

	// Dequeue half to create wrap-around scenario
	for i := 0; i < initialCap/2; i++ {
		val, ok := queue.Dequeue()
		if !ok || val != i {
			t.Errorf("Dequeue %d should return %d, got %v", i, i, val)
		}
	}

	// Enqueue more to exceed capacity and trigger grow
	for i := initialCap; i < initialCap+10; i++ {
		queue.Enqueue(i)
	}

	// Verify all remaining elements are correct
	expected := initialCap / 2
	for !queue.IsEmpty() {
		val, ok := queue.Dequeue()
		if !ok || val != expected {
			t.Errorf("Dequeue should return %d, got %v", expected, val)
		}
		expected++
	}
	if expected != initialCap+10 {
		t.Errorf("Expected to dequeue up to %d, stopped at %d", initialCap+9, expected-1)
	}
}

// Test shrink behavior with wrap-around
func TestShrinkWithWrapAround(t *testing.T) {
	queue := New[int]()

	// Enqueue many elements to trigger multiple grows
	for i := 0; i < 100; i++ {
		queue.Enqueue(i)
	}

	// Dequeue most elements to trigger shrink
	for i := 0; i < 90; i++ {
		val, ok := queue.Dequeue()
		if !ok || val != i {
			t.Errorf("Dequeue %d should return %d, got %v", i, i, val)
		}
	}

	// Verify remaining elements
	expected := 90
	for !queue.IsEmpty() {
		val, ok := queue.Dequeue()
		if !ok || val != expected {
			t.Errorf("Dequeue should return %d, got %v", expected, val)
		}
		expected++
	}
	if expected != 100 {
		t.Errorf("Expected to dequeue up to 99, stopped at %d", expected-1)
	}
}

// Test comprehensive scenario with mixed operations
func TestQueueComprehensive(t *testing.T) {
	queue := New[int]()

	// Empty queue operations
	if !queue.IsEmpty() {
		t.Error("Queue should be empty initially")
	}

	// Enqueue and verify
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.Size() != 3 {
		t.Errorf("Length should be 3 after enqueues, got %d", queue.Size())
	}

	// Peek
	if val, ok := queue.Peek(); !ok || val != 1 {
		t.Errorf("Peek should return 1, got %v", val)
	}

	// Dequeue
	val, ok := queue.Dequeue()
	if !ok || val != 1 {
		t.Errorf("Dequeue should return 1, got %v", val)
	}
	if queue.Size() != 2 {
		t.Errorf("Length should be 2 after dequeue, got %d", queue.Size())
	}

	// Enqueue more
	queue.Enqueue(4)
	queue.Enqueue(5)
	if queue.Size() != 4 {
		t.Errorf("Length should be 4 after more enqueues, got %d", queue.Size())
	}

	// Dequeue all
	expected := []int{2, 3, 4, 5}
	for i, exp := range expected {
		val, ok := queue.Dequeue()
		if !ok || val != exp {
			t.Errorf("Dequeue %d should return %d, got %v", i, exp, val)
		}
	}

	// Verify empty
	if !queue.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}

	// Clear and reuse
	queue.Clear()
	queue.Enqueue(100)
	if val, ok := queue.Dequeue(); !ok || val != 100 {
		t.Errorf("After clear and enqueue: dequeue should return 100, got %v", val)
	}
}

// Test with different data types
func TestQueueWithDifferentTypes(t *testing.T) {
	// Test with string
	strQueue := New[string]()
	strQueue.Enqueue("hello")
	strQueue.Enqueue("world")
	if val, ok := strQueue.Dequeue(); !ok || val != "hello" {
		t.Errorf("String queue dequeue should return 'hello', got %v", val)
	}
	if val, ok := strQueue.Peek(); !ok || val != "world" {
		t.Errorf("String queue peek should return 'world', got %v", val)
	}

	// Test with float
	floatQueue := New[float64]()
	floatQueue.Enqueue(3.14)
	floatQueue.Enqueue(2.718)
	if val, ok := floatQueue.Dequeue(); !ok || val != 3.14 {
		t.Errorf("Float queue dequeue should return 3.14, got %v", val)
	}

	// Test with struct
	type Person struct {
		name string
		age  int
	}
	personQueue := New[Person]()
	personQueue.Enqueue(Person{"Alice", 30})
	personQueue.Enqueue(Person{"Bob", 25})
	if val, ok := personQueue.Dequeue(); !ok || val.name != "Alice" {
		t.Errorf("Person queue dequeue should return Alice, got %v", val)
	}
}