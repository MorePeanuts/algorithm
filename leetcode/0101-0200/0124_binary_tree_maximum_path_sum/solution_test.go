// Package leetcode0124 solves LeetCode 124. Binary Tree Maximum Path Sum
package leetcode0124

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestMaxPathSumExamples(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want int
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)},
			6,
		},
		{
			"example_2",
			[]*int{utils.IntPtr(-10), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)},
			42,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := maxPathSum(root)
			if got != tst.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tst.want)
			}
		})
	}
}

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want int
	}{
		{"single_node", []*int{utils.IntPtr(1)}, 1},
		{"single_negative_node", []*int{utils.IntPtr(-100)}, -100},
		{"single_node_zero", []*int{utils.IntPtr(0)}, 0},
		{"two_nodes_left", []*int{utils.IntPtr(1), utils.IntPtr(2), nil}, 3},
		{"two_nodes_right", []*int{utils.IntPtr(1), nil, utils.IntPtr(3)}, 4},
		{"all_negative_small", []*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3)}, -1},
		{"all_negative_large", []*int{utils.IntPtr(-10), utils.IntPtr(-20), utils.IntPtr(-5), utils.IntPtr(-3), nil, utils.IntPtr(-1), nil}, -1},
		{"left_chain", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, 10},
		{"right_chain", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, 10},
		{"root_with_negative_children", []*int{utils.IntPtr(10), utils.IntPtr(-1), utils.IntPtr(-2)}, 10},
		{"root_with_one_positive_child", []*int{utils.IntPtr(5), utils.IntPtr(-10), utils.IntPtr(10)}, 15},
		{"path_through_root", []*int{utils.IntPtr(1), utils.IntPtr(-2), utils.IntPtr(3)}, 4},
		{"max_path_in_left_subtree", []*int{utils.IntPtr(-10), utils.IntPtr(5), utils.IntPtr(-20), utils.IntPtr(3), utils.IntPtr(4)}, 12},
		{"max_path_in_right_subtree", []*int{utils.IntPtr(-10), utils.IntPtr(-20), utils.IntPtr(5), utils.IntPtr(3), utils.IntPtr(4)}, 5},
		{"zigzag_path", []*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}, 47},
		{"deep_tree", []*int{
			utils.IntPtr(1),
			utils.IntPtr(2), utils.IntPtr(3),
			utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7),
		}, 18},
		{"mixed_values_1", []*int{utils.IntPtr(2), utils.IntPtr(-1), utils.IntPtr(-2)}, 2},
		{"mixed_values_2", []*int{utils.IntPtr(5), utils.IntPtr(4), utils.IntPtr(8), utils.IntPtr(11), nil, utils.IntPtr(13), utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(2), nil, nil, nil, utils.IntPtr(1)}, 48},
		{"node_val_min", []*int{utils.IntPtr(-1000)}, -1000},
		{"node_val_max", []*int{utils.IntPtr(1000)}, 1000},
		{"two_node_negative", []*int{utils.IntPtr(-1), utils.IntPtr(-2), nil}, -1},
		{"three_node_middle_positive", []*int{utils.IntPtr(-1), utils.IntPtr(3), utils.IntPtr(-2)}, 3},
		{"cross_path", []*int{utils.IntPtr(10), utils.IntPtr(2), utils.IntPtr(10), utils.IntPtr(20), utils.IntPtr(1), nil, utils.IntPtr(-25), nil, nil, nil, nil, utils.IntPtr(3), utils.IntPtr(4)}, 42},
		{"left_only_tree", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(-4)}, 6},
		{"right_only_tree", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(-4)}, 6},
		{"negative_and_zero", []*int{utils.IntPtr(0), utils.IntPtr(-1), utils.IntPtr(-2)}, 0},
		{"zero_in_middle", []*int{utils.IntPtr(-1), utils.IntPtr(0), utils.IntPtr(-2)}, 0},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := maxPathSum(root)
			if got != tst.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tst.want)
			}
		})
	}
}
