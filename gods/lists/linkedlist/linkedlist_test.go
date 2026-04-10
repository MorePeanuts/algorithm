package linkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {
	// Test empty list
	list := New[int]()
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}
	if list.Len() != 0 {
		t.Errorf("New list length should be 0, got %d", list.Len())
	}

	// Test list with initial values
	list = New(1, 2, 3)
	if list.IsEmpty() {
		t.Error("List with initial values should not be empty")
	}
	if list.Len() != 3 {
		t.Errorf("List length should be 3, got %d", list.Len())
	}
	if val, ok := list.Get(0); !ok || val != 1 {
		t.Errorf("First element should be 1, got %v", val)
	}
	if val, ok := list.Get(2); !ok || val != 3 {
		t.Errorf("Last element should be 3, got %v", val)
	}
}

func TestGet(t *testing.T) {
	list := New(10, 20, 30)

	// Test valid indices
	tests := []struct {
		idx      int
		expected int
		found    bool
	}{
		{0, 10, true},
		{1, 20, true},
		{2, 30, true},
	}

	for _, tt := range tests {
		val, ok := list.Get(tt.idx)
		if ok != tt.found {
			t.Errorf("Get(%d) ok = %v, want %v", tt.idx, ok, tt.found)
		}
		if ok && val != tt.expected {
			t.Errorf("Get(%d) = %v, want %v", tt.idx, val, tt.expected)
		}
	}

	// Test invalid indices
	invalidIndices := []int{-1, 3, 100}
	for _, idx := range invalidIndices {
		val, ok := list.Get(idx)
		if ok {
			t.Errorf("Get(%d) should return false, got value %v", idx, val)
		}
	}

	// Test empty list
	emptyList := New[int]()
	val, ok := emptyList.Get(0)
	if ok {
		t.Errorf("Get(0) on empty list should return false, got value %v", val)
	}
}

func TestSet(t *testing.T) {
	// Test Set on existing index
	list := New(1, 2, 3)
	list.Set(1, 200)
	if val, ok := list.Get(1); !ok || val != 200 {
		t.Errorf("Set(1, 200) failed, got %v", val)
	}

	// Test Set on first index
	list.Set(0, 100)
	if val, ok := list.Get(0); !ok || val != 100 {
		t.Errorf("Set(0, 100) failed, got %v", val)
	}

	// Test Set on last index
	list.Set(2, 300)
	if val, ok := list.Get(2); !ok || val != 300 {
		t.Errorf("Set(2, 300) failed, got %v", val)
	}

	// Test Set on index equal to length (should append)
	list.Set(3, 400)
	if list.Len() != 4 {
		t.Errorf("Set(3, 400) should append, length should be 4, got %d", list.Len())
	}
	if val, ok := list.Get(3); !ok || val != 400 {
		t.Errorf("Set(3, 400) failed, got %v", val)
	}

	// Test Set on out of range index (should not append)
	list.Set(10, 1000)
	if list.Len() != 4 {
		t.Errorf("Set(10, 1000) should not modify list, length should be 4, got %d", list.Len())
	}

	// Test Set on negative index (should not modify)
	list.Set(-1, -100)
	if list.Len() != 4 {
		t.Errorf("Set(-1, -100) should not modify list, length should be 4, got %d", list.Len())
	}

	// Test Set on empty list at index 0 (should append)
	emptyList := New[int]()
	emptyList.Set(0, 50)
	if emptyList.Len() != 1 {
		t.Errorf("Set(0, 50) on empty list should append, length should be 1, got %d", emptyList.Len())
	}
	if val, ok := emptyList.Get(0); !ok || val != 50 {
		t.Errorf("Set(0, 50) on empty list failed, got %v", val)
	}
}

func TestAppend(t *testing.T) {
	// Test append to empty list
	list := New[int]()
	list.Append(1)
	if list.Len() != 1 {
		t.Errorf("Append to empty list: length should be 1, got %d", list.Len())
	}
	if val, ok := list.Get(0); !ok || val != 1 {
		t.Errorf("Append to empty list: first element should be 1, got %v", val)
	}

	// Test append multiple values
	list.Append(2, 3, 4)
	if list.Len() != 4 {
		t.Errorf("Append multiple: length should be 4, got %d", list.Len())
	}
	expected := []int{1, 2, 3, 4}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Append multiple: index %d should be %d, got %v", i, exp, val)
		}
	}

	// Test append single value
	list.Append(5)
	if list.Len() != 5 {
		t.Errorf("Append single: length should be 5, got %d", list.Len())
	}
	if val, ok := list.Get(4); !ok || val != 5 {
		t.Errorf("Append single: last element should be 5, got %v", val)
	}

	// Test append no values
	list.Append()
	if list.Len() != 5 {
		t.Errorf("Append no values: length should remain 5, got %d", list.Len())
	}

	// Test append to list with one element
	singleList := New(100)
	singleList.Append(200)
	if singleList.Len() != 2 {
		t.Errorf("Append to single-element list: length should be 2, got %d", singleList.Len())
	}
	if val, ok := singleList.Get(1); !ok || val != 200 {
		t.Errorf("Append to single-element list: second element should be 200, got %v", val)
	}
}

