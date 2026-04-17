package arraystack

import (
	"testing"
)

func TestNew(t *testing.T) {
	// Test empty stack
	stack := New[int]()
	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}
	if stack.Size() != 0 {
		t.Errorf("New stack length should be 0, got %d", stack.Size())
	}
}

func TestPush(t *testing.T) {
	// Test push to empty stack
	stack := New[int]()
	stack.Push(1)
	if stack.Size() != 1 {
		t.Errorf("Push to empty stack: length should be 1, got %d", stack.Size())
	}
	if val, ok := stack.Peek(); !ok || val != 1 {
		t.Errorf("Push to empty stack: top element should be 1, got %v", val)
	}

	// Test push multiple values
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("Push multiple: length should be 3, got %d", stack.Size())
	}
	if val, ok := stack.Peek(); !ok || val != 3 {
		t.Errorf("Push multiple: top element should be 3, got %v", val)
	}

	// Test push many values
	largeStack := New[int]()
	for i := 0; i < 100; i++ {
		largeStack.Push(i)
	}
	if largeStack.Size() != 100 {
		t.Errorf("Push 100 values: length should be 100, got %d", largeStack.Size())
	}
	if val, ok := largeStack.Peek(); !ok || val != 99 {
		t.Errorf("Push 100 values: top element should be 99, got %v", val)
	}
}

func TestPop(t *testing.T) {
	// Test pop from stack with elements
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	val, ok := stack.Pop()
	if !ok || val != 3 {
		t.Errorf("Pop should return 3, got %v (ok: %v)", val, ok)
	}
	if stack.Size() != 2 {
		t.Errorf("After pop: length should be 2, got %d", stack.Size())
	}

	// Test pop second element
	val, ok = stack.Pop()
	if !ok || val != 2 {
		t.Errorf("Pop should return 2, got %v (ok: %v)", val, ok)
	}
	if stack.Size() != 1 {
		t.Errorf("After pop: length should be 1, got %d", stack.Size())
	}

	// Test pop last element
	val, ok = stack.Pop()
	if !ok || val != 1 {
		t.Errorf("Pop should return 1, got %v (ok: %v)", val, ok)
	}
	if !stack.IsEmpty() {
		t.Error("After popping all elements, stack should be empty")
	}

	// Test pop from empty stack
	emptyStack := New[int]()
	val, ok = emptyStack.Pop()
	if ok {
		t.Errorf("Pop from empty stack should not return ok, got value %v", val)
	}

	// Test with string type
	strStack := New[string]()
	strStack.Push("a")
	strStack.Push("b")
	strVal, ok := strStack.Pop()
	if !ok || strVal != "b" {
		t.Errorf("Pop string stack should return 'b', got %v", strVal)
	}
}

func TestPeek(t *testing.T) {
	// Test peek on stack with elements
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)

	val, ok := stack.Peek()
	if !ok || val != 2 {
		t.Errorf("Peek should return 2, got %v (ok: %v)", val, ok)
	}
	if stack.Size() != 2 {
		t.Errorf("Peek should not modify length, got %d", stack.Size())
	}

	// Test peek after push
	stack.Push(3)
	val, ok = stack.Peek()
	if !ok || val != 3 {
		t.Errorf("Peek should return 3, got %v (ok: %v)", val, ok)
	}

	// Test peek after pop
	stack.Pop()
	val, ok = stack.Peek()
	if !ok || val != 2 {
		t.Errorf("Peek should return 2, got %v (ok: %v)", val, ok)
	}

	// Test peek on empty stack
	emptyStack := New[int]()
	val, ok = emptyStack.Peek()
	if ok {
		t.Errorf("Peek on empty stack should not return ok, got value %v", val)
	}

	// Test peek on single-element stack
	singleStack := New[int]()
	singleStack.Push(42)
	val, ok = singleStack.Peek()
	if !ok || val != 42 {
		t.Errorf("Peek on single-element stack should return 42, got %v", val)
	}
}

func TestIsEmpty(t *testing.T) {
	// Test empty stack
	stack := New[int]()
	if !stack.IsEmpty() {
		t.Error("IsEmpty() on new stack should return true")
	}

	// Test non-empty stack
	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("IsEmpty() on stack with element should return false")
	}

	// Test stack cleared
	stack.Clear()
	if !stack.IsEmpty() {
		t.Error("IsEmpty() on cleared stack should return true")
	}

	// Test stack with multiple elements
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.IsEmpty() {
		t.Error("IsEmpty() on stack with multiple elements should return false")
	}

	// Test after popping all elements
	stack.Pop()
	stack.Pop()
	stack.Pop()
	if !stack.IsEmpty() {
		t.Error("IsEmpty() after popping all elements should return true")
	}
}

