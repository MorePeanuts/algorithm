package leetcode0226

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	root.Left = invertTree2(root.Left)
	root.Right = invertTree2(root.Right)
	return root
}
