package leetcode0143

func reorderList3(head *ListNode) {
	table := make([]*ListNode, 0)
	for ptr := head; ptr != nil; ptr = ptr.Next {
		table = append(table, ptr)
	}
	for l, r := 0, len(table)-1; l < r; l, r = l+1, r-1 {
		table[l].Next = table[r]
		if l+1 < r {
			table[r].Next = table[l+1]
			if l+1 == r-1 {
				table[l+1].Next = nil
			}
		} else {
			table[r].Next = nil
		}
	}
}
