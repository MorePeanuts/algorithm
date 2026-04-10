package leetcode0143

func reorderList2(head *ListNode) {
	// Split the linked list in half
	var prevSlow *ListNode
	slow, fast := head, head
	for fast != nil {
		prevSlow = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	prevSlow.Next = nil

	// Reverse the linked list pointed to by `slow`
	var p1, p2, p3 *ListNode
	p2 = slow
	for p2 != nil {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	// Cross-connect two linked lists
	p2, p3 = head.Next, head
	for p1 != nil || p2 != nil {
		p3.Next = p1
		p3 = p3.Next
		if p1 != nil {
			p1 = p1.Next
			p3.Next = p2
			p3 = p3.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
}
