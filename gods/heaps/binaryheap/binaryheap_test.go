package binaryheap

import (
	"testing"
)

func TestNew(t *testing.T) {
	heap := New[int]()
	if !heap.IsEmpty() {
		t.Error("New heap should be empty")
	}
	if heap.Size() != 0 {
		t.Errorf("New heap size should be 0, got %d", heap.Size())
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
	maxHeap := NewWith[int](reverseInt)
	if !maxHeap.IsEmpty() {
		t.Error("New max heap should be empty")
	}
}

func TestInsertSingleValue(t *testing.T) {
	heap := New[int]()
	heap.Insert(5)
	if heap.Size() != 1 {
		t.Errorf("After insert single value, size should be 1, got %d", heap.Size())
	}
	if val, ok := heap.FindMin(); !ok || val != 5 {
		t.Errorf("FindMin should return 5, got %v (ok: %v)", val, ok)
	}
}

func TestInsertMultipleValues(t *testing.T) {
	heap := New[int]()
	heap.Insert(5, 3, 7, 1, 9)
	if heap.Size() != 5 {
		t.Errorf("After insert multiple values, size should be 5, got %d", heap.Size())
	}
	if val, ok := heap.FindMin(); !ok || val != 1 {
		t.Errorf("FindMin should return 1, got %v (ok: %v)", val, ok)
	}
}

func TestInsertIncrementally(t *testing.T) {
	heap := New[int]()
	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(7)
	heap.Insert(1)
	heap.Insert(9)
	if heap.Size() != 5 {
		t.Errorf("Size should be 5, got %d", heap.Size())
	}
	if val, ok := heap.FindMin(); !ok || val != 1 {
		t.Errorf("FindMin should return 1, got %v", val)
	}
}

func TestDeleteMin(t *testing.T) {
	heap := New[int]()
	heap.Insert(5, 3, 7, 1, 9, 2, 8, 4, 6)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, exp := range expected {
		val, ok := heap.DeleteMin()
		if !ok || val != exp {
			t.Errorf("DeleteMin %d should return %d, got %v (ok: %v)", i, exp, val, ok)
		}
	}
	if !heap.IsEmpty() {
		t.Error("Heap should be empty after deleting all elements")
	}
}

func TestDeleteMinFromEmpty(t *testing.T) {
	heap := New[int]()
	val, ok := heap.DeleteMin()
	if ok {
		t.Errorf("DeleteMin from empty should not return ok, got %v", val)
	}
}

func TestFindMin(t *testing.T) {
	heap := New[int]()
	heap.Insert(5, 3, 7)
	if val, ok := heap.FindMin(); !ok || val != 3 {
		t.Errorf("FindMin should return 3, got %v (ok: %v)", val, ok)
	}
	if heap.Size() != 3 {
		t.Error("FindMin should not modify heap size")
	}
}

func TestFindMinFromEmpty(t *testing.T) {
	heap := New[int]()
	val, ok := heap.FindMin()
	if ok {
		t.Errorf("FindMin from empty should not return ok, got %v", val)
	}
}

func TestIsEmpty(t *testing.T) {
	heap := New[int]()
	if !heap.IsEmpty() {
		t.Error("New heap should be empty")
	}
	heap.Insert(1)
	if heap.IsEmpty() {
		t.Error("Heap with element should not be empty")
	}
	heap.Clear()
	if !heap.IsEmpty() {
		t.Error("Cleared heap should be empty")
	}
}

func TestSize(t *testing.T) {
	heap := New[int]()
	if heap.Size() != 0 {
		t.Errorf("Empty heap size should be 0, got %d", heap.Size())
	}
	heap.Insert(1, 2, 3)
	if heap.Size() != 3 {
		t.Errorf("Heap size should be 3, got %d", heap.Size())
	}
	heap.DeleteMin()
	if heap.Size() != 2 {
		t.Errorf("After DeleteMin, size should be 2, got %d", heap.Size())
	}
	heap.Clear()
	if heap.Size() != 0 {
		t.Errorf("After Clear, size should be 0, got %d", heap.Size())
	}
}

func TestClear(t *testing.T) {
	heap := New[int]()
	heap.Insert(1, 2, 3, 4, 5)
	heap.Clear()
	if !heap.IsEmpty() {
		t.Error("Heap should be empty after Clear")
	}
	if heap.Size() != 0 {
		t.Errorf("Size should be 0 after Clear, got %d", heap.Size())
	}
	heap.Clear()
	if !heap.IsEmpty() {
		t.Error("Clear on empty heap should be safe")
	}
	heap.Insert(10)
	if heap.Size() != 1 {
		t.Error("Should be able to insert after Clear")
	}
}

func TestMaxHeap(t *testing.T) {
	reverseInt := func(a, b int) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	}
	maxHeap := NewWith[int](reverseInt)
	maxHeap.Insert(5, 3, 7, 1, 9)
	if val, ok := maxHeap.FindMin(); !ok || val != 9 {
		t.Errorf("Max heap FindMin should return 9, got %v", val)
	}

	expected := []int{9, 7, 5, 3, 1}
	for i, exp := range expected {
		val, ok := maxHeap.DeleteMin()
		if !ok || val != exp {
			t.Errorf("Max heap DeleteMin %d should return %d, got %v", i, exp, val)
		}
	}
}

func TestHeapWithString(t *testing.T) {
	heap := New[string]()
	heap.Insert("banana", "apple", "cherry", "date")
	if val, ok := heap.FindMin(); !ok || val != "apple" {
		t.Errorf("FindMin should return 'apple', got %v", val)
	}

	heap.DeleteMin()
	if val, ok := heap.FindMin(); !ok || val != "banana" {
		t.Errorf("FindMin should return 'banana' after delete, got %v", val)
	}
}

func TestHeapComprehensive(t *testing.T) {
	heap := New[int]()

	if !heap.IsEmpty() {
		t.Error("Heap should be empty initially")
	}

	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(5)
	if heap.Size() != 3 {
		t.Error("Size should be 3")
	}
	if val, ok := heap.FindMin(); !ok || val != 5 {
		t.Error("FindMin should return 5")
	}

	heap.DeleteMin()
	if val, ok := heap.FindMin(); !ok || val != 10 {
		t.Error("FindMin should return 10 after delete")
	}

	heap.Insert(3, 15, 7)
	if heap.Size() != 5 {
		t.Error("Size should be 5")
	}

	expected := []int{3, 7, 10, 15, 20}
	for _, exp := range expected {
		val, _ := heap.DeleteMin()
		if val != exp {
			t.Errorf("Expected %d, got %d", exp, val)
		}
	}

	if !heap.IsEmpty() {
		t.Error("Heap should be empty now")
	}

	heap.Clear()
	heap.Insert(100)
	if val, ok := heap.DeleteMin(); !ok || val != 100 {
		t.Error("Should work after Clear and reinsert")
	}
}

func TestFloydAlgorithm(t *testing.T) {
	heap := New[int]()
	values := make([]int, 100)
	for i := 0; i < 100; i++ {
		values[i] = 99 - i
	}
	heap.Insert(values...)

	if heap.Size() != 100 {
		t.Errorf("Size should be 100, got %d", heap.Size())
	}

	for i := 0; i < 100; i++ {
		val, ok := heap.DeleteMin()
		if !ok || val != i {
			t.Errorf("DeleteMin should return %d, got %d", i, val)
		}
	}
}
