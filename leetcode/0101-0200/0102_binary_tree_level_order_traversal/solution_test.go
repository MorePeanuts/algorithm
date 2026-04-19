// Package leetcode0102 solves LeetCode 102. Binary Tree Level Order Traversal
package leetcode0102

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestLevelOrderExamples(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want [][]int
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)},
			[][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			"example_2",
			[]*int{utils.IntPtr(1)},
			[][]int{{1}},
		},
		{
			"example_3",
			[]*int{},
			[][]int{},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			rootTree := utils.SliceToTreeNode(tst.root)
			got := levelOrder(rootTree)
			if !equal2DIntSlice(got, tst.want) {
				t.Errorf("levelOrder(%v) = %v, want %v", tst.root, got, tst.want)
			}
		})
	}
}

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want [][]int
	}{
		{"empty_tree", []*int{}, [][]int{}},
		{"single_node", []*int{utils.IntPtr(1)}, [][]int{{1}}},
		{"single_node_negative", []*int{utils.IntPtr(-100)}, [][]int{{-100}}},
		{"single_node_max", []*int{utils.IntPtr(1000)}, [][]int{{1000}}},
		{"single_node_min", []*int{utils.IntPtr(-1000)}, [][]int{{-1000}}},
		{"only_left_child", []*int{utils.IntPtr(1), utils.IntPtr(2)}, [][]int{{1}, {2}}},
		{"only_right_child", []*int{utils.IntPtr(1), nil, utils.IntPtr(2)}, [][]int{{1}, {2}}},
		{"two_levels_complete", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}, [][]int{{1}, {2, 3}}},
		{"left_skewed_tree", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, [][]int{{1}, {2}, {3}, {4}}},
		{"right_skewed_tree", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, [][]int{{1}, {2}, {3}, {4}}},
		{"three_levels_complete", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}, [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
		{"missing_middle_nodes", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, nil, utils.IntPtr(6), utils.IntPtr(7)}, [][]int{{1}, {2, 3}, {6, 7}}},
		{"sparse_left_subtree", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, nil}, [][]int{{1}, {2, 3}, {4}}},
		{"sparse_right_subtree", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, nil, nil, utils.IntPtr(7)}, [][]int{{1}, {2, 3}, {7}}},
		{"zigzag_shape", []*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"alternating_nodes", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), utils.IntPtr(3), nil, nil, utils.IntPtr(4), nil, utils.IntPtr(5)}, [][]int{{1}, {2}, {3}, {4}, {5}}},
		{"negative_values_tree", []*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3), utils.IntPtr(-4), utils.IntPtr(-5)}, [][]int{{-1}, {-2, -3}, {-4, -5}}},
		{"mixed_values_tree", []*int{utils.IntPtr(0), utils.IntPtr(-1), utils.IntPtr(1), utils.IntPtr(-100), utils.IntPtr(-50), utils.IntPtr(50), utils.IntPtr(100)}, [][]int{{0}, {-1, 1}, {-100, -50, 50, 100}}},
		{"four_levels_complete", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10), utils.IntPtr(11), utils.IntPtr(12), utils.IntPtr(13), utils.IntPtr(14), utils.IntPtr(15)}, [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14, 15}}},
		{"deep_left_chain", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5), nil, utils.IntPtr(6)}, [][]int{{1}, {2}, {3}, {4}, {5}, {6}}},
		{"deep_right_chain", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5), nil, utils.IntPtr(6)}, [][]int{{1}, {2}, {3}, {4}, {5}, {6}}},
		{"symmetric_tree", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3)}, [][]int{{1}, {2, 2}, {3, 4, 4, 3}}},
		{"asymmetric_tree", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(3)}, [][]int{{1}, {2, 2}, {3, 3}}},
		{"single_left_grandchild", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3)}, [][]int{{1}, {2}, {3}}},
		{"single_right_grandchild", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3)}, [][]int{{1}, {2}, {3}}},
		{"mixed_null_pattern", []*int{utils.IntPtr(5), utils.IntPtr(4), utils.IntPtr(8), utils.IntPtr(11), nil, utils.IntPtr(13), utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(2), nil, nil, nil, utils.IntPtr(1)}, [][]int{{5}, {4, 8}, {11, 13, 4}, {7, 2, 1}}},
		{"all_same_values", []*int{utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5)}, [][]int{{5}, {5, 5}, {5, 5, 5, 5}}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			rootTree := utils.SliceToTreeNode(tst.root)
			got := levelOrder(rootTree)
			if !equal2DIntSlice(got, tst.want) {
				t.Errorf("levelOrder(%v) = %v, want %v", tst.root, got, tst.want)
			}
		})
	}
}

// equal2DIntSlice checks if two 2D int slices are exactly equal in order and values.
func equal2DIntSlice(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if len(got[i]) != len(want[i]) {
			return false
		}
		for j := range got[i] {
			if got[i][j] != want[i][j] {
				return false
			}
		}
	}
	return true
}
