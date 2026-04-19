package leetcode0199

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestRightSideViewExamples(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want []int
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(5), nil, utils.IntPtr(4)},
			[]int{1, 3, 4},
		},
		{
			"example_2",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, nil, utils.IntPtr(5)},
			[]int{1, 3, 4, 5},
		},
		{
			"example_3",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(3)},
			[]int{1, 3},
		},
		{
			"example_4",
			[]*int{},
			[]int{},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := rightSideView(root)
			if !utils.MatchIntSlice(got, tst.want) {
				t.Errorf("rightSideView(%v) = %v, want %v", tst.root, got, tst.want)
			}

			got = rightSideView2(root)
			if !utils.MatchIntSlice(got, tst.want) {
				t.Errorf("rightSideView2(%v) = %v, want %v", tst.root, got, tst.want)
			}
		})
	}
}

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want []int
	}{
		{
			"single_node",
			[]*int{utils.IntPtr(1)},
			[]int{1},
		},
		{
			"left_skewed_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)},
			[]int{1, 2, 3, 4},
		},
		{
			"right_skewed_tree",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)},
			[]int{1, 2, 3, 4},
		},
		{
			"left_longer_than_right",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4)},
			[]int{1, 3, 4},
		},
		{
			"right_longer_than_left",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, nil, nil, utils.IntPtr(4)},
			[]int{1, 3, 4},
		},
		{
			"full_binary_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
			[]int{1, 3, 7},
		},
		{
			"perfect_binary_tree_depth_4",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10), utils.IntPtr(11), utils.IntPtr(12), utils.IntPtr(13), utils.IntPtr(14), utils.IntPtr(15)},
			[]int{1, 3, 7, 15},
		},
		{
			"left_node_at_deeper_level",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(5), utils.IntPtr(6), nil, utils.IntPtr(7)},
			[]int{1, 3, 6, 7},
		},
		{
			"alternating_left_right",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"alternating_right_left",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"two_nodes_left_only",
			[]*int{utils.IntPtr(1), utils.IntPtr(2)},
			[]int{1, 2},
		},
		{
			"two_nodes_right_only",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2)},
			[]int{1, 2},
		},
		{
			"middle_left_missing",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, nil, utils.IntPtr(4), utils.IntPtr(5), nil, utils.IntPtr(6)},
			[]int{1, 3, 5, 6},
		},
		{
			"middle_right_missing",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), nil, nil, utils.IntPtr(6)},
			[]int{1, 3, 5, 6},
		},
		{
			"sparse_tree_1",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, utils.IntPtr(5), utils.IntPtr(6), nil, nil, utils.IntPtr(7)},
			[]int{1, 3, 5, 7},
		},
		{
			"sparse_tree_2",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), nil, nil, nil, utils.IntPtr(6)},
			[]int{1, 2, 3, 5, 6},
		},
		{
			"negative_values",
			[]*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3), nil, utils.IntPtr(-5), nil, utils.IntPtr(-4)},
			[]int{-1, -3, -4},
		},
		{
			"mixed_values",
			[]*int{utils.IntPtr(100), utils.IntPtr(-100), utils.IntPtr(50), nil, utils.IntPtr(-50), nil, utils.IntPtr(75)},
			[]int{100, 50, 75},
		},
		{
			"all_same_values",
			[]*int{utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5)},
			[]int{5, 5, 5},
		},
		{
			"zigzag_right_view",
			[]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7), utils.IntPtr(1), nil, nil, utils.IntPtr(2)},
			[]int{3, 20, 7, 2},
		},
		{
			"max_nodes_100",
			func() []*int {
				res := make([]*int, 100)
				for i := 0; i < 100; i++ {
					val := i + 1
					res[i] = &val
				}
				return res
			}(),
			[]int{1, 3, 7, 15, 31, 63, 100},
		},
		{
			"left_child_has_right_only",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, nil, utils.IntPtr(3), nil, utils.IntPtr(4)},
			[]int{1, 2, 3, 4},
		},
		{
			"right_child_has_left_only",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4)},
			[]int{1, 2, 3, 4},
		},
		{
			"complex_tree_1",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), nil, utils.IntPtr(7), utils.IntPtr(8), nil, utils.IntPtr(9), utils.IntPtr(10), nil, nil, utils.IntPtr(11)},
			[]int{1, 3, 6, 10, 11},
		},
		{
			"complex_tree_2",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), nil, nil, nil, utils.IntPtr(8), utils.IntPtr(9), nil, utils.IntPtr(10)},
			[]int{1, 3, 5, 7, 9, 10},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := rightSideView(root)
			if !utils.MatchIntSlice(got, tst.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tst.want)
			}

			got = rightSideView2(root)
			if !utils.MatchIntSlice(got, tst.want) {
				t.Errorf("rightSideView2() = %v, want %v", got, tst.want)
			}
		})
	}
}
