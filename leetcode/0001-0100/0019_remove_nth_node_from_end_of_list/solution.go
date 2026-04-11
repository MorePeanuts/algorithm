// Package leetcode0019 solves LeetCode 19. Remove Nth Node From End of List
package leetcode0019

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type ListNode = utils.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for range n {
		fast = fast.Next
	}
	var prev *ListNode

	for fast != nil {
		fast = fast.Next
		prev = slow
		slow = slow.Next
	}

	if prev != nil {
		prev.Next = slow.Next
	} else {
		head = slow.Next
	}
	return head
}