func TestRemove(t *testing.T) {
	// Test remove from middle
	list := New(1, 2, 3, 4, 5)
	list.Remove(2)
	if list.Len() != 4 {
		t.Errorf("Remove middle: length should be 4, got %d", list.Len())
	}
	expected := []int{1, 2, 4, 5}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Remove middle: index %d should be %d, got %v", i, exp, val)
		}
	}

	// Test remove first element
	list = New(1, 2, 3)
	list.Remove(0)
	if list.Len() != 2 {
		t.Errorf("Remove first: length should be 2, got %d", list.Len())
	}
	if val, ok := list.Get(0); !ok || val != 2 {
		t.Errorf("Remove first: first element should be 2, got %v", val)
	}

	// Test remove last element
	list = New(1, 2, 3)
	list.Remove(2)
	if list.Len() != 2 {
		t.Errorf("Remove last: length should be 2, got %d", list.Len())
	}
	if val, ok := list.Get(1); !ok || val != 2 {
		t.Errorf("Remove last: last element should be 2, got %v", val)
	}

	// Test remove from single-element list
	list = New(42)
	list.Remove(0)
	if !list.IsEmpty() {
		t.Error("Remove from single-element list should leave it empty")
	}
	if list.Len() != 0 {
		t.Errorf("Remove from single-element list: length should be 0, got %d", list.Len())
	}

	// Test remove invalid indices
	list = New(1, 2, 3)
	list.Remove(-1)
	if list.Len() != 3 {
		t.Errorf("Remove negative index: length should remain 3, got %d", list.Len())
	}
	list.Remove(3)
	if list.Len() != 3 {
		t.Errorf("Remove out of range index: length should remain 3, got %d", list.Len())
	}

	// Test remove from empty list
	emptyList := New[int]()
	emptyList.Remove(0) // Should not panic
	if emptyList.Len() != 0 {
		t.Errorf("Remove from empty list: length should remain 0, got %d", emptyList.Len())
	}
}

func TestInsert(t *testing.T) {
	// Test insert at beginning
	list := New(3, 4, 5)
	list.Insert(0, 1, 2)
	if list.Len() != 5 {
		t.Errorf("Insert at beginning: length should be 5, got %d", list.Len())
	}
	expected := []int{1, 2, 3, 4, 5}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Insert at beginning: index %d should be %d, got %v", i, exp, val)
		}
	}

	// Test insert in middle
	list = New(1, 2, 5, 6)
	list.Insert(2, 3, 4)
	if list.Len() != 6 {
		t.Errorf("Insert in middle: length should be 6, got %d", list.Len())
	}
	expected = []int{1, 2, 3, 4, 5, 6}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Insert in middle: index %d should be %d, got %v", i, exp, val)
		}
	}

	// Test insert at end (using list.Len())
	list = New(1, 2, 3)
	list.Insert(3, 4, 5)
	if list.Len() != 5 {
		t.Errorf("Insert at end: length should be 5, got %d", list.Len())
	}
	expected = []int{1, 2, 3, 4, 5}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Insert at end: index %d should be %d, got %v", i, exp, val)
		}
	}

	// Test insert single value
	list = New(1, 3)
	list.Insert(1, 2)
	if list.Len() != 3 {
		t.Errorf("Insert single: length should be 3, got %d", list.Len())
	}
	if val, ok := list.Get(1); !ok || val != 2 {
		t.Errorf("Insert single: index 1 should be 2, got %v", val)
	}

	// Test insert no values
	list = New(1, 2, 3)
	list.Insert(1)
	if list.Len() != 3 {
		t.Errorf("Insert no values: length should remain 3, got %d", list.Len())
	}

	// Test insert at 0 in empty list
	emptyList := New[int]()
	emptyList.Insert(0, 1, 2, 3)
	if emptyList.Len() != 3 {
		t.Errorf("Insert in empty list: length should be 3, got %d", emptyList.Len())
	}

	// Test insert at invalid index (should not modify)
	list = New(1, 2, 3)
	list.Insert(5, 10)
	if list.Len() != 3 {
		t.Errorf("Insert invalid index: length should remain 3, got %d", list.Len())
	}
	list.Insert(-1, 10)
	if list.Len() != 3 {
		t.Errorf("Insert negative index: length should remain 3, got %d", list.Len())
	}
}

