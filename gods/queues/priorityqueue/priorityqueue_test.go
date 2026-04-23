package priorityqueue

import (
	"testing"
)

func TestNew(t *testing.T) {
	queue := New[int]()
	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if queue.Size() != 0 {
		t.Errorf("New queue size should be 0, got %d", queue.Size())
	}
}

func TestNewWith(t *testing.T) {
	reverseInt := func(a, b int) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	}
	maxQueue := NewWith[int](reverseInt)
	if !maxQueue.IsEmpty() {
		t.Error("New max queue should be empty")
	}
}

func TestEnqueueSingleValue(t *testing.T) {
	queue := New[int]()
	queue.Enqueue(5)
	if queue.Size() != 1 {
		t.Errorf("After enqueue single value, size should be 1, got %d", queue.Size())
	}
	if val, ok := queue.Peek(); !ok || val != 5 {
		t.Errorf("Peek should return 5, got %v (ok: %v)", val, ok)
	}
}

func TestEnqueueMultipleValues(t *testing.T) {
	queue := New[int]()
	queue.Enqueue(5)
	queue.Enqueue(3)
	queue.Enqueue(7)
	queue.Enqueue(1)
	queue.Enqueue(9)
	if queue.Size() != 5 {
		t.Errorf("After enqueue multiple values, size should be 5, got %d", queue.Size())
	}
	if val, ok := queue.Peek(); !ok || val != 1 {
		t.Errorf("Peek should return 1, got %v (ok: %v)", val, ok)
	}
}

func TestDequeue(t *testing.T) {
	queue := New[int]()
	queue.Enqueue(5)
	queue.Enqueue(3)
	queue.Enqueue(7)
	queue.Enqueue(1)
	queue.Enqueue(9)
	queue.Enqueue(2)
	queue.Enqueue(8)
	queue.Enqueue(4)
	queue.Enqueue(6)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, exp := range expected {
		val, ok := queue.Dequeue()
		if !ok || val != exp {
			t.Errorf("Dequeue %d should return %d, got %v (ok: %v)", i, exp, val, ok)
		}
	}
	if !queue.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}
}

func TestDequeueFromEmpty(t *testing.T) {
	queue := New[int]()
	val, ok := queue.Dequeue()
	if ok {
		t.Errorf("Dequeue from empty should not return ok, got %v", val)
	}
}

func TestPeek(t *testing.T) {
	queue := New[int]()
	queue.Enqueue(5)
	queue.Enqueue(3)
	queue.Enqueue(7)
	if val, ok := queue.Peek(); !ok || val != 3 {
		t.Errorf("Peek should return 3, got %v (ok: %v)", val, ok)
	}
	if queue.Size() != 3 {
		t.Error("Peek should not modify queue size")
	}
}

func TestPeekFromEmpty(t *testing.T) {
	queue := New[int]()
	val, ok := queue.Peek()
	if ok {
		t.Errorf("Peek from empty should not return ok, got %v", val)
	}
}

func TestIsEmpty(t *testing.T) {
	queue := New[int]()
	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}
	queue.Enqueue(1)
	if queue.IsEmpty() {
		t.Error("Queue with element should not be empty")
	}
	queue.Clear()
	if !queue.IsEmpty() {
		t.Error("Cleared queue should be empty")
	}
}

func TestSize(t *testing.T) {
	queue := New[int]()
	if queue.Size() != 0 {
		t.Errorf("Empty queue size should be 0, got %d", queue.Size())
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.Size() != 3 {
		t.Errorf("Queue size should be 3, got %d", queue.Size())
	}
	queue.Dequeue()
	if queue.Size() != 2 {
		t.Errorf("After Dequeue, size should be 2, got %d", queue.Size())
	}
	queue.Clear()
	if queue.Size() != 0 {
		t.Errorf("After Clear, size should be 0, got %d", queue.Size())
	}
}

func TestClear(t *testing.T) {
	queue := New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Clear()
	if !queue.IsEmpty() {
		t.Error("Queue should be empty after Clear")
	}
	if queue.Size() != 0 {
		t.Errorf("Size should be 0 after Clear, got %d", queue.Size())
	}
	queue.Clear()
	if !queue.IsEmpty() {
		t.Error("Clear on empty queue should be safe")
	}
	queue.Enqueue(10)
	if queue.Size() != 1 {
		t.Error("Should be able to enqueue after Clear")
	}
}

func TestMaxPriorityQueue(t *testing.T) {
	reverseInt := func(a, b int) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	}
	maxQueue := NewWith[int](reverseInt)
	maxQueue.Enqueue(5)
	maxQueue.Enqueue(3)
	maxQueue.Enqueue(7)
	maxQueue.Enqueue(1)
	maxQueue.Enqueue(9)
	if val, ok := maxQueue.Peek(); !ok || val != 9 {
		t.Errorf("Max queue Peek should return 9, got %v", val)
	}

	expected := []int{9, 7, 5, 3, 1}
	for i, exp := range expected {
		val, ok := maxQueue.Dequeue()
		if !ok || val != exp {
			t.Errorf("Max queue Dequeue %d should return %d, got %v", i, exp, val)
		}
	}
}

