package leetcode0110

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestIsBalancedExamples(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want bool
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(3), utils.IntPtr(9), utils.IntPtr(20), nil, nil, utils.IntPtr(15), utils.IntPtr(7)},
			true,
		},
		{
			"example_2",
			[]*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(3), nil, nil, utils.IntPtr(4), utils.IntPtr(4)},
			false,
		},
		{
			"example_3",
			[]*int{},
			true,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := isBalanced(root)
			if got != tst.want {
				t.Errorf("isBalanced() = %v, want %v", got, tst.want)
			}
		})
	}
}

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name string
		root []*int
		want bool
	}{
		{"single_node", []*int{utils.IntPtr(1)}, true},
		{"two_nodes_left", []*int{utils.IntPtr(1), utils.IntPtr(2)}, true},
		{"two_nodes_right", []*int{utils.IntPtr(1), nil, utils.IntPtr(2)}, true},
		{"three_nodes_complete", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}, true},
		{"left_skewed_3", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3)}, false},
		{"right_skewed_3", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3)}, false},
		{"left_skewed_4", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, false},
		{"right_skewed_4", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, false},
		{"balanced_depth_3", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)}, true},
		{"almost_balanced_left", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, nil}, true},
		{"almost_balanced_right", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, nil, nil, utils.IntPtr(4)}, true},
		{"unbalanced_left_deeper", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), nil, nil, utils.IntPtr(6)}, false},
		{"unbalanced_right_deeper", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, nil, utils.IntPtr(4), utils.IntPtr(5), nil, nil, nil, utils.IntPtr(6)}, false},
		{"balanced_with_single_child", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(5), nil}, true},
		{"unbalanced_in_subtree", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, utils.IntPtr(5), utils.IntPtr(6), nil, nil, utils.IntPtr(7)}, false},
		{"balanced_perfect_4", []*int{
			utils.IntPtr(1),
			utils.IntPtr(2), utils.IntPtr(3),
			utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7),
		}, true},
		{"unbalanced_at_root", []*int{
			utils.IntPtr(1),
			utils.IntPtr(2), nil,
			utils.IntPtr(3), nil,
			utils.IntPtr(4), nil,
		}, false},
		{"complex_balanced", []*int{
			utils.IntPtr(1),
			utils.IntPtr(2), utils.IntPtr(3),
			utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7),
			utils.IntPtr(8), nil, utils.IntPtr(9), nil, nil, nil, utils.IntPtr(10), nil,
		}, true},
		{"complex_unbalanced", []*int{
			utils.IntPtr(1),
			utils.IntPtr(2), utils.IntPtr(3),
			utils.IntPtr(4), utils.IntPtr(5), nil, nil,
			utils.IntPtr(6), utils.IntPtr(7), nil, nil,
			utils.IntPtr(8), nil,
		}, false},
		{"single_left_chain_5", []*int{
			utils.IntPtr(1),
			utils.IntPtr(2), nil,
			utils.IntPtr(3), nil,
			utils.IntPtr(4), nil,
			utils.IntPtr(5),
		}, false},
		{"single_right_chain_5", []*int{
			utils.IntPtr(1),
			nil, utils.IntPtr(2),
			nil, utils.IntPtr(3),
			nil, utils.IntPtr(4),
			nil, utils.IntPtr(5),
		}, false},
		{"balanced_but_not_perfect", []*int{
			utils.IntPtr(1),
			utils.IntPtr(2), utils.IntPtr(3),
			utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), nil,
			utils.IntPtr(7), utils.IntPtr(8), nil, nil, utils.IntPtr(9), nil,
		}, false},
		{"all_left_children_depth_3", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil}, false},
		{"all_right_children_depth_3", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3)}, false},
		{"zig_zag_balanced", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), nil, utils.IntPtr(4), utils.IntPtr(5), nil}, true},
		{"zig_zag_unbalanced", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, false},
		{"large_balanced_tree_15_nodes", []*int{
			utils.IntPtr(1),
			utils.IntPtr(2), utils.IntPtr(3),
			utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7),
			utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10), utils.IntPtr(11), utils.IntPtr(12), utils.IntPtr(13), utils.IntPtr(14), utils.IntPtr(15),
		}, true},
		{"large_unbalanced_tree", []*int{
			utils.IntPtr(1),
			utils.IntPtr(2), utils.IntPtr(3),
			utils.IntPtr(4), utils.IntPtr(5), nil, nil,
			utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9),
			utils.IntPtr(10), nil,
		}, false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.root)
			got := isBalanced(root)
			if got != tst.want {
				t.Errorf("isBalanced() = %v, want %v", got, tst.want)
			}
		})
	}
}
