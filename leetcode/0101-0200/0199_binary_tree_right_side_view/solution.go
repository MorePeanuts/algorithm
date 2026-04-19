// Package leetcode0199 solves LeetCode 199. Binary Tree Right Side View
package leetcode0199

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type TreeNode = utils.TreeNode

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		nums := len(queue)
		for i := range nums {
			node := queue[0]
			if i == nums-1 {
				res = append(res, node.Val)
			}
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