func TestPriorityQueueWithString(t *testing.T) {
	queue := New[string]()
	queue.Enqueue("banana")
	queue.Enqueue("apple")
	queue.Enqueue("cherry")
	queue.Enqueue("date")
	if val, ok := queue.Peek(); !ok || val != "apple" {
		t.Errorf("Peek should return 'apple', got %v", val)
	}

	queue.Dequeue()
	if val, ok := queue.Peek(); !ok || val != "banana" {
		t.Errorf("Peek should return 'banana' after dequeue, got %v", val)
	}
}

func TestPriorityQueueWithCustomComparator(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	compareByAge := func(a, b Person) int {
		if a.Age < b.Age {
			return -1
		}
		if a.Age > b.Age {
			return 1
		}
		return 0
	}

	queue := NewWith[Person](compareByAge)
	queue.Enqueue(Person{"Alice", 30})
	queue.Enqueue(Person{"Bob", 25})
	queue.Enqueue(Person{"Charlie", 35})

	if val, ok := queue.Peek(); !ok || val.Name != "Bob" {
		t.Errorf("Peek should return 'Bob' (youngest), got %v", val)
	}

	val, ok := queue.Dequeue()
	if !ok || val.Name != "Bob" {
		t.Errorf("Dequeue should return 'Bob', got %v", val)
	}

	if val, ok := queue.Peek(); !ok || val.Name != "Alice" {
		t.Errorf("Peek should return 'Alice' after dequeue, got %v", val)
	}
}

func TestPriorityQueueComprehensive(t *testing.T) {
	queue := New[int]()

	if !queue.IsEmpty() {
		t.Error("Queue should be empty initially")
	}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(5)
	if queue.Size() != 3 {
		t.Error("Size should be 3")
	}
	if val, ok := queue.Peek(); !ok || val != 5 {
		t.Error("Peek should return 5")
	}

	queue.Dequeue()
	if val, ok := queue.Peek(); !ok || val != 10 {
		t.Error("Peek should return 10 after dequeue")
	}

	queue.Enqueue(3)
	queue.Enqueue(15)
	queue.Enqueue(7)
	if queue.Size() != 5 {
		t.Error("Size should be 5")
	}

	expected := []int{3, 7, 10, 15, 20}
	for _, exp := range expected {
		val, _ := queue.Dequeue()
		if val != exp {
			t.Errorf("Expected %d, got %d", exp, val)
		}
	}

	if !queue.IsEmpty() {
		t.Error("Queue should be empty now")
	}

	queue.Clear()
	queue.Enqueue(100)
	if val, ok := queue.Dequeue(); !ok || val != 100 {
		t.Error("Should work after Clear and reinsert")
	}
}

func TestPriorityQueueLargeScale(t *testing.T) {
	queue := New[int]()
	const n = 100
	for i := 0; i < n; i++ {
		queue.Enqueue(n - 1 - i)
	}

	if queue.Size() != n {
		t.Errorf("Size should be %d, got %d", n, queue.Size())
	}

	for i := 0; i < n; i++ {
		val, ok := queue.Dequeue()
		if !ok || val != i {
			t.Errorf("Dequeue should return %d, got %d", i, val)
		}
	}
}

func TestPriorityQueueNewWith(t *testing.T) {
	type Item struct {
		Name     string
		Priority int
	}

	// Higher priority first
	comparator := func(a, b Item) int {
		if a.Priority > b.Priority {
			return -1
		}
		if a.Priority < b.Priority {
			return 1
		}
		return 0
	}

	queue := NewWith[Item](comparator)
	queue.Enqueue(Item{"low", 1})
	queue.Enqueue(Item{"high", 3})
	queue.Enqueue(Item{"medium", 2})

	if val, ok := queue.Peek(); !ok || val.Name != "high" {
		t.Errorf("Peek should return high priority item, got %v", val.Name)
	}

	expected := []string{"high", "medium", "low"}
	for _, exp := range expected {
		val, ok := queue.Dequeue()
		if !ok || val.Name != exp {
			t.Errorf("Expected %s, got %s", exp, val.Name)
		}
	}
}

func TestPriorityQueueWithComparator(t *testing.T) {
	queue := NewWith[string](func(x, y string) int {
		if len(x) < len(y) {
			return -1
		} else if len(x) > len(y) {
			return 1
		} else {
			return 0
		}
	})
	queue.Enqueue("zebra")
	queue.Enqueue("apple")
	queue.Enqueue("monkey")

	if val, ok := queue.Peek(); !ok || val != "zebra" {
		t.Errorf("Peek should return 'apple', got %v", val)
	}
}

