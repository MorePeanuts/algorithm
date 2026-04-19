// Package leetcode0098 solves LeetCode 98. Validate Binary Search Tree
package leetcode0098

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	right := root.Right
	for right != nil && right.Left != nil {
		right = right.Left
	}
	if right != nil && right.Val <= root.Val {
		return false
	}
	left := root.Left
	for left != nil && left.Right != nil {
		left = left.Right
	}
	if left != nil && left.Val >= root.Val {
		return false
	}
	if !isValidBST(root.Left) || !isValidBST(root.Right) {
		return false
	}
	return true
}
