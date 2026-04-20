// Package leetcode0297 solves LeetCode 297. Serialize and Deserialize Binary Tree
package leetcode0297

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestCodecSerializeDeserializeExamples(t *testing.T) {
	tests := []struct {
		name string
		tree []*int
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, nil, utils.IntPtr(4), utils.IntPtr(5)},
		},
		{
			"example_2",
			[]*int{},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.tree)
			codec := Constructor()
			data := codec.serialize(root)
			got := codec.deserialize(data)

			if !utils.EqualTrees(root, got) {
				t.Errorf("serialize/deserialize failed for tree %v", tst.tree)
			}
		})
	}
}

func TestCodecSerializeDeserialize(t *testing.T) {
	tests := []struct {
		name string
		tree []*int
	}{
		{
			"single_node",
			[]*int{utils.IntPtr(1)},
		},
		{
			"left_skewed_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)},
		},
		{
			"right_skewed_tree",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)},
		},
		{
			"full_binary_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)},
		},
		{
			"complete_binary_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6)},
		},
		{
			"tree_with_duplicate_values_1",
			[]*int{utils.IntPtr(1), utils.IntPtr(1)},
		},
		{
			"deep_left_tree",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
		},
		{
			"deep_right_tree",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
		},
		{
			"balanced_tree_1",
			[]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)},
		},
		{
			"balanced_tree_2",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(4), utils.IntPtr(3)},
		},
		{
			"tree_with_nils_in_middle",
			[]*int{utils.IntPtr(5), utils.IntPtr(4), utils.IntPtr(8), utils.IntPtr(11), nil, utils.IntPtr(13), utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(2), nil, nil, nil, utils.IntPtr(1)},
		},
		{
			"small_tree_1",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)},
		},
		{
			"small_tree_2",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2)},
		},
		{
			"small_tree_3",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil},
		},
		{
			"tree_with_three_levels",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10), utils.IntPtr(11), utils.IntPtr(12), utils.IntPtr(13), utils.IntPtr(14), utils.IntPtr(15)},
		},
		{
			"sparse_tree_1",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), utils.IntPtr(4), nil, utils.IntPtr(5)},
		},
		{
			"sparse_tree_2",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5), nil, utils.IntPtr(6)},
		},
		{
			"complex_tree",
			[]*int{utils.IntPtr(4), utils.IntPtr(-7), utils.IntPtr(-3), nil, nil, utils.IntPtr(-9), utils.IntPtr(-3), utils.IntPtr(9), utils.IntPtr(-7), utils.IntPtr(-4), nil, utils.IntPtr(6), nil, utils.IntPtr(-6), utils.IntPtr(-6), nil, nil, utils.IntPtr(0), utils.IntPtr(6), utils.IntPtr(5), nil, utils.IntPtr(9), nil, nil, utils.IntPtr(-1), utils.IntPtr(-4), nil, nil, nil, utils.IntPtr(-2)},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.tree)
			codec := Constructor()
			data := codec.serialize(root)
			got := codec.deserialize(data)

			if !utils.EqualTrees(root, got) {
				t.Errorf("serialize/deserialize failed for tree %v", tst.tree)
			}
		})
	}
}