func TestContains(t *testing.T) {
	list := New(1, 2, 3, 4, 5)

	// Test single value exists
	if !list.Contains(3) {
		t.Error("Contains(3) should return true")
	}
	if !list.Contains(1) {
		t.Error("Contains(1) should return true")
	}
	if !list.Contains(5) {
		t.Error("Contains(5) should return true")
	}

	// Test single value not exists
	if list.Contains(0) {
		t.Error("Contains(0) should return false")
	}
	if list.Contains(100) {
		t.Error("Contains(100) should return false")
	}

	// Test multiple values all exist
	if !list.Contains(2, 4) {
		t.Error("Contains(2, 4) should return true")
	}
	if !list.Contains(1, 3, 5) {
		t.Error("Contains(1, 3, 5) should return true")
	}

	// Test multiple values with one not existing
	if list.Contains(2, 6) {
		t.Error("Contains(2, 6) should return false")
	}
	if list.Contains(0, 1) {
		t.Error("Contains(0, 1) should return false")
	}

	// Test no values (should return true)
	if !list.Contains() {
		t.Error("Contains() with no values should return true")
	}

	// Test empty list
	emptyList := New[int]()
	if emptyList.Contains(1) {
		t.Error("Empty list Contains(1) should return false")
	}
	if !emptyList.Contains() {
		t.Error("Empty list Contains() should return true")
	}

	// Test with string type
	strList := New("a", "b", "c")
	if !strList.Contains("b") {
		t.Error("strList.Contains('b') should return true")
	}
	if strList.Contains("d") {
		t.Error("strList.Contains('d') should return false")
	}
}

func TestSwap(t *testing.T) {
	// Test swap in middle
	list := New(1, 2, 3, 4, 5)
	list.Swap(1, 3)
	if list.Len() != 5 {
		t.Errorf("Swap: length should remain 5, got %d", list.Len())
	}
	expected := []int{1, 4, 3, 2, 5}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Swap: index %d should be %d, got %v", i, exp, val)
		}
	}

	// Test swap first and last
	list = New(1, 2, 3)
	list.Swap(0, 2)
	expected = []int{3, 2, 1}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Swap first and last: index %d should be %d, got %v", i, exp, val)
		}
	}

	// Test swap adjacent elements
	list = New(1, 2, 3, 4)
	list.Swap(1, 2)
	expected = []int{1, 3, 2, 4}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Swap adjacent: index %d should be %d, got %v", i, exp, val)
		}
	}

	// Test swap with self (should do nothing)
	list = New(1, 2, 3)
	list.Swap(1, 1)
	expected = []int{1, 2, 3}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Swap self: index %d should be %d, got %v", i, exp, val)
		}
	}

	// Test swap with invalid indices (should do nothing)
	list = New(1, 2, 3)
	list.Swap(-1, 1)
	expected = []int{1, 2, 3}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Swap negative index: index %d should be %d, got %v", i, exp, val)
		}
	}
	list.Swap(1, 3)
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Swap out of range: index %d should be %d, got %v", i, exp, val)
		}
	}

	// Test swap on list with two elements
	list = New(10, 20)
	list.Swap(0, 1)
	if val, ok := list.Get(0); !ok || val != 20 {
		t.Errorf("Swap two elements: index 0 should be 20, got %v", val)
	}
	if val, ok := list.Get(1); !ok || val != 10 {
		t.Errorf("Swap two elements: index 1 should be 10, got %v", val)
	}
}

