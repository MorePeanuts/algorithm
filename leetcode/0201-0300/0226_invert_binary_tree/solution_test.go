package leetcode0226

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestInvertTreeExamples(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want []*int
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(4), utils.IntPtr(2), utils.IntPtr(7), utils.IntPtr(1), utils.IntPtr(3), utils.IntPtr(6), utils.IntPtr(9)},
			[]*int{utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(2), utils.IntPtr(9), utils.IntPtr(6), utils.IntPtr(3), utils.IntPtr(1)},
		},
		{
			"example_2",
			[]*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(3)},
			[]*int{utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(1)},
		},
		{
			"example_3",
			[]*int{},
			[]*int{},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			want := utils.SliceToTreeNode(tst.want)

			got := invertTree(root)
			if !utils.EqualTrees(got, want) {
				t.Errorf("invertTree() = %v, want %v", utils.TreeNodeToSlice(got), utils.TreeNodeToSlice(want))
			}

			// Reconstruct tree for second solution test
			root2 := utils.SliceToTreeNode(tst.root)
			got2 := invertTree2(root2)
			if !utils.EqualTrees(got2, want) {
				t.Errorf("invertTree2() = %v, want %v", utils.TreeNodeToSlice(got2), utils.TreeNodeToSlice(want))
			}
		})
	}
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want []*int
	}{
		{
			"single_node",
			[]*int{utils.IntPtr(1)},
			[]*int{utils.IntPtr(1)},
		},
		{
			"left_skewed_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil},
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3)},
		},
		{
			"right_skewed_tree",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3)},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil},
		},
		{
			"only_left_child",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil},
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2)},
		},
		{
			"only_right_child",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2)},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil},
		},
		{
			"full_binary_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
			[]*int{utils.IntPtr(1), utils.IntPtr(3), utils.IntPtr(2), utils.IntPtr(7), utils.IntPtr(6), utils.IntPtr(5), utils.IntPtr(4)},
		},
		{
			"all_same_values",
			[]*int{utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5)},
			[]*int{utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5), utils.IntPtr(5)},
		},
		{
			"negative_values",
			[]*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3), utils.IntPtr(-4), utils.IntPtr(-5)},
			[]*int{utils.IntPtr(-1), utils.IntPtr(-3), utils.IntPtr(-2), nil, nil, utils.IntPtr(-5), utils.IntPtr(-4)},
		},
		{
			"mixed_positive_negative",
			[]*int{utils.IntPtr(0), utils.IntPtr(-1), utils.IntPtr(1), utils.IntPtr(-2), utils.IntPtr(-1), utils.IntPtr(1), utils.IntPtr(2)},
			[]*int{utils.IntPtr(0), utils.IntPtr(1), utils.IntPtr(-1), utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(-1), utils.IntPtr(-2)},
		},
		{
			"symmetric_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3)},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3)},
		},
		{
			"asymmetric_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(5), nil},
			[]*int{utils.IntPtr(1), utils.IntPtr(3), utils.IntPtr(2), nil, utils.IntPtr(5), utils.IntPtr(4), nil},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			want := utils.SliceToTreeNode(tst.want)

			got := invertTree(root)
			if !utils.EqualTrees(got, want) {
				t.Errorf("invertTree() = %v, want %v", utils.TreeNodeToSlice(got), utils.TreeNodeToSlice(want))
			}

			// Reconstruct tree for second solution test
			root2 := utils.SliceToTreeNode(tst.root)
			got2 := invertTree2(root2)
			if !utils.EqualTrees(got2, want) {
				t.Errorf("invertTree2() = %v, want %v", utils.TreeNodeToSlice(got2), utils.TreeNodeToSlice(want))
			}
		})
	}
}
