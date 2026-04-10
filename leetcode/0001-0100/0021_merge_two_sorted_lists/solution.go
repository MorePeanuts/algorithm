// Package leetcode0021 solves LeetCode 21. Merge Two Sorted Lists
package leetcode0021

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type ListNode = utils.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var res, winner, p *ListNode
	for p1, p2 := list1, list2; p1 != nil || p2 != nil; {
		if p1 == nil || p2 != nil && p2.Val <= p1.Val {
			winner = p2
			p2 = p2.Next
		} else {
			winner = p1
			p1 = p1.Next
		}
		if res == nil {
			res = winner
			p = res
		} else {
			p.Next = winner
			p = p.Next
		}
	}
	return res
}
