// Package leetcode0025 solves LeetCode 25. Reverse Nodes in k-Group
package leetcode0025

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type ListNode = utils.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	var p1, p2, res *ListNode
	p2 = head
	for {
		newHead, newTail, next := reverseList(p2, k)
		if p1 != nil {
			p1.Next = newHead
		}
		newTail.Next = next
		p1, p2 = newTail, next
		if res == nil {
			res = newHead
		}
		if next == nil {
			break
		}
	}
	return res
}

func reverseList(head *ListNode, k int) (newHead, newTail, next *ListNode) {
	l := 1
	var p *ListNode
	for p = head; p.Next != nil; l++ {
		p = p.Next
	}
	if l < k {
		return head, p, nil
	}

	var p1, p2, p3 *ListNode
	p2 = head
	newTail = head
	for range k {
		p3 = p2.Next
		p2.Next = p1
		p1, p2 = p2, p3
	}
	newHead = p1
	next = p2
	return
}
