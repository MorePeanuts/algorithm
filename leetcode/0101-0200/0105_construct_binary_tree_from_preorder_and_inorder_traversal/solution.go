// Package leetcode0105 solves LeetCode 105. Construct Binary Tree from Preorder and Inorder Traversal
package leetcode0105

import (
	"slices"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

type TreeNode = utils.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	rootIdx := slices.Index(inorder, preorder[0])
	leftTree := buildTree(preorder[1:rootIdx+1], inorder[:rootIdx])
	var rightTree *TreeNode
	if rootIdx != len(inorder)-1 {
		rightTree = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])
	}
	root.Left, root.Right = leftTree, rightTree
	return root
}
