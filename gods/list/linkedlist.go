// Package list provides linked list data structure implementations.
package list

// node represents a singly linked list node that holds an element and a pointer to the next node.
type node[T comparable] struct {
	element T
	next    *node[T]
}

// IsEmpty checks whether the linked list is empty.
// Returns true if the list contains no elements, false otherwise.
func (n *node[T]) IsEmpty() bool {
	return n.next == nil
}

// IsLast checks whether the given node is the last node in the list.
// Returns true if p is the last node, false otherwise.
func (n *node[T]) IsLast(p *node[T]) bool {
	return p.next == nil
}

// Find searches for the first occurrence of the element x in the list.
// Returns a pointer to the node containing x, or nil if x is not found.
func (n *node[T]) Find(x T) *node[T] {
	p := n.next
	for p != nil && p.element != x {
		p = p.next
	}
	return p
}

// FindPrevious searches for the node preceding the first occurrence of element x.
// Returns a pointer to the node before x, or the header node if x is not found.
func (n *node[T]) FindPrevious(x T) *node[T] {
	p := n
	for p.next != nil && p.next.element != x {
		p = p.next
	}
	return p
}

// Delete removes the first occurrence of element x from the list.
// If x is not found, the list remains unchanged.
func (n *node[T]) Delete(x T) {
	var p, temp *node[T]
	p = n.FindPrevious(x)

	if !n.IsLast(p) {
		temp = p.next
		p.next = temp.next
	}
}

// Insert inserts element x immediately after node p.
func (n *node[T]) Insert(x T, p *node[T]) {
	temp := &node[T]{x, p.next}
	p.next = temp
}

// MakeEmpty removes all elements from the list, leaving it empty.
func (n *node[T]) MakeEmpty() {
	n.next = nil
}

// LinkedList represents a singly linked list with a header node.
// It is an alias for node[T], where the node serves as the list header.
type LinkedList[T comparable] = node[T]
