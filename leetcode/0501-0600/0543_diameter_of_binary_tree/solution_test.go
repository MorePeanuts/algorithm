package leetcode0543

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestDiameterOfBinaryTreeExamples(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want int
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5)},
			3,
		},
		{
			"example_2",
			[]*int{utils.IntPtr(1), utils.IntPtr(2)},
			1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := diameterOfBinaryTree(root)
			if got != tst.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tst.want)
			}
		})
	}
}

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want int
	}{
		{
			"single_node",
			[]*int{utils.IntPtr(1)},
			0,
		},
		{
			"only_left_child",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil},
			1,
		},
		{
			"only_right_child",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2)},
			1,
		},
		{
			"left_skewed_tree_depth_3",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3)},
			2,
		},
		{
			"right_skewed_tree_depth_3",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3)},
			2,
		},
		{
			"left_skewed_tree_depth_5",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
			4,
		},
		{
			"right_skewed_tree_depth_5",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
			4,
		},
		{
			"full_binary_tree_depth_2",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)},
			2,
		},
		{
			"full_binary_tree_depth_3",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
			4,
		},
		{
			"full_binary_tree_depth_4",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10), utils.IntPtr(11), utils.IntPtr(12), utils.IntPtr(13), utils.IntPtr(14), utils.IntPtr(15)},
			6,
		},
		{
			"diameter_passes_through_root",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, utils.IntPtr(5)},
			4,
		},
		{
			"diameter_in_left_subtree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), nil, utils.IntPtr(6), nil, nil, utils.IntPtr(7)},
			5,
		},
		{
			"diameter_in_right_subtree",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), nil, utils.IntPtr(6), nil, nil, utils.IntPtr(7)},
			5,
		},
		{
			"unbalanced_left_heavy",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), nil, nil, utils.IntPtr(6), utils.IntPtr(7), nil, nil, utils.IntPtr(8)},
			5,
		},
		{
			"unbalanced_right_heavy",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, nil, utils.IntPtr(4), utils.IntPtr(5), nil, nil, nil, utils.IntPtr(6), utils.IntPtr(7), nil, utils.IntPtr(8)},
			6,
		},
		{
			"all_same_values",
			[]*int{utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5)},
			3,
		},
		{
			"negative_values",
			[]*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3), utils.IntPtr(-4), utils.IntPtr(-5), nil, utils.IntPtr(-6)},
			4,
		},
		{
			"mixed_positive_negative",
			[]*int{utils.IntPtr(0), utils.IntPtr(-1), utils.IntPtr(1), utils.IntPtr(-2), nil, nil, utils.IntPtr(2), utils.IntPtr(-3)},
			5,
		},
		{
			"zigzag_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, nil, utils.IntPtr(3), utils.IntPtr(4), nil, nil, utils.IntPtr(5)},
			4,
		},
		{
			"tree_with_null_gaps",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(5), nil, nil, nil, utils.IntPtr(6), utils.IntPtr(7), nil, utils.IntPtr(8)},
			6,
		},
		{
			"sparse_left_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(5), nil, utils.IntPtr(6), nil, utils.IntPtr(7)},
			4,
		},
		{
			"sparse_right_tree",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), utils.IntPtr(4), nil, utils.IntPtr(5), nil, utils.IntPtr(6), nil, utils.IntPtr(7)},
			6,
		},
		{
			"balanced_tree_depth_5",
			[]*int{
				utils.IntPtr(1),
				utils.IntPtr(2), utils.IntPtr(3),
				utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7),
				utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10), utils.IntPtr(11), utils.IntPtr(12), utils.IntPtr(13), utils.IntPtr(14), utils.IntPtr(15),
				utils.IntPtr(16), utils.IntPtr(17), utils.IntPtr(18), utils.IntPtr(19), utils.IntPtr(20), utils.IntPtr(21), utils.IntPtr(22), utils.IntPtr(23),
			},
			7,
		},
		{
			"two_nodes_left",
			[]*int{utils.IntPtr(10), utils.IntPtr(20)},
			1,
		},
		{
			"two_nodes_right",
			[]*int{utils.IntPtr(10), nil, utils.IntPtr(20)},
			1,
		},
		{
			"three_nodes_chain_left",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3)},
			2,
		},
		{
			"three_nodes_chain_right",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3)},
			2,
		},
		{
			"symmetric_tree_depth_3",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3)},
			4,
		},
		{
			"asymmetric_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(5), nil, utils.IntPtr(6), utils.IntPtr(7), nil, utils.IntPtr(8)},
			6,
		},
		{
			"tree_with_many_leaves",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), nil, nil, nil, utils.IntPtr(8), nil, nil, nil, utils.IntPtr(9)},
			6,
		},
		{
			"deep_left_subtree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5), utils.IntPtr(6), nil, utils.IntPtr(7), nil, utils.IntPtr(8)},
			5,
		},
		{
			"deep_right_subtree",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5), utils.IntPtr(6), nil, utils.IntPtr(7), nil, utils.IntPtr(8)},
			7,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := diameterOfBinaryTree(root)
			if got != tst.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tst.want)
			}
		})
	}
}
