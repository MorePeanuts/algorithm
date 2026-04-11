package leetcode0138

func copyRandomList3(head *Node) *Node {
	if head == nil {
		return nil
	}
	for ptr := head; ptr != nil; ptr = ptr.Next.Next {
		node := &Node{Val: ptr.Val, Next: ptr.Next, Random: nil}
		ptr.Next = node
	}
	for ptr := head; ptr != nil; ptr = ptr.Next.Next {
		if ptr.Random != nil {
			ptr.Next.Random = ptr.Random.Next
		}
	}
	res := head.Next
	for ptr, nextPtr := head, head.Next.Next; ptr != nil; ptr, nextPtr = nextPtr, nextPtr.Next.Next {
		if nextPtr != nil {
			ptr.Next.Next = nextPtr.Next
		}
		ptr.Next = nextPtr
		if nextPtr == nil {
			break
		}
	}
	return res
}
