// Package leetcode0110 solves LeetCode 110. Balanced Binary Tree
package leetcode0110

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

func isBalanced(root *TreeNode) bool {
	balance := true
	height(root, &balance)
	return balance
}

func height(node *TreeNode, balance *bool) int {
	if node == nil {
		return 0
	}
	leftHeight := height(node.Left, balance)
	if !(*balance) {
		return -1
	}
	rightHeight := height(node.Right, balance)
	if !(*balance) {
		return -1
	}
	if leftHeight-rightHeight > 1 || leftHeight-rightHeight < -1 {
		*balance = false
	}
	return max(leftHeight, rightHeight) + 1
}
