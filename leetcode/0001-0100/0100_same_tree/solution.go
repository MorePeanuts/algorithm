// Package leetcode0100 solves LeetCode 100. Same Tree
package leetcode0100

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
