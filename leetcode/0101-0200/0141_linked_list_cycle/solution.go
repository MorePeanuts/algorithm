// Package leetcode0141 solves LeetCode 141. Linked List Cycle
package leetcode0141

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type ListNode = utils.ListNode

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}
