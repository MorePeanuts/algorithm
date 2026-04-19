package leetcode0235

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	ptr := root
	for ptr != nil {
		if p.Val > ptr.Val {
			ptr = ptr.Right
		} else if q.Val < ptr.Val {
			ptr = ptr.Left
		} else {
			break
		}
	}
	return ptr
}
