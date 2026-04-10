// Package leetcode0206 solves LeetCode 206. Reverse Linked List
package leetcode0206

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type ListNode = utils.ListNode

func reverseList(head *ListNode) *ListNode {
	var p1, p2, p3 *ListNode
	p2 = head
	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}
