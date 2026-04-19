// Package leetcode0230 solves LeetCode 230. Kth Smallest Element in a BST
package leetcode0230

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	res, _ := inorder(root, k, &count)
	return res
}

func inorder(root *TreeNode, k int, curr *int) (int, bool) {
	if root == nil {
		return 0, false
	}
	leftNum, leftOk := inorder(root.Left, k, curr)
	if leftOk {
		return leftNum, leftOk
	}
	*curr++
	if *curr == k {
		return root.Val, true
	}
	rightNum, rightOk := inorder(root.Right, k, curr)
	if rightOk {
		return rightNum, rightOk
	}
	return 0, false
}
