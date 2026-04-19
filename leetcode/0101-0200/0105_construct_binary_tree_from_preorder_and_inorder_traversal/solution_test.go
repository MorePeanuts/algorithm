package leetcode0105

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestBuildTreeExamples(t *testing.T) {
	tests := []struct {
		name      string
		preorder  []int
		inorder   []int
		wantSlice []*int
	}{
		{
			"example_1",
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			[]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)},
		},
		{
			"example_2",
			[]int{-1},
			[]int{-1},
			[]*int{utils.IntPtr(-1)},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			want := utils.SliceToTreeNode(tst.wantSlice)

			got := buildTree(tst.preorder, tst.inorder)
			if !utils.EqualTrees(got, want) {
				t.Errorf("buildTree(%v, %v) = %v, want %v", tst.preorder, tst.inorder, utils.TreeNodeToSlice(got), tst.wantSlice)
			}

			got2 := buildTree2(tst.preorder, tst.inorder)
			if !utils.EqualTrees(got2, want) {
				t.Errorf("buildTree2(%v, %v) = %v, want %v", tst.preorder, tst.inorder, utils.TreeNodeToSlice(got2), tst.wantSlice)
			}
		})
	}
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name      string
		preorder  []int
		inorder   []int
		wantSlice []*int
	}{
		{
			"left_skewed_tree",
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
		},
		{
			"right_skewed_tree",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
		},
		{
			"complete_binary_tree",
			[]int{1, 2, 4, 5, 3, 6, 7},
			[]int{4, 2, 5, 1, 6, 3, 7},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
		},
		{
			"perfect_binary_tree_height_2",
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)},
		},
		{
			"tree_with_only_left_child",
			[]int{1, 2, 3},
			[]int{3, 2, 1},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3)},
		},
		{
			"tree_with_only_right_child",
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3)},
		},
		{
			"balanced_tree_1",
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			[]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)},
		},
		{
			"negative_values",
			[]int{-1, -2, -3, -4, -5},
			[]int{-5, -4, -3, -2, -1},
			[]*int{utils.IntPtr(-1), utils.IntPtr(-2), nil, utils.IntPtr(-3), nil, utils.IntPtr(-4), nil, utils.IntPtr(-5)},
		},
		{
			"mixed_positive_negative",
			[]int{1, -2, -4, 5, 3},
			[]int{-4, -2, 5, 1, 3},
			[]*int{utils.IntPtr(1), utils.IntPtr(-2), utils.IntPtr(3), utils.IntPtr(-4), utils.IntPtr(5)},
		},
		{
			"all_same_sign_negative",
			[]int{-10, -20, -40, -50, -30},
			[]int{-40, -20, -50, -10, -30},
			[]*int{utils.IntPtr(-10), utils.IntPtr(-20), utils.IntPtr(-30), utils.IntPtr(-40), utils.IntPtr(-50)},
		},
		{
			"three_nodes_left_heavy",
			[]int{1, 2, 3},
			[]int{3, 2, 1},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3)},
		},
		{
			"three_nodes_right_heavy",
			[]int{1, 3, 2},
			[]int{1, 3, 2},
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(3), nil, utils.IntPtr(2)},
		},
		{
			"five_nodes_balanced",
			[]int{1, 2, 4, 5, 3},
			[]int{4, 2, 5, 1, 3},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5)},
		},
		{
			"five_nodes_unbalanced_left",
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
		},
		{
			"five_nodes_unbalanced_right",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
		},
		{
			"tree_height_3",
			[]int{1, 2, 4, 5, 3, 6, 7},
			[]int{4, 2, 5, 1, 6, 3, 7},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			want := utils.SliceToTreeNode(tst.wantSlice)

			got := buildTree(tst.preorder, tst.inorder)
			if !utils.EqualTrees(got, want) {
				t.Errorf("buildTree(%v, %v) = %v, want %v", tst.preorder, tst.inorder, utils.TreeNodeToSlice(got), tst.wantSlice)
			}

			got2 := buildTree2(tst.preorder, tst.inorder)
			if !utils.EqualTrees(got2, want) {
				t.Errorf("buildTree2(%v, %v) = %v, want %v", tst.preorder, tst.inorder, utils.TreeNodeToSlice(got2), tst.wantSlice)
			}
		})
	}
}
