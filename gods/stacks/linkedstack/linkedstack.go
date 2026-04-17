// Package linkedstack provides a generic stack implementation using a singly linked list.
package linkedstack

// node is a singly linked list node that holds a value and a pointer to the next node.
type node[T any] struct {
	value T
	next  *node[T]
}

// Stack is a generic LIFO (Last In, First Out) stack that stores elements of type T.
// It implements the stacks.Stack interface and containers.Container interface.
type Stack[T any] struct {
	top *node[T]
	len int
}

// New creates a new empty stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// stacks.Stack interface implementation

// Push adds an element to the top of the stack.
func (stack *Stack[T]) Push(value T) {
	newNode := &node[T]{value, stack.top}
	stack.top = newNode
	stack.len++
}

// Pop removes and returns the element from the top of the stack.
// It returns the element and true if the stack is not empty,
// otherwise it returns the zero value of T and false.
func (stack *Stack[T]) Pop() (value T, ok bool) {
	if stack.len > 0 {
		value, ok = stack.top.value, true
		stack.top = stack.top.next
		stack.len--
	}
	return
}

// Peek returns the element from the top of the stack without removing it.
// It returns the element and true if the stack is not empty,
// otherwise it returns the zero value of T and false.
func (stack *Stack[T]) Peek() (value T, ok bool) {
	if stack.len > 0 {
		value, ok = stack.top.value, true
	}
	return
}

// containers.Container interface implementation

// IsEmpty returns true if the stack contains no elements.
func (stack *Stack[T]) IsEmpty() bool {
	return stack.len == 0
}

// Size returns the number of elements in the stack.
func (stack *Stack[T]) Size() int {
	return stack.len
}

// Clear removes all elements from the stack.
func (stack *Stack[T]) Clear() {
	stack.len = 0
	stack.top = nil
}
