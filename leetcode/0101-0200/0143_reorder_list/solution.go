// Package leetcode0143 solves LeetCode 143. Reorder List
package leetcode0143

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type ListNode = utils.ListNode

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	p2 := slow.Next
	slow.Next = nil

	var p1 *ListNode
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1, p2 = p2, p3
	}

	p2 = head
	for p1 != nil && p2 != nil {
		nextP1 := p1.Next
		nextP2 := p2.Next

		p2.Next = p1
		p2 = nextP2

		p1.Next = p2
		p1 = nextP1
	}
}
