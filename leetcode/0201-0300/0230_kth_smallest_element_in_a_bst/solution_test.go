// Package leetcode0230 solves LeetCode 230. Kth Smallest Element in a BST
package leetcode0230

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestKthSmallestExamples(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		k    int
		want int
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(3), utils.IntPtr(1), utils.IntPtr(4), nil, utils.IntPtr(2)},
			1,
			1,
		},
		{
			"example_2",
			[]*int{utils.IntPtr(5), utils.IntPtr(3), utils.IntPtr(6), utils.IntPtr(2), utils.IntPtr(4), nil, nil, utils.IntPtr(1)},
			3,
			3,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			rootTree := utils.SliceToTreeNode(tst.root)
			got := kthSmallest(rootTree, tst.k)
			if got != tst.want {
				t.Errorf("kthSmallest(%v, %d) = %d, want %d", tst.root, tst.k, got, tst.want)
			}
		})
	}
}

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		k    int
		want int
	}{
		{"single_node", []*int{utils.IntPtr(1)}, 1, 1},
		{"single_node_min_val", []*int{utils.IntPtr(0)}, 1, 0},
		{"single_node_max_val", []*int{utils.IntPtr(10000)}, 1, 10000},
		{"two_nodes_k1", []*int{utils.IntPtr(2), utils.IntPtr(1)}, 1, 1},
		{"two_nodes_k2", []*int{utils.IntPtr(2), utils.IntPtr(1)}, 2, 2},
		{"three_nodes_k1", []*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(3)}, 1, 1},
		{"three_nodes_k2", []*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(3)}, 2, 2},
		{"three_nodes_k3", []*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(3)}, 3, 3},
		{"left_skewed_k1", []*int{utils.IntPtr(5), utils.IntPtr(4), nil, utils.IntPtr(3), nil, utils.IntPtr(2), nil, utils.IntPtr(1)}, 1, 1},
		{"left_skewed_k3", []*int{utils.IntPtr(5), utils.IntPtr(4), nil, utils.IntPtr(3), nil, utils.IntPtr(2), nil, utils.IntPtr(1)}, 3, 3},
		{"left_skewed_k5", []*int{utils.IntPtr(5), utils.IntPtr(4), nil, utils.IntPtr(3), nil, utils.IntPtr(2), nil, utils.IntPtr(1)}, 5, 5},
		{"right_skewed_k1", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)}, 1, 1},
		{"right_skewed_k3", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)}, 3, 3},
		{"right_skewed_k5", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)}, 5, 5},
		{"balanced_bst_k1", []*int{utils.IntPtr(4), utils.IntPtr(2), utils.IntPtr(6), utils.IntPtr(1), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(7)}, 1, 1},
		{"balanced_bst_k4", []*int{utils.IntPtr(4), utils.IntPtr(2), utils.IntPtr(6), utils.IntPtr(1), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(7)}, 4, 4},
		{"balanced_bst_k7", []*int{utils.IntPtr(4), utils.IntPtr(2), utils.IntPtr(6), utils.IntPtr(1), utils.IntPtr(3), utils.IntPtr(5), utils.IntPtr(7)}, 7, 7},
		{"all_same_values", []*int{utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5)}, 2, 5},
		{"deep_left_branch", []*int{utils.IntPtr(10), utils.IntPtr(5), nil, utils.IntPtr(3), nil, utils.IntPtr(2), nil, utils.IntPtr(1)}, 3, 3},
		{"deep_right_branch", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(5), utils.IntPtr(4)}, 5, 5},
		{"mixed_values_k1", []*int{utils.IntPtr(100), utils.IntPtr(50), utils.IntPtr(200), utils.IntPtr(25), utils.IntPtr(75), utils.IntPtr(150), utils.IntPtr(250)}, 1, 25},
		{"mixed_values_k4", []*int{utils.IntPtr(100), utils.IntPtr(50), utils.IntPtr(200), utils.IntPtr(25), utils.IntPtr(75), utils.IntPtr(150), utils.IntPtr(250)}, 4, 100},
		{"mixed_values_k7", []*int{utils.IntPtr(100), utils.IntPtr(50), utils.IntPtr(200), utils.IntPtr(25), utils.IntPtr(75), utils.IntPtr(150), utils.IntPtr(250)}, 7, 250},
		{"single_left_child", []*int{utils.IntPtr(3), utils.IntPtr(2), nil, utils.IntPtr(1)}, 2, 2},
		{"single_right_child", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3)}, 2, 2},
		{"sparse_tree_k1", []*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), nil, utils.IntPtr(7), nil, utils.IntPtr(20)}, 1, 5},
		{"sparse_tree_k3", []*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), nil, utils.IntPtr(7), nil, utils.IntPtr(20)}, 3, 10},
		{"sparse_tree_k5", []*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), nil, utils.IntPtr(7), nil, utils.IntPtr(20)}, 5, 20},
		{"negative_values_k1", []*int{utils.IntPtr(-2), utils.IntPtr(-3), utils.IntPtr(-1)}, 1, -3},
		{"negative_values_k2", []*int{utils.IntPtr(-2), utils.IntPtr(-3), utils.IntPtr(-1)}, 2, -2},
		{"negative_values_k3", []*int{utils.IntPtr(-2), utils.IntPtr(-3), utils.IntPtr(-1)}, 3, -1},
		{"zero_in_middle", []*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), utils.IntPtr(-5), utils.IntPtr(7), nil, nil, utils.IntPtr(-10), nil, utils.IntPtr(0)}, 4, 0},
		{"complex_bst_k1", []*int{utils.IntPtr(8), utils.IntPtr(3), utils.IntPtr(10), utils.IntPtr(1), utils.IntPtr(6), nil, utils.IntPtr(14), nil, nil, utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(13)}, 1, 1},
		{"complex_bst_k5", []*int{utils.IntPtr(8), utils.IntPtr(3), utils.IntPtr(10), utils.IntPtr(1), utils.IntPtr(6), nil, utils.IntPtr(14), nil, nil, utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(13)}, 5, 7},
		{"complex_bst_k9", []*int{utils.IntPtr(8), utils.IntPtr(3), utils.IntPtr(10), utils.IntPtr(1), utils.IntPtr(6), nil, utils.IntPtr(14), nil, nil, utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(13)}, 9, 14},
		{"root_only_k1", []*int{utils.IntPtr(42)}, 1, 42},
		{"left_full_right_empty", []*int{utils.IntPtr(5), utils.IntPtr(3), nil, utils.IntPtr(2), utils.IntPtr(4)}, 3, 4},
		{"right_full_left_empty", []*int{utils.IntPtr(1), nil, utils.IntPtr(3), utils.IntPtr(2), utils.IntPtr(5), nil, nil, utils.IntPtr(4)}, 4, 4},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			rootTree := utils.SliceToTreeNode(tst.root)
			got := kthSmallest(rootTree, tst.k)
			if got != tst.want {
				t.Errorf("kthSmallest(%v, %d) = %d, want %d", tst.root, tst.k, got, tst.want)
			}
		})
	}
}