func TestIsEmpty(t *testing.T) {
	// Test empty list
	list := New[int]()
	if !list.IsEmpty() {
		t.Error("IsEmpty() on new list should return true")
	}

	// Test non-empty list
	list.Append(1)
	if list.IsEmpty() {
		t.Error("IsEmpty() on list with element should return false")
	}

	// Test list cleared
	list.Clear()
	if !list.IsEmpty() {
		t.Error("IsEmpty() on cleared list should return true")
	}

	// Test list with multiple elements
	list = New(1, 2, 3)
	if list.IsEmpty() {
		t.Error("IsEmpty() on list with multiple elements should return false")
	}

	// Test after removing last element
	list.Remove(2)
	list.Remove(1)
	list.Remove(0)
	if !list.IsEmpty() {
		t.Error("IsEmpty() after removing all elements should return true")
	}
}

func TestLen(t *testing.T) {
	// Test empty list
	list := New[int]()
	if list.Len() != 0 {
		t.Errorf("Len() on new list should be 0, got %d", list.Len())
	}

	// Test after append
	list.Append(1)
	if list.Len() != 1 {
		t.Errorf("Len() after append should be 1, got %d", list.Len())
	}

	// Test after multiple appends
	list.Append(2, 3)
	if list.Len() != 3 {
		t.Errorf("Len() after multiple appends should be 3, got %d", list.Len())
	}

	// Test after insert
	list.Insert(1, 10)
	if list.Len() != 4 {
		t.Errorf("Len() after insert should be 4, got %d", list.Len())
	}

	// Test after remove
	list.Remove(2)
	if list.Len() != 3 {
		t.Errorf("Len() after remove should be 3, got %d", list.Len())
	}

	// Test after clear
	list.Clear()
	if list.Len() != 0 {
		t.Errorf("Len() after clear should be 0, got %d", list.Len())
	}

	// Test with initial values
	list = New(1, 2, 3, 4, 5)
	if list.Len() != 5 {
		t.Errorf("Len() with initial values should be 5, got %d", list.Len())
	}
}

func TestClear(t *testing.T) {
	// Test clear non-empty list
	list := New(1, 2, 3, 4, 5)
	list.Clear()
	if !list.IsEmpty() {
		t.Error("Clear() should leave list empty")
	}
	if list.Len() != 0 {
		t.Errorf("Clear() should set length to 0, got %d", list.Len())
	}

	// Test clear empty list (should not panic)
	list = New[int]()
	list.Clear()
	if !list.IsEmpty() {
		t.Error("Clear() on empty list should leave it empty")
	}

	// Test clear and then use
	list = New(1, 2, 3)
	list.Clear()
	list.Append(10, 20)
	if list.Len() != 2 {
		t.Errorf("Clear() then append: length should be 2, got %d", list.Len())
	}
	if val, ok := list.Get(0); !ok || val != 10 {
		t.Errorf("Clear() then append: first element should be 10, got %v", val)
	}

	// Test clear single-element list
	list = New(42)
	list.Clear()
	if !list.IsEmpty() {
		t.Error("Clear() on single-element list should leave it empty")
	}
}

// Test comprehensive scenario
func TestListComprehensive(t *testing.T) {
	list := New[int]()

	// Empty list operations
	if !list.IsEmpty() {
		t.Error("List should be empty initially")
	}

	// Append and verify
	list.Append(1, 2, 3)
	if list.Len() != 3 {
		t.Errorf("Length should be 3 after append, got %d", list.Len())
	}
	if !list.Contains(1, 2, 3) {
		t.Error("List should contain 1, 2, 3")
	}

	// Insert
	list.Insert(1, 10, 20)
	expected := []int{1, 10, 20, 2, 3}
	for i, exp := range expected {
		if val, ok := list.Get(i); !ok || val != exp {
			t.Errorf("Index %d should be %d, got %v", i, exp, val)
		}
	}

	// Set
	list.Set(0, 100)
	if val, ok := list.Get(0); !ok || val != 100 {
		t.Errorf("Index 0 should be 100 after set, got %v", val)
	}

	// Remove
	list.Remove(2)
	if list.Len() != 4 {
		t.Errorf("Length should be 4 after remove, got %d", list.Len())
	}
	if list.Contains(20) {
		t.Error("List should not contain 20 after removal")
	}

	// Swap
	list.Swap(0, 3)
	if val, ok := list.Get(0); !ok || val != 3 {
		t.Errorf("Index 0 should be 3 after swap, got %v", val)
	}
	if val, ok := list.Get(3); !ok || val != 100 {
		t.Errorf("Index 3 should be 100 after swap, got %v", val)
	}

	// Clear
	list.Clear()
	if !list.IsEmpty() {
		t.Error("List should be empty after clear")
	}
	if list.Len() != 0 {
		t.Errorf("Length should be 0 after clear, got %d", list.Len())
	}
}