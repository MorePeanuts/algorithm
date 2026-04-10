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
