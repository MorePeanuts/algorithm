// Package leetcode0138 solves LeetCode 138. Copy List with Random Pointer
package leetcode0138

import "github.com/MorePeanuts/algorithm/leetcode/utils"

type Node = utils.Node

func copyRandomList(head *Node) *Node {
	table := make(map[*Node][]int)
	for ptr, i := head, 0; ptr != nil; ptr, i = ptr.Next, i+1 {
		table[ptr] = []int{i}
	}
	for ptr := range table {
		if ptr.Random != nil {
			table[ptr] = append(table[ptr], table[ptr.Random][0])
		}
	}

	// copy the table
	newTable := make([]*Node, len(table))
	for ptr := range table {
		newTable[table[ptr][0]] = &Node{Val: ptr.Val}
	}
	for _, i := range table {
		if i[0] < len(table)-1 {
			newTable[i[0]].Next = newTable[i[0]+1]
		}
		if len(i) > 1 {
			newTable[i[0]].Random = newTable[i[1]]
		}
	}

	if len(newTable) == 0 {
		return nil
	} else {
		return newTable[0]
	}
}
