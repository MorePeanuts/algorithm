// Package leetcode0098 solves LeetCode 98. Validate Binary Search Tree
package leetcode0098

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestIsValidBSTExamples(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want bool
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(3)},
			true,
		},
		{
			"example_2",
			[]*int{utils.IntPtr(5), utils.IntPtr(1), utils.IntPtr(4), nil, nil, utils.IntPtr(3), utils.IntPtr(6)},
			false,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			rootTree := utils.SliceToTreeNode(tst.root)
			got := isValidBST(rootTree)
			if got != tst.want {
				t.Errorf("isValidBST(%v) = %v, want %v", tst.root, got, tst.want)
			}
		})
	}
}

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want bool
	}{
		{"single_node", []*int{utils.IntPtr(1)}, true},
		{"single_node_min_val", []*int{utils.IntPtr(-2147483648)}, true},
		{"single_node_max_val", []*int{utils.IntPtr(2147483647)}, true},
		{"two_nodes_valid_left", []*int{utils.IntPtr(2), utils.IntPtr(1)}, true},
		{"two_nodes_valid_right", []*int{utils.IntPtr(1), nil, utils.IntPtr(2)}, true},
		{"two_nodes_invalid_left", []*int{utils.IntPtr(1), utils.IntPtr(2)}, false},
		{"two_nodes_invalid_right", []*int{utils.IntPtr(2), nil, utils.IntPtr(1)}, false},
		{"two_nodes_equal_left", []*int{utils.IntPtr(2), utils.IntPtr(2)}, false},
		{"two_nodes_equal_right", []*int{utils.IntPtr(2), nil, utils.IntPtr(2)}, false},
		{"three_nodes_valid_bst", []*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(3)}, true},
		{"three_nodes_invalid_right_child", []*int{utils.IntPtr(5), utils.IntPtr(1), utils.IntPtr(4)}, false},
		{"three_nodes_invalid_left_child", []*int{utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}, false},
		{"three_nodes_equal_root", []*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(2)}, false},
		{"deep_valid_bst", []*int{utils.IntPtr(8), utils.IntPtr(3), utils.IntPtr(10), utils.IntPtr(1), utils.IntPtr(6), nil, utils.IntPtr(14), nil, nil, utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(13)}, true},
		{"deep_invalid_subtree", []*int{utils.IntPtr(8), utils.IntPtr(3), utils.IntPtr(10), utils.IntPtr(1), utils.IntPtr(9), nil, utils.IntPtr(14)}, false},
		{"invalid_right_subtree_less_than_ancestor", []*int{utils.IntPtr(5), utils.IntPtr(3), utils.IntPtr(6), nil, nil, utils.IntPtr(4), utils.IntPtr(7)}, false},
		{"invalid_left_subtree_greater_than_ancestor", []*int{utils.IntPtr(5), utils.IntPtr(2), utils.IntPtr(7), utils.IntPtr(1), utils.IntPtr(6)}, false},
		{"left_skewed_valid", []*int{utils.IntPtr(5), utils.IntPtr(4), nil, utils.IntPtr(3), nil, utils.IntPtr(2), nil, utils.IntPtr(1)}, true},
		{"left_skewed_invalid", []*int{utils.IntPtr(5), utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(2)}, false},
		{"right_skewed_valid", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)}, true},
		{"right_skewed_invalid", []*int{utils.IntPtr(1), nil, utils.IntPtr(3), nil, utils.IntPtr(2), nil, utils.IntPtr(4)}, false},
		{"equal_values_in_deep_left", []*int{utils.IntPtr(3), utils.IntPtr(2), utils.IntPtr(4), utils.IntPtr(1), utils.IntPtr(2)}, false},
		{"equal_values_in_deep_right", []*int{utils.IntPtr(3), utils.IntPtr(2), utils.IntPtr(4), nil, nil, utils.IntPtr(3), utils.IntPtr(5)}, false},
		{"min_value_boundary_valid", []*int{utils.IntPtr(-2147483648), nil, utils.IntPtr(2147483647)}, true},
		{"min_value_boundary_invalid", []*int{utils.IntPtr(-2147483648), utils.IntPtr(-2147483648)}, false},
		{"max_value_boundary_valid", []*int{utils.IntPtr(0), utils.IntPtr(-2147483648), utils.IntPtr(2147483647)}, true},
		{"max_value_boundary_invalid", []*int{utils.IntPtr(2147483647), nil, utils.IntPtr(2147483647)}, false},
		{"complex_valid_bst", []*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), utils.IntPtr(3), utils.IntPtr(7), nil, utils.IntPtr(18), utils.IntPtr(2), utils.IntPtr(4), utils.IntPtr(6), utils.IntPtr(8), nil, utils.IntPtr(20)}, true},
		{"complex_invalid_at_mid_level", []*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), utils.IntPtr(3), utils.IntPtr(12), nil, utils.IntPtr(18)}, false},
		{"left_child_greater_than_root", []*int{utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}, false},
		{"right_child_less_than_root", []*int{utils.IntPtr(5), utils.IntPtr(3), utils.IntPtr(4)}, false},
		{"valid_with_negative_values", []*int{utils.IntPtr(-10), utils.IntPtr(-20), utils.IntPtr(-5), utils.IntPtr(-30), nil, utils.IntPtr(-8), utils.IntPtr(-3)}, true},
		{"invalid_with_negative_values", []*int{utils.IntPtr(-10), utils.IntPtr(-5), utils.IntPtr(-20)}, false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			rootTree := utils.SliceToTreeNode(tst.root)
			got := isValidBST(rootTree)
			if got != tst.want {
				t.Errorf("isValidBST(%v) = %v, want %v", tst.root, got, tst.want)
			}
		})
	}
}
