// Package arraystack provides a generic stack implementation using a dynamic array.
package arraystack

import "github.com/MorePeanuts/algorithm/gods/stacks"

// Interface guard
var _ stacks.Stack[any] = (*Stack[any])(nil)

// Stack is a generic LIFO (Last In, First Out) stack that stores elements of type T.
// It implements the containers.Container interface.
type Stack[T any] struct {
	elements []T
}

// New creates a new empty stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

// stacks.Stack interface implementation

// Push adds an element to the top of the stack.
func (stack *Stack[T]) Push(value T) {
	stack.elements = append(stack.elements, value)
}

// Pop removes and returns the element from the top of the stack.
// It returns the element and true if the stack is not empty,
// otherwise it returns the zero value of T and false.
func (stack *Stack[T]) Pop() (value T, ok bool) {
	if stack.Size() > 0 {
		value, ok = stack.elements[stack.Size()-1], true
		stack.elements = stack.elements[:stack.Size()-1]
	}
	return
}

// Peek returns the element from the top of the stack without removing it.
// It returns the element and true if the stack is not empty,
// otherwise it returns the zero value of T and false.
func (stack *Stack[T]) Peek() (value T, ok bool) {
	if stack.Size() > 0 {
		value, ok = stack.elements[stack.Size()-1], true
	}
	return
}

// containers.Container interface implementation

// IsEmpty returns true if the stack contains no elements.
func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.elements) == 0
}

// Size returns the number of elements in the stack.
func (stack *Stack[T]) Size() int {
	return len(stack.elements)
}

// Clear removes all elements from the stack.
func (stack *Stack[T]) Clear() {
	clear(stack.elements[:cap(stack.elements)])
	stack.elements = stack.elements[:0]
}
