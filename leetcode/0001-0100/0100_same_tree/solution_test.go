// Package leetcode0100 solves LeetCode 100. Same Tree
package leetcode0100

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestIsSameTreeExamples(t *testing.T) {
	tests := []struct {
		name string
		p    []*int
		q    []*int
		want bool
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)},
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)},
			true,
		},
		{
			"example_2",
			[]*int{utils.IntPtr(1), utils.IntPtr(2)},
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2)},
			false,
		},
		{
			"example_3",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(1)},
			[]*int{utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(2)},
			false,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			pTree := utils.SliceToTreeNode(tst.p)
			qTree := utils.SliceToTreeNode(tst.q)
			got := isSameTree(pTree, qTree)
			if got != tst.want {
				t.Errorf("isSameTree(%v, %v) = %v, want %v", tst.p, tst.q, got, tst.want)
			}
		})
	}
}

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    []*int
		q    []*int
		want bool
	}{
		{"both_empty", []*int{}, []*int{}, true},
		{"one_empty_one_single", []*int{}, []*int{utils.IntPtr(1)}, false},
		{"one_single_one_empty", []*int{utils.IntPtr(1)}, []*int{}, false},
		{"single_node_same", []*int{utils.IntPtr(5)}, []*int{utils.IntPtr(5)}, true},
		{"single_node_diff", []*int{utils.IntPtr(5)}, []*int{utils.IntPtr(6)}, false},
		{"same_structure_diff_values", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}, []*int{utils.IntPtr(1), utils.IntPtr(4), utils.IntPtr(3)}, false},
		{"left_subtree_diff", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}, []*int{utils.IntPtr(1), nil, utils.IntPtr(3)}, false},
		{"right_subtree_diff", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}, []*int{utils.IntPtr(1), utils.IntPtr(2), nil}, false},
		{"deep_same_tree", []*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}, []*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}, true},
		{"deep_diff_tree", []*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)}, []*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(8)}, false},
		{"unbalanced_same", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, true},
		{"unbalanced_diff", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(5)}, false},
		{"negative_values_same", []*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3)}, []*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3)}, true},
		{"negative_values_diff", []*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3)}, []*int{utils.IntPtr(-1), utils.IntPtr(2), utils.IntPtr(-3)}, false},
		{"max_min_values", []*int{utils.IntPtr(10000), utils.IntPtr(-10000)}, []*int{utils.IntPtr(10000), utils.IntPtr(-10000)}, true},
		{"max_min_values_diff", []*int{utils.IntPtr(10000), utils.IntPtr(-10000)}, []*int{utils.IntPtr(10000), utils.IntPtr(-9999)}, false},
		{"one_side_long_chain", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)}, []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)}, true},
		{"one_side_long_chain_diff", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)}, []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(5), nil, utils.IntPtr(5)}, false},
		{"mirror_structure", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}, []*int{utils.IntPtr(1), utils.IntPtr(3), utils.IntPtr(2)}, false},
		{"same_nodes_different_order", []*int{utils.IntPtr(2), utils.IntPtr(1), utils.IntPtr(3)}, []*int{utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(1)}, false},
		{"single_child_left", []*int{utils.IntPtr(1), utils.IntPtr(2)}, []*int{utils.IntPtr(1), utils.IntPtr(2)}, true},
		{"single_child_right", []*int{utils.IntPtr(1), nil, utils.IntPtr(2)}, []*int{utils.IntPtr(1), nil, utils.IntPtr(2)}, true},
		{"single_child_left_vs_right", []*int{utils.IntPtr(1), utils.IntPtr(2)}, []*int{utils.IntPtr(1), nil, utils.IntPtr(2)}, false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			pTree := utils.SliceToTreeNode(tst.p)
			qTree := utils.SliceToTreeNode(tst.q)
			got := isSameTree(pTree, qTree)
			if got != tst.want {
				t.Errorf("isSameTree(%v, %v) = %v, want %v", tst.p, tst.q, got, tst.want)
			}
		})
	}
}
