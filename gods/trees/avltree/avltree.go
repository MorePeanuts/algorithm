// Package avltree provides a generic AVL tree implementation.
// AVL tree is a self-balancing binary search tree where the difference
// between heights of left and right subtrees (balance factor) for any
// node is at most 1.
package avltree

import (
	"cmp"

	"github.com/MorePeanuts/algorithm/gods/utils"
)

// Node represents a single node in the AVL tree.
type Node[K comparable, V any] struct {
	Key        K
	Value      V
	Parent     *Node[K, V]
	LeftChild  *Node[K, V]
	RightChild *Node[K, V]
	height     int
}

// Tree implements an AVL tree, a self-balancing binary search tree.
// The difference between heights of left and right subtrees (balance factor)
// for any node is at most 1.
type Tree[K comparable, V any] struct {
	Root       *Node[K, V]
	Comparator utils.Comparator[K]
	size       int
}

// New creates a new AVL tree with ordered keys using the default comparator.
func New[K cmp.Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{Comparator: cmp.Compare[K]}
}

// NewWith creates a new AVL tree with a custom comparator for key comparison.
func NewWith[K comparable, V any](comparator utils.Comparator[K]) *Tree[K, V] {
	return &Tree[K, V]{Comparator: comparator}
}

// Put inserts a key-value pair into the AVL tree. If the key already exists,
// its value is updated. The tree maintains its AVL property after insertion.
func (tree *Tree[K, V]) Put(key K, value V) {
	tree.Root = tree.insert(key, value, tree.Root)
	tree.Root.Parent = nil // Ensure root's parent is always nil
}

// insert recursively inserts a key-value pair into the subtree rooted at node
// and returns the new root of the subtree after balancing.
func (tree *Tree[K, V]) insert(key K, value V, node *Node[K, V]) *Node[K, V] {
	if node == nil {
		tree.size++
		return &Node[K, V]{Key: key, Value: value, height: 0}
	}

	c := tree.Comparator(key, node.Key)
	if c == 0 {
		// Key already exists, update value
		node.Key, node.Value = key, value
		return node
	} else if c < 0 {
		// Insert into left subtree
		node.LeftChild = tree.insert(key, value, node.LeftChild)
		node.LeftChild.Parent = node

		// Check if left subtree is unbalanced
		if node.getBalance() == 2 {
			cLeft := tree.Comparator(key, node.LeftChild.Key)
			if cLeft < 0 {
				// Left-left case: single right rotation
				node = node.singleRotateWithLeft()
			} else {
				// Left-right case: double rotation (left then right)
				node = node.doubleRotateWithLeft()
			}
		}
	} else {
		// Insert into right subtree
		node.RightChild = tree.insert(key, value, node.RightChild)
		node.RightChild.Parent = node

		// Check if right subtree is unbalanced
		if node.getBalance() == -2 {
			cRight := tree.Comparator(key, node.RightChild.Key)
			if cRight > 0 {
				// Right-right case: single left rotation
				node = node.singleRotateWithRight()
			} else {
				// Right-left case: double rotation (right then left)
				node = node.doubleRotateWithRight()
			}
		}
	}

	// Update height of current node
	node.height = 1 + max(node.LeftChild.Height(), node.RightChild.Height())
	return node
}

// Get retrieves the value associated with the given key. Returns the value
// and true if found, zero value and false otherwise.
func (tree *Tree[K, V]) Get(key K) (value V, ok bool) {
	node := tree.GetNode(key)
	if node != nil {
		return node.Value, true
	}
	return value, false
}

// GetNode returns the node containing the given key, or nil if not found.
func (tree *Tree[K, V]) GetNode(key K) *Node[K, V] {
	node := tree.Root
	for node != nil {
		c := tree.Comparator(key, node.Key)
		if c == 0 {
			return node
		} else if c > 0 {
			// Key is greater, search right subtree
			node = node.RightChild
		} else {
			// Key is smaller, search left subtree
			node = node.LeftChild
		}
	}
	return node
}

// Remove removes a key from the AVL tree and rebalances the tree if necessary.
func (tree *Tree[K, V]) Remove(key K) {
	tree.Root = tree.remove(key, tree.Root)
	if tree.Root != nil {
		tree.Root.Parent = nil // Ensure root's parent is always nil
	}
}

