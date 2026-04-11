// Package leetcode0142 solves LeetCode 142. Linked List Cycle II
package leetcode0142

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type ListNode = utils.ListNode

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	// Calculate the length of the ring
	length := 1
	fast = slow.Next
	for fast != slow {
		length++
		fast = fast.Next
	}

	slow, fast = head, head
	for range length {
		fast = fast.Next
	}
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
