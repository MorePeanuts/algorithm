// Package leetcode0124 solves LeetCode 124. Binary Tree Maximum Path Sum
package leetcode0124

import (
	"math"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

type TreeNode = utils.TreeNode

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, currMax *int) int {
	if root == nil {
		return 0
	}
	leftMax := dfs(root.Left, currMax)
	leftMax = max(root.Val, root.Val+leftMax)
	rightMax := dfs(root.Right, currMax)
	rightMax = max(root.Val, root.Val+rightMax)
	*currMax = max(*currMax, leftMax+rightMax-root.Val)
	return max(leftMax, rightMax)
}