// remove recursively removes a key from the subtree rooted at node
// and returns the new root of the subtree after balancing.
func (tree *Tree[K, V]) remove(key K, node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}

	c := tree.Comparator(key, node.Key)
	if c < 0 {
		// Key is smaller, remove from left subtree
		node.LeftChild = tree.remove(key, node.LeftChild)
		if node.LeftChild != nil {
			node.LeftChild.Parent = node
		}
	} else if c > 0 {
		// Key is larger, remove from right subtree
		node.RightChild = tree.remove(key, node.RightChild)
		if node.RightChild != nil {
			node.RightChild.Parent = node
		}
	} else {
		// Case 1: Leaf node or node with only one child
		if node.LeftChild == nil {
			tree.size--
			return node.RightChild
		} else if node.RightChild == nil {
			tree.size--
			return node.LeftChild
		}

		// Case 2: Node with two children - get inorder successor (smallest in right subtree)
		successor := node.RightChild
		for successor.LeftChild != nil {
			successor = successor.LeftChild
		}

		// Copy successor's key and value to this node
		node.Key = successor.Key
		node.Value = successor.Value

		// Remove the successor (this will decrement tree.size)
		node.RightChild = tree.remove(successor.Key, node.RightChild)
		if node.RightChild != nil {
			node.RightChild.Parent = node
		}
	}

	// Update height of current node
	node.height = 1 + max(node.LeftChild.Height(), node.RightChild.Height())

	// Get balance factor to check if this node is now unbalanced
	balance := node.getBalance()

	// Left Left Case
	if balance > 1 && node.LeftChild.getBalance() >= 0 {
		return node.singleRotateWithLeft()
	}

	// Left Right Case
	if balance > 1 && node.LeftChild.getBalance() < 0 {
		return node.doubleRotateWithLeft()
	}

	// Right Right Case
	if balance < -1 && node.RightChild.getBalance() <= 0 {
		return node.singleRotateWithRight()
	}

	// Right Left Case
	if balance < -1 && node.RightChild.getBalance() > 0 {
		return node.doubleRotateWithRight()
	}

	return node
}

// containers.Container interface implementation

// IsEmpty returns true if the tree contains no nodes.
func (tree *Tree[K, V]) IsEmpty() bool {
	return tree.size == 0
}

// Size returns the number of nodes in the tree.
func (tree *Tree[K, V]) Size() int {
	return tree.size
}

// Clear removes all nodes from the tree.
func (tree *Tree[K, V]) Clear() {
	tree.size = 0
	tree.Root = nil
}

// Node utility methods

// Size returns the number of nodes in the subtree rooted at this node.
func (n2 *Node[K, V]) Size() int {
	if n2 == nil {
		return 0
	}
	return 1 + n2.LeftChild.Size() + n2.RightChild.Size()
}

// Height returns the height of the subtree rooted at this node.
// Returns -1 for nil nodes.
func (n2 *Node[K, V]) Height() int {
	if n2 == nil {
		return -1
	}
	return n2.height
}

// getBalance calculates and returns the balance factor of this node.
// Balance factor is the difference between heights of left and right subtrees.
func (n2 *Node[K, V]) getBalance() int {
	if n2 == nil {
		return 0
	}
	return n2.LeftChild.Height() - n2.RightChild.Height()
}

// singleRotateWithLeft performs a single right rotation (LL case).
// n2 is the unbalanced node, returns the new root after rotation.
//
// Before:
//
//	    n2
//	   /  \
//	  n1   Z
//	 /  \
//	X    Y
//
// After:
//
//	  n1
//	 /  \
//	X    n2
//	    /  \
//	   Y    Z
func (n2 *Node[K, V]) singleRotateWithLeft() *Node[K, V] {
	n1 := n2.LeftChild

	// Reassign children
	if n1.RightChild != nil {
		n1.RightChild.Parent = n2
	}
	n2.LeftChild = n1.RightChild
	n2.Parent = n1
	n1.RightChild = n2

	// Update heights
	n1.height = max(n1.LeftChild.Height(), n1.RightChild.Height()) + 1
	n2.height = max(n2.LeftChild.Height(), n2.RightChild.Height()) + 1

	return n1
}

// singleRotateWithRight performs a single left rotation (RR case).
// n1 is the unbalanced node, returns the new root after rotation.
//
// Before:
//
//	  n1
//	 /  \
//	X    n2
//	    /  \
//	   Y    Z
//
// After:
//
//	    n2
//	   /  \
//	  n1   Z
//	 /  \
//	X    Y
func (n1 *Node[K, V]) singleRotateWithRight() *Node[K, V] {
	n2 := n1.RightChild

	// Reassign children
	if n2.LeftChild != nil {
		n2.LeftChild.Parent = n1
	}
	n1.RightChild = n2.LeftChild
	n1.Parent = n2
	n2.LeftChild = n1

	// Update heights
	n1.height = max(n1.LeftChild.Height(), n1.RightChild.Height()) + 1
	n2.height = max(n2.LeftChild.Height(), n2.RightChild.Height()) + 1

	return n2
}

// doubleRotateWithLeft performs a double rotation for LR case.
// First rotates left child right, then rotates node left.
func (n3 *Node[K, V]) doubleRotateWithLeft() *Node[K, V] {
	n3.LeftChild = n3.LeftChild.singleRotateWithRight()
	n3.LeftChild.Parent = n3
	return n3.singleRotateWithLeft()
}

// doubleRotateWithRight performs a double rotation for RL case.
// First rotates right child left, then rotates node right.
func (n1 *Node[K, V]) doubleRotateWithRight() *Node[K, V] {
	n1.RightChild = n1.RightChild.singleRotateWithLeft()
	n1.RightChild.Parent = n1
	return n1.singleRotateWithRight()
}