func TestSize(t *testing.T) {
	// Test empty stack
	stack := New[int]()
	if stack.Size() != 0 {
		t.Errorf("Len() on new stack should be 0, got %d", stack.Size())
	}

	// Test after push
	stack.Push(1)
	if stack.Size() != 1 {
		t.Errorf("Len() after push should be 1, got %d", stack.Size())
	}

	// Test after multiple pushes
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("Len() after multiple pushes should be 3, got %d", stack.Size())
	}

	// Test after pop
	stack.Pop()
	if stack.Size() != 2 {
		t.Errorf("Len() after pop should be 2, got %d", stack.Size())
	}

	// Test after clear
	stack.Clear()
	if stack.Size() != 0 {
		t.Errorf("Len() after clear should be 0, got %d", stack.Size())
	}
}

func TestClear(t *testing.T) {
	// Test clear non-empty stack
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Clear()
	if !stack.IsEmpty() {
		t.Error("Clear() should leave stack empty")
	}
	if stack.Size() != 0 {
		t.Errorf("Clear() should set length to 0, got %d", stack.Size())
	}

	// Test clear empty stack (should not panic)
	stack = New[int]()
	stack.Clear()
	if !stack.IsEmpty() {
		t.Error("Clear() on empty stack should leave it empty")
	}

	// Test clear and then use
	stack = New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Clear()
	stack.Push(10)
	stack.Push(20)
	if stack.Size() != 2 {
		t.Errorf("Clear() then push: length should be 2, got %d", stack.Size())
	}
	if val, ok := stack.Peek(); !ok || val != 20 {
		t.Errorf("Clear() then push: top element should be 20, got %v", val)
	}

	// Test clear single-element stack
	stack = New[int]()
	stack.Push(42)
	stack.Clear()
	if !stack.IsEmpty() {
		t.Error("Clear() on single-element stack should leave it empty")
	}
}

// Test comprehensive scenario
func TestStackComprehensive(t *testing.T) {
	stack := New[int]()

	// Empty stack operations
	if !stack.IsEmpty() {
		t.Error("Stack should be empty initially")
	}

	// Push and verify
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("Length should be 3 after pushes, got %d", stack.Size())
	}

	// Peek
	if val, ok := stack.Peek(); !ok || val != 3 {
		t.Errorf("Peek should return 3, got %v", val)
	}

	// Pop
	val, ok := stack.Pop()
	if !ok || val != 3 {
		t.Errorf("Pop should return 3, got %v", val)
	}
	if stack.Size() != 2 {
		t.Errorf("Length should be 2 after pop, got %d", stack.Size())
	}

	// Push more
	stack.Push(4)
	stack.Push(5)
	if stack.Size() != 4 {
		t.Errorf("Length should be 4 after more pushes, got %d", stack.Size())
	}

	// Pop all
	expected := []int{5, 4, 2, 1}
	for i, exp := range expected {
		val, ok := stack.Pop()
		if !ok || val != exp {
			t.Errorf("Pop %d should return %d, got %v", i, exp, val)
		}
	}

	// Verify empty
	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}

	// Clear and reuse
	stack.Clear()
	stack.Push(100)
	if val, ok := stack.Pop(); !ok || val != 100 {
		t.Errorf("After clear and push: pop should return 100, got %v", val)
	}
}

// Test with different data types
func TestStackWithDifferentTypes(t *testing.T) {
	// Test with string
	strStack := New[string]()
	strStack.Push("hello")
	strStack.Push("world")
	if val, ok := strStack.Pop(); !ok || val != "world" {
		t.Errorf("String stack pop should return 'world', got %v", val)
	}
	if val, ok := strStack.Peek(); !ok || val != "hello" {
		t.Errorf("String stack peek should return 'hello', got %v", val)
	}

	// Test with float
	floatStack := New[float64]()
	floatStack.Push(3.14)
	floatStack.Push(2.718)
	if val, ok := floatStack.Pop(); !ok || val != 2.718 {
		t.Errorf("Float stack pop should return 2.718, got %v", val)
	}

	// Test with struct
	type Person struct {
		name string
		age  int
	}
	personStack := New[Person]()
	personStack.Push(Person{"Alice", 30})
	personStack.Push(Person{"Bob", 25})
	if val, ok := personStack.Pop(); !ok || val.name != "Bob" {
		t.Errorf("Person stack pop should return Bob, got %v", val)
	}
}