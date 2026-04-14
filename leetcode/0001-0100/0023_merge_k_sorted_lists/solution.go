// Package leetcode0023 solves LeetCode 23. Merge k Sorted Lists
package leetcode0023

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type ListNode = utils.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		j := 0
		for i := 0; i < len(lists); i, j = i+2, j+1 {
			if i+1 < len(lists) {
				lists[j] = merge2Lists(lists[i], lists[i+1])
			} else {
				lists[j] = merge2Lists(lists[i], nil)
			}
		}
		lists = lists[:j]
	}
	return lists[0]
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res, p3, winner *ListNode
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			winner = p1
			p1 = p1.Next
		} else {
			winner = p2
			p2 = p2.Next
		}
		if res == nil {
			res = winner
			p3 = res
		} else {
			p3.Next = winner
			p3 = p3.Next
		}
	}
	if p1 != nil {
		p3.Next = p1
	}
	if p2 != nil {
		p3.Next = p2
	}
	return res
}
