package leetcode0138

func copyRandomList2(head *Node) *Node {
	old2new := make(map[*Node]*Node)

	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, exist := old2new[node]; exist {
			return n
		}
		n := &Node{Val: node.Val}
		old2new[node] = n
		n.Next = deepCopy(node.Next)
		n.Random = deepCopy(node.Random)
		return n
	}

	return deepCopy(head)
}
