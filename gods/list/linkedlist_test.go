package list

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Create a new linked list (header node)
	var list LinkedList[int]

	// Test initial state
	if !list.IsEmpty() {
		t.Error("New list should be empty")
	}

	// Test insertion - insert after header
	list.Insert(10, &list)
	if list.IsEmpty() {
		t.Error("List should not be empty after insertion")
	}

	// Test Find method
	found := list.Find(10)
	if found == nil {
		t.Error("Should find element 10")
	} else if found.element != 10 {
		t.Errorf("Expected element 10, got %v", found.element)
	}

	// Test IsLast method
	if !list.IsLast(found) {
		t.Error("Node with 10 should be last")
	}

	// Insert more elements
	list.Insert(20, &list)
	list.Insert(30, &list)

	// Current list order: header -> 30 -> 20 -> 10

	// Test FindPrevious
	prev := list.FindPrevious(20)
	if prev == nil || prev.element != 30 {
		t.Error("FindPrevious(20) should return node with 30")
	}

	// Test Delete
	list.Delete(20)
	if list.Find(20) != nil {
		t.Error("Element 20 should be deleted")
	}

	// Now list: header -> 30 -> 10
	prev = list.FindPrevious(10)
	if prev == nil || prev.element != 30 {
		t.Error("FindPrevious(10) should return node with 30 after deleting 20")
	}

	// Test deleting non-existent element (should not panic)
	list.Delete(999)

	// Test MakeEmpty
	list.MakeEmpty()
	if !list.IsEmpty() {
		t.Error("List should be empty after MakeEmpty")
	}
}

func TestLinkedListString(t *testing.T) {
	var list LinkedList[string]

	list.Insert("world", &list)
	list.Insert("hello", &list)

	if list.Find("hello") == nil {
		t.Error("Should find 'hello'")
	}
	if list.Find("world") == nil {
		t.Error("Should find 'world'")
	}

	list.Delete("hello")
	if list.Find("hello") != nil {
		t.Error("'hello' should be deleted")
	}
}

func TestNodeEdgeCases(t *testing.T) {
	var list LinkedList[int]

	// Test Find on empty list
	if list.Find(10) != nil {
		t.Error("Find on empty list should return nil")
	}

	// Test FindPrevious on empty list
	prev := list.FindPrevious(10)
	if prev != &list {
		t.Error("FindPrevious on empty list should return header")
	}

	// Test Delete on empty list (should not panic)
	list.Delete(10)

	// Test IsLast on header node
	if !list.IsLast(&list) {
		t.Error("Header node in empty list should be considered last")
	}
}

