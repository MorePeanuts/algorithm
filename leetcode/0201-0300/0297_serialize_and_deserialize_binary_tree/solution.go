// Package leetcode0297 solves LeetCode 297. Serialize and Deserialize Binary Tree
package leetcode0297

import (
	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

type TreeNode = utils.TreeNode

type Codec struct {
	levelOrder []*int
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.levelOrder = nil
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			val := node.Val
			this.levelOrder = append(this.levelOrder, &val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			var val *int
			this.levelOrder = append(this.levelOrder, val)
		}
	}

	runes := make([]rune, 0)
	for _, p := range this.levelOrder {
		if p == nil {
			runes = append(runes, 5000)
		} else {
			runes = append(runes, rune(*p)+1000)
		}
	}
	return string(runes)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.levelOrder = nil
	for _, r := range data {
		if r == 5000 {
			var val *int
			this.levelOrder = append(this.levelOrder, val)
		} else {
			val := int(r) - 1000
			this.levelOrder = append(this.levelOrder, &val)
		}
	}

	if this.levelOrder[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *this.levelOrder[0]}
	queue := []*TreeNode{root}
	for i := 1; len(queue) > 0; {
		node := queue[0]
		queue = queue[1:]
		if this.levelOrder[i] != nil {
			node.Left = &TreeNode{Val: *this.levelOrder[i]}
			queue = append(queue, node.Left)
		}
		i++
		if this.levelOrder[i] != nil {
			node.Right = &TreeNode{Val: *this.levelOrder[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
