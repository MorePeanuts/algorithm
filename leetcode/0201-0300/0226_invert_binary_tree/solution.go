// Package leetcode0226 solves LeetCode 226. Invert Binary Tree
package leetcode0226

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
