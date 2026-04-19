// Package leetcode1448 solves LeetCode 1448. Count Good Nodes in Binary Tree
package leetcode1448

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, currMax int) int {
	if root == nil {
		return 0
	}
	if root.Val >= currMax {
		currMax = root.Val
		return dfs(root.Left, currMax) + dfs(root.Right, currMax) + 1
	}
	return dfs(root.Left, currMax) + dfs(root.Right, currMax)
}
