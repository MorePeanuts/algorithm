// Package leetcode0572 solves LeetCode 572. Subtree of Another Tree
package leetcode0572

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestIsSubtreeExamples(t *testing.T) {
	tests := []struct {
		name    string
		root    []*int
		subRoot []*int
		want    bool
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(1), utils.IntPtr(2)},
			[]*int{utils.IntPtr(4), utils.IntPtr(1), utils.IntPtr(2)},
			true,
		},
		{
			"example_2",
			[]*int{utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(1), utils.IntPtr(2), nil, nil, nil, nil, utils.IntPtr(0)},
			[]*int{utils.IntPtr(4), utils.IntPtr(1), utils.IntPtr(2)},
			false,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			rootTree := utils.SliceToTreeNode(tst.root)
			subRootTree := utils.SliceToTreeNode(tst.subRoot)
			got := isSubtree(rootTree, subRootTree)
			if got != tst.want {
				t.Errorf("isSubtree(%v, %v) = %v, want %v", tst.root, tst.subRoot, got, tst.want)
			}

			got2 := isSubtree2(rootTree, subRootTree)
			if got2 != tst.want {
				t.Errorf("isSubtree2(%v, %v) = %v, want %v", tst.root, tst.subRoot, got2, tst.want)
			}
		})
	}
}

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name    string
		root    []*int
		subRoot []*int
		want    bool
	}{
		// Single node cases
		{"single_node_same", []*int{utils.IntPtr(1)}, []*int{utils.IntPtr(1)}, true},
		{"single_node_diff", []*int{utils.IntPtr(1)}, []*int{utils.IntPtr(2)}, false},

		// Subtree is root itself
		{"subtree_is_root", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}, []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3)}, true},

		// Empty subtree (but according to constraints, subRoot has at least 1 node)
		// But testing edge case behavior
		{"root_empty_subroot_not", []*int{}, []*int{utils.IntPtr(1)}, false},

		// Subtree at different positions
		{"subtree_at_left_child", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5)}, []*int{utils.IntPtr(2), utils.IntPtr(4), utils.IntPtr(5)}, true},
		{"subtree_at_right_child_single", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5)}, []*int{utils.IntPtr(3)}, true},

		// Subtree with similar structure but different values
		{"similar_structure_diff_values", []*int{utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(1), utils.IntPtr(2)}, []*int{utils.IntPtr(4), utils.IntPtr(1), utils.IntPtr(3)}, false},

		// Subtree with extra node in root (subtree must include all descendants)
		{"subtree_with_extra_node_in_root", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3)}, []*int{utils.IntPtr(2), utils.IntPtr(3)}, true},

		// Deep subtree
		{"deep_subtree", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9)}, []*int{utils.IntPtr(2), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(8), utils.IntPtr(9)}, true},
		{"deep_subtree_not_matching", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9)}, []*int{utils.IntPtr(2), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(8), utils.IntPtr(10)}, false},

		// Negative values
		{"negative_values_match", []*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3)}, []*int{utils.IntPtr(-2)}, true},
		{"negative_values_no_match", []*int{utils.IntPtr(-1), utils.IntPtr(-2), utils.IntPtr(-3)}, []*int{utils.IntPtr(2)}, false},

		// Mixed positive and negative
		{"mixed_values_match", []*int{utils.IntPtr(1), utils.IntPtr(-2), utils.IntPtr(3), utils.IntPtr(-4), utils.IntPtr(5)}, []*int{utils.IntPtr(-2), utils.IntPtr(-4), utils.IntPtr(5)}, true},

		// Edge values (-10^4 to 10^4)
		{"edge_values_match", []*int{utils.IntPtr(10000), utils.IntPtr(-10000)}, []*int{utils.IntPtr(-10000)}, true},
		{"edge_values_root_match", []*int{utils.IntPtr(10000), utils.IntPtr(-10000)}, []*int{utils.IntPtr(10000), utils.IntPtr(-10000)}, true},

		// Unbalanced trees
		{"unbalanced_root_single_node_subtree", []*int{utils.IntPtr(1), utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, []*int{utils.IntPtr(3)}, false},
		{"unbalanced_subtree_match", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, utils.IntPtr(5), nil, utils.IntPtr(6)}, []*int{utils.IntPtr(4), nil, utils.IntPtr(6)}, true},

		// Subtree with only left children
		{"left_chain_subtree", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), nil, nil, nil, utils.IntPtr(5)}, []*int{utils.IntPtr(4), utils.IntPtr(5)}, true},

		// Subtree with only right children
		{"right_chain_subtree", []*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4)}, []*int{utils.IntPtr(2), nil, utils.IntPtr(3)}, false},

		// Partial matches that aren't exact
		{"partial_match_not_exact", []*int{utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(6)}, []*int{utils.IntPtr(4), utils.IntPtr(1)}, false},

		// Subtree is deeper in the tree
		{"subtree_deep_in_tree", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10), utils.IntPtr(11)}, []*int{utils.IntPtr(5), utils.IntPtr(10), utils.IntPtr(11)}, true},

		// Same values but different structure
		{"same_values_diff_structure", []*int{utils.IntPtr(1), utils.IntPtr(1), utils.IntPtr(1)}, []*int{utils.IntPtr(1), utils.IntPtr(1)}, false},

		// Large subtree
		{"large_subtree", []*int{utils.IntPtr(0), utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9), utils.IntPtr(10), utils.IntPtr(11), utils.IntPtr(12), utils.IntPtr(13), utils.IntPtr(14)}, []*int{utils.IntPtr(2), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(11), utils.IntPtr(12), utils.IntPtr(13), utils.IntPtr(14)}, true},

		// Subtree that almost matches but has one different value deep down
		{"almost_match_deep_diff", []*int{utils.IntPtr(1), utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7), utils.IntPtr(8), utils.IntPtr(9)}, []*int{utils.IntPtr(2), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(8), utils.IntPtr(10)}, false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			rootTree := utils.SliceToTreeNode(tst.root)
			subRootTree := utils.SliceToTreeNode(tst.subRoot)
			got := isSubtree(rootTree, subRootTree)
			if got != tst.want {
				t.Errorf("isSubtree(%v, %v) = %v, want %v", tst.root, tst.subRoot, got, tst.want)
			}

			got2 := isSubtree2(rootTree, subRootTree)
			if got2 != tst.want {
				t.Errorf("isSubtree2(%v, %v) = %v, want %v", tst.root, tst.subRoot, got2, tst.want)
			}
		})
	}
}
