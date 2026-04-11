// Package utils: contains common data structures and util functions used in leetcode
package utils

// ListNode represents a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// SliceToList converts an int slice to a linked list.
func SliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

// ListToSlice converts a linked list to an int slice.
func ListToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// EqualLists compares two linked lists for equality.
func EqualLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// SliceToCyclicList converts an int slice to a linked list with a cycle.
// pos is the index (0-based) of the node that the tail connects to.
// If pos is -1, no cycle is created.
func SliceToCyclicList(nums []int, pos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	nodes := []*ListNode{head}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
		nodes = append(nodes, curr)
	}
	if pos >= 0 && pos < len(nodes) {
		curr.Next = nodes[pos]
	}
	return head
}

// Node represents A node containing a random pointer.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// NodeWithRandomInput represents a node value and its random index for input.
type NodeWithRandomInput struct {
	Val         int
	RandomIndex *int // nil means random pointer is nil
}

// SliceToNodeWithRandom converts a slice of NodeWithRandomInput to a linked list with random pointers.
func SliceToNodeWithRandom(input []NodeWithRandomInput) *Node {
	if len(input) == 0 {
		return nil
	}

	// Create all nodes first
	nodes := make([]*Node, len(input))
	for i := range input {
		nodes[i] = &Node{Val: input[i].Val}
	}

	// Set Next and Random pointers
	for i := range input {
		if i < len(input)-1 {
			nodes[i].Next = nodes[i+1]
		}
		if input[i].RandomIndex != nil {
			nodes[i].Random = nodes[*input[i].RandomIndex]
		}
	}

	return nodes[0]
}

// NodeWithRandomToSlice converts a linked list with random pointers to a slice of NodeWithRandomInput.
func NodeWithRandomToSlice(head *Node) []NodeWithRandomInput {
	if head == nil {
		return nil
	}

	// First pass: collect all nodes and build node to index map
	nodes := []*Node{}
	nodeToIndex := make(map[*Node]int)
	for curr := head; curr != nil; curr = curr.Next {
		nodeToIndex[curr] = len(nodes)
		nodes = append(nodes, curr)
	}

	// Second pass: build result slice
	result := make([]NodeWithRandomInput, len(nodes))
	for i, node := range nodes {
		result[i].Val = node.Val
		if node.Random != nil {
			if idx, ok := nodeToIndex[node.Random]; ok {
				result[i].RandomIndex = &idx
			}
		}
	}

	return result
}

// EqualNodeWithRandom compares two linked lists with random pointers for equality.
// It checks:
// 1. Same sequence of Val values
// 2. Same Next pointer structure
// 3. Same Random pointer relationships
// 4. All nodes in the copied list are new nodes (not pointing to original list)
func EqualNodeWithRandom(original, got *Node) bool {
	// Handle nil case
	if original == nil && got == nil {
		return true
	}
	if original == nil || got == nil {
		return false
	}

	// Build maps for both lists
	originalNodes := []*Node{}
	originalNodeToIndex := make(map[*Node]int)
	for curr := original; curr != nil; curr = curr.Next {
		originalNodeToIndex[curr] = len(originalNodes)
		originalNodes = append(originalNodes, curr)
	}

	gotNodes := []*Node{}
	gotNodeToIndex := make(map[*Node]int)
	for curr := got; curr != nil; curr = curr.Next {
		// Check that this node is not from the original list
		if _, isOriginal := originalNodeToIndex[curr]; isOriginal {
			return false
		}
		gotNodeToIndex[curr] = len(gotNodes)
		gotNodes = append(gotNodes, curr)
	}

	// Check length
	if len(originalNodes) != len(gotNodes) {
		return false
	}

	// Check Val and Random relationships
	for i := range originalNodes {
		// Check Val
		if originalNodes[i].Val != gotNodes[i].Val {
			return false
		}

		// Check Random
		if originalNodes[i].Random == nil {
			if gotNodes[i].Random != nil {
				return false
			}
		} else {
			if gotNodes[i].Random == nil {
				return false
			}
			originalRandomIdx := originalNodeToIndex[originalNodes[i].Random]
			gotRandomIdx := gotNodeToIndex[gotNodes[i].Random]
			if originalRandomIdx != gotRandomIdx {
				return false
			}
		}
	}

	return true
}
