// Package linkedlist provides a generic doubly linked list implementation.
package linkedlist

// node represents a single element in the linked list.
type node[T comparable] struct {
	value T
	next  *node[T]
}

// List is a generic singly linked list that stores elements of type T.
// It implements the list.List interface and containers.Container interface.
type List[T comparable] struct {
	first *node[T]
	last  *node[T]
	len   int
}

// New creates a new linked list with the given values (optional).
// If no values are provided, an empty list is returned.
func New[T comparable](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Append(values...)
	}
	return list
}

// list.List interface implementation

// Get returns the element at the specified index.
// It returns the element and true if the index is valid,
// otherwise it returns the zero value of T and false.
func (list *List[T]) Get(idx int) (T, bool) {
	if !list.withinRange(idx) {
		var zero T
		return zero, false
	}
	ptr := list.first
	for range idx {
		ptr = ptr.next
	}
	return ptr.value, true
}

// Set updates the element at the specified index.
// If the index is equal to the list length, the value is appended.
// If the index is out of range, no action is taken.
func (list *List[T]) Set(idx int, value T) {
	if !list.withinRange(idx) {
		if idx == list.len {
			list.Append(value)
		}
		return
	}

	ptr := list.first
	for range idx {
		ptr = ptr.next
	}
	ptr.value = value
}

// Append adds one or more values to the end of the list.
func (list *List[T]) Append(values ...T) {
	for _, value := range values {
		newNode := &node[T]{value, nil}
		if list.last == nil {
			list.last = newNode
			list.first = newNode
		} else {
			list.last.next = newNode
			list.last = newNode
		}
		list.len++
	}
}

// Remove removes the element at the specified index.
// If the index is out of range, no action is taken.
func (list *List[T]) Remove(idx int) {
	if !list.withinRange(idx) {
		return
	}

	if list.len == 1 {
		list.Clear()
		return
	}

	beforePtr := new(node[T])
	ptr := list.first
	for range idx {
		beforePtr, ptr = ptr, ptr.next
	}
	switch ptr {
	case list.first:
		list.first = ptr.next
	case list.last:
		beforePtr.next = nil
		list.last = beforePtr
	default:
		beforePtr.next = ptr.next
	}
	list.len--
}

// Insert inserts one or more values at the specified index.
// If the index is equal to the list length, the values are appended.
// If the index is out of range or no values are provided, no action is taken.
func (list *List[T]) Insert(idx int, values ...T) {
	if !list.withinRange(idx) {
		if idx == list.len {
			list.Append(values...)
		}
		return
	}
	if len(values) == 0 {
		return
	}

	beforePtr := new(node[T])
	ptr := list.first
	for range idx {
		beforePtr, ptr = ptr, ptr.next
	}
	newList := New(values...)
	newList.last.next = ptr
	if idx == 0 {
		list.first = newList.first
	} else {
		beforePtr.next = newList.first
	}

	list.len += len(values)
}

// Contains checks if all the specified values exist in the list.
// Returns true if all values are found, false otherwise.
// If no values are provided, returns true.
func (list *List[T]) Contains(values ...T) bool {
	set := make(map[T]bool)
	for ptr := list.first; ptr != nil; ptr = ptr.next {
		set[ptr.value] = true
	}
	for _, value := range values {
		if _, exist := set[value]; !exist {
			return false
		}
	}
	return true
}

// Swap swaps the elements at the specified indices i and j.
// If either index is out of range or i == j, no action is taken.
func (list *List[T]) Swap(i, j int) {
	if !list.withinRange(i) || !list.withinRange(j) || i == j {
		return
	}

	var iPtr, jPtr *node[T]
	for k, ptr := 0, list.first; iPtr == nil || jPtr == nil; k, ptr = k+1, ptr.next {
		switch k {
		case i:
			iPtr = ptr
		case j:
			jPtr = ptr
		}
	}
	iPtr.value, jPtr.value = jPtr.value, iPtr.value
}

// containers.Container interface implementation

// IsEmpty returns true if the list contains no elements.
func (list *List[T]) IsEmpty() bool {
	return list.len == 0
}

// Len returns the number of elements in the list.
func (list *List[T]) Len() int {
	return list.len
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	list.first = nil
	list.last = nil
	list.len = 0
}

// withinRange checks if the given index is valid for the list.
// Returns true if idx is between 0 (inclusive) and list.len (exclusive).
func (list *List[T]) withinRange(idx int) bool {
	return idx >= 0 && idx < list.len
}