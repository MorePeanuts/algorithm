// Package leetcode0543 solves LeetCode 543. Diameter of Binary Tree
package leetcode0543

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	maxDepth(root, &diameter)
	return diameter
}

func maxDepth(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}
	leftDepth := maxDepth(node.Left, diameter)
	rightDepth := maxDepth(node.Right, diameter)
	*diameter = max(*diameter, leftDepth+rightDepth)
	return max(leftDepth, rightDepth) + 1
}
