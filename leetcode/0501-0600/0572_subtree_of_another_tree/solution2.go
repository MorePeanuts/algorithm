package leetcode0572

func isSubtree2(root *TreeNode, subRoot *TreeNode) bool {
	subHeight := height(subRoot)
	_, found := findSubtree(root, subRoot, subHeight)
	return found
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := height(node.Left)
	rightHeight := height(node.Right)
	return max(leftHeight, rightHeight) + 1
}

func findSubtree(node *TreeNode, subRoot *TreeNode, subHeight int) (int, bool) {
	if node == nil {
		return 0, false
	}
	leftHeight, leftFound := findSubtree(node.Left, subRoot, subHeight)
	if leftFound {
		return -1, true
	}
	rightHeight, rightFound := findSubtree(node.Right, subRoot, subHeight)
	if rightFound {
		return -1, true
	}
	rootHeight := max(leftHeight, rightHeight) + 1
	if rootHeight == subHeight && isSameTree(node, subRoot) {
		return rootHeight, true
	} else {
		return rootHeight, false
	}
}
