// Package leetcode1448 solves LeetCode 1448. Count Good Nodes in Binary Tree
package leetcode1448

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestGoodNodesExamples(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want int
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(3), utils.IntPtr(1), utils.IntPtr(4), utils.IntPtr(3), nil, utils.IntPtr(1), utils.IntPtr(5)},
			4,
		},
		{
			"example_2",
			[]*int{utils.IntPtr(3), utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(2)},
			3,
		},
		{
			"example_3",
			[]*int{utils.IntPtr(1)},
			1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := goodNodes(root)
			if got != tst.want {
				t.Errorf("goodNodes() = %v, want %v", got, tst.want)
			}
		})
	}
}

func TestGoodNodes(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want int
	}{
		{
			"single_node",
			[]*int{utils.IntPtr(5)},
			1,
		},
		{
			"single_node_negative",
			[]*int{utils.IntPtr(-100)},
			1,
		},
		{
			"left_skewed_tree_increasing",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
			5,
		},
		{
			"left_skewed_tree_decreasing",
			[]*int{utils.IntPtr(5), utils.IntPtr(4), nil, utils.IntPtr(3), nil, utils.IntPtr(2), nil, utils.IntPtr(1)},
			1,
		},
		{
			"right_skewed_tree_increasing",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
			5,
		},
		{
			"right_skewed_tree_decreasing",
			[]*int{utils.IntPtr(5), nil, utils.IntPtr(4), nil, utils.IntPtr(3), nil, utils.IntPtr(2), nil, utils.IntPtr(1)},
			1,
		},
		{
			"all_same_values",
			[]*int{utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(2)},
			7,
		},
		{
			"perfect_binary_tree_increasing",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
			7,
		},
		{
			"perfect_binary_tree_decreasing",
			[]*int{utils.IntPtr(7), utils.IntPtr(6), utils.IntPtr(5), utils.IntPtr(4), utils.IntPtr(3), utils.IntPtr(2), utils.IntPtr(1)},
			1,
		},
		{
			"peak_in_middle",
			[]*int{utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(1), utils.IntPtr(6), utils.IntPtr(4), utils.IntPtr(2), utils.IntPtr(7)},
			4,
		},
		{
			"valley_in_middle",
			[]*int{utils.IntPtr(5), utils.IntPtr(3), utils.IntPtr(7), utils.IntPtr(2), utils.IntPtr(4), utils.IntPtr(6), utils.IntPtr(8)},
			3,
		},
		{
			"alternating_values",
			[]*int{utils.IntPtr(1), utils.IntPtr(10), utils.IntPtr(2), utils.IntPtr(9), utils.IntPtr(3), utils.IntPtr(8), utils.IntPtr(4)},
			5,
		},
		{
			"negative_values_all_good",
			[]*int{utils.IntPtr(-5), utils.IntPtr(-4), utils.IntPtr(-3), utils.IntPtr(-2), utils.IntPtr(-1)},
			5,
		},
		{
			"negative_values_none_good_except_root",
			[]*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3), utils.IntPtr(-4), utils.IntPtr(-5)},
			1,
		},
		{
			"mixed_negative_and_positive",
			[]*int{utils.IntPtr(-1), utils.IntPtr(5), utils.IntPtr(-2), utils.IntPtr(4), utils.IntPtr(3), utils.IntPtr(-3), utils.IntPtr(6)},
			3,
		},
		{
			"sparse_left_tree",
			[]*int{utils.IntPtr(10), utils.IntPtr(5), nil, utils.IntPtr(15), nil, nil, nil, utils.IntPtr(20)},
			2,
		},
		{
			"sparse_right_tree",
			[]*int{utils.IntPtr(10), nil, utils.IntPtr(5), nil, nil, nil, utils.IntPtr(15), nil, nil, nil, nil, nil, nil, nil, utils.IntPtr(20)},
			1,
		},
		{
			"zigzag_values",
			[]*int{utils.IntPtr(3), utils.IntPtr(1), utils.IntPtr(5), utils.IntPtr(2), utils.IntPtr(4), utils.IntPtr(0), utils.IntPtr(6)},
			4,
		},
		{
			"root_max_value",
			[]*int{utils.IntPtr(100), utils.IntPtr(50), utils.IntPtr(75), utils.IntPtr(25), utils.IntPtr(60), utils.IntPtr(70), utils.IntPtr(80)},
			1,
		},
		{
			"leaves_max_values",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(100), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(200)},
			7,
		},
		{
			"two_nodes_left",
			[]*int{utils.IntPtr(1), utils.IntPtr(2)},
			2,
		},
		{
			"two_nodes_right",
			[]*int{utils.IntPtr(2), nil, utils.IntPtr(1)},
			1,
		},
		{
			"three_nodes_all_increasing",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)},
			3,
		},
		{
			"three_nodes_all_decreasing",
			[]*int{utils.IntPtr(3), utils.IntPtr(2), utils.IntPtr(1)},
			1,
		},
		{
			"equal_path_max",
			[]*int{utils.IntPtr(5), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(5), nil, nil, utils.IntPtr(5)},
			4,
		},
		{
			"min_and_max_values",
			[]*int{utils.IntPtr(-10000), nil, utils.IntPtr(10000)},
			2,
		},
		{
			"deep_tree_with_good_nodes_scattered",
			[]*int{utils.IntPtr(5), utils.IntPtr(3), utils.IntPtr(8), utils.IntPtr(6), utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(9), utils.IntPtr(10), utils.IntPtr(5), nil, utils.IntPtr(6), nil, utils.IntPtr(8), nil, utils.IntPtr(11)},
			8,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := goodNodes(root)
			if got != tst.want {
				t.Errorf("goodNodes() = %v, want %v", got, tst.want)
			}
		})
	}
}
