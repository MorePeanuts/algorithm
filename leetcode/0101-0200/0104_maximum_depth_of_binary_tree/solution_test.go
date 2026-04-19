package leetcode0104

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestMaxDepthExamples(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want int
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)},
			3,
		},
		{
			"example_2",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2)},
			2,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := maxDepth(root)
			if got != tst.want {
				t.Errorf("maxDepth() = %v, want %v", got, tst.want)
			}

			// Test second solution if it exists
			root2 := utils.SliceToTreeNode(tst.root)
			got2 := maxDepth2(root2)
			if got2 != tst.want {
				t.Errorf("maxDepth2() = %v, want %v", got2, tst.want)
			}
		})
	}
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want int
	}{
		{
			"empty_tree",
			[]*int{},
			0,
		},
		{
			"single_node",
			[]*int{utils.IntPtr(1)},
			1,
		},
		{
			"left_skewed_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil},
			4,
		},
		{
			"right_skewed_tree",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)},
			4,
		},
		{
			"only_left_child",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil},
			2,
		},
		{
			"only_right_child",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2)},
			2,
		},
		{
			"full_binary_tree_depth_3",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
			3,
		},
		{
			"full_binary_tree_depth_4",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10), utils.IntPtr(11), utils.IntPtr(12), utils.IntPtr(13), utils.IntPtr(14), utils.IntPtr(15)},
			4,
		},
		{
			"all_same_values",
			[]*int{utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5)},
			3,
		},
		{
			"negative_values",
			[]*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3), utils.IntPtr(-4), utils.IntPtr(-5), nil, utils.IntPtr(-6)},
			3,
		},
		{
			"mixed_positive_negative",
			[]*int{utils.IntPtr(0), utils.IntPtr(-1), utils.IntPtr(1), utils.IntPtr(-2), nil, nil, utils.IntPtr(2), utils.IntPtr(-3)},
			4,
		},
		{
			"unbalanced_left_heavy",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), nil, nil, utils.IntPtr(6), utils.IntPtr(7)},
			4,
		},
		{
			"unbalanced_right_heavy",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, nil, utils.IntPtr(4), utils.IntPtr(5), nil, nil, nil, utils.IntPtr(6)},
			4,
		},
		{
			"alternating_levels",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
			5,
		},
		{
			"sparse_tree",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5), nil, utils.IntPtr(6)},
			6,
		},
		{
			"tree_with_many_leaves",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), nil, nil, nil, utils.IntPtr(8), nil, nil, nil, utils.IntPtr(9)},
			4,
		},
		{
			"single_left_chain",
			[]*int{utils.IntPtr(10), utils.IntPtr(9), nil, utils.IntPtr(8), nil, utils.IntPtr(7), nil, utils.IntPtr(6), nil, utils.IntPtr(5)},
			6,
		},
		{
			"single_right_chain",
			[]*int{utils.IntPtr(5), nil, utils.IntPtr(6), nil, utils.IntPtr(7), nil, utils.IntPtr(8), nil, utils.IntPtr(9), nil, utils.IntPtr(10)},
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
			5,
		},
		{
			"tree_with_null_gaps",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(5), nil, nil, nil, utils.IntPtr(6)},
			4,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := maxDepth(root)
			if got != tst.want {
				t.Errorf("maxDepth() = %v, want %v", got, tst.want)
			}

			// Test second solution if it exists
			root2 := utils.SliceToTreeNode(tst.root)
			got2 := maxDepth2(root2)
			if got2 != tst.want {
				t.Errorf("maxDepth2() = %v, want %v", got2, tst.want)
			}
		})
	}
}
