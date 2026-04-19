// Package leetcode0235 solves LeetCode 235. Lowest Common Ancestor of a Binary Search Tree
package leetcode0235

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath := getPath(root, p)
	qPath := getPath(root, q)
	for i, pp := range pPath {
		if i+1 == len(pPath) || i+1 == len(qPath) {
			return pp
		}
		if pPath[i+1] != qPath[i+1] {
			return pp
		}
	}
	return nil
}

func getPath(root, n *TreeNode) []*TreeNode {
	path := make([]*TreeNode, 0)
	for ptr := root; ; {
		path = append(path, ptr)
		if n.Val < ptr.Val {
			ptr = ptr.Left
		} else if n.Val > ptr.Val {
			ptr = ptr.Right
		} else {
			break
		}
	}
	return path
}
