package leetcode0235

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor2(root.Right, p, q)
	} else if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor2(root.Left, p, q)
	}
	return root
}
