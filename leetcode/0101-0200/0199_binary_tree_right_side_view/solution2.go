package leetcode0199

func rightSideView2(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, depth int, view *[]int) {
	if root == nil {
		return
	}
	if depth == len(*view) {
		*view = append(*view, root.Val)
	}
	dfs(root.Right, depth+1, view)
	dfs(root.Left, depth+1, view)
}
