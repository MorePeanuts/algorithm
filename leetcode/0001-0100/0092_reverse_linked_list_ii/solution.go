// Package leetcode0092 solves LeetCode 92. Reverse Linked List II
package leetcode0092

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type ListNode = utils.ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var p1, p2 *ListNode
	i := 1
	for p2 = head; i < left; i++ {
		p1 = p2
		p2 = p2.Next
	}
	leftNode := p2
	for range right - left + 1 {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	// Now p1 points to rightNode
	if leftNode.Next != nil {
		leftNode.Next.Next = p1
	} else {
		head = p1
	}
	leftNode.Next = p2

	return head
}
