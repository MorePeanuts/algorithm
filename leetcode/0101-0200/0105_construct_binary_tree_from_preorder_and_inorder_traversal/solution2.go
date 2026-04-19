package leetcode0105

func buildTree2(preorder []int, inorder []int) *TreeNode {
	idxMap := make(map[int]int)
	for i, num := range inorder {
		idxMap[num] = i
	}

	var dfs func(int, int, int) *TreeNode
	dfs = func(rootAtPreorder, leftAtInorder, rightAtInorder int) *TreeNode {
		if leftAtInorder > rightAtInorder {
			return nil
		}
		root := &TreeNode{Val: preorder[rootAtPreorder]}
		rootAtInorder := idxMap[preorder[rootAtPreorder]]
		root.Left = dfs(rootAtPreorder+1, leftAtInorder, rootAtInorder-1)
		root.Right = dfs(rootAtPreorder+1+rootAtInorder-leftAtInorder, rootAtInorder+1, rightAtInorder)
		return root
	}

	return dfs(0, 0, len(inorder)-1)
}
