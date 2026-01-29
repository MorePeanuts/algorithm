// Package leetcode0002 solves LeetCode 2. Add Two Numbers
package leetcode0002

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type ListNode = utils.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	carry := 0
	var res, p0 *ListNode

	for p1 != nil || p2 != nil {
		sum := carry
		if p1 != nil {
			sum += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			sum += p2.Val
			p2 = p2.Next
		}

		node := ListNode{Val: sum % 10}
		carry = sum / 10

		if res == nil {
			res = &node
			p0 = res
		} else {
			p0.Next = &node
			p0 = p0.Next
		}
	}

	if carry > 0 {
		p0.Next = &ListNode{Val: carry}
	}

	return res
}
