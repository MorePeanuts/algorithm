package leetcode0235

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestLowestCommonAncestorExamples(t *testing.T) {
	tests := []struct {
		name     string
		rootVals []*int
		pVal     int
		qVal     int
		wantVal  int
	}{
		{
			"example_1",
			[]*int{utils.IntPtr(6), utils.IntPtr(2), utils.IntPtr(8), utils.IntPtr(0), utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(9), nil, nil, utils.IntPtr(3), utils.IntPtr(5)},
			2,
			8,
			6,
		},
		{
			"example_2",
			[]*int{utils.IntPtr(6), utils.IntPtr(2), utils.IntPtr(8), utils.IntPtr(0), utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(9), nil, nil, utils.IntPtr(3), utils.IntPtr(5)},
			2,
			4,
			2,
		},
		{
			"example_3",
			[]*int{utils.IntPtr(2), utils.IntPtr(1)},
			2,
			1,
			2,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.rootVals)
			p := utils.FindNode(root, tst.pVal)
			q := utils.FindNode(root, tst.qVal)
			want := utils.FindNode(root, tst.wantVal)

			got := lowestCommonAncestor(root, p, q)
			if got != want {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, want.Val)
			}

			got = lowestCommonAncestor2(root, p, q)
			if got != want {
				t.Errorf("lowestCommonAncestor2() = %v, want %v", got.Val, want.Val)
			}

			got = lowestCommonAncestor3(root, p, q)
			if got != want {
				t.Errorf("lowestCommonAncestor3() = %v, want %v", got.Val, want.Val)
			}
		})
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		rootVals []*int
		pVal     int
		qVal     int
		wantVal  int
	}{
		{
			"two_nodes_root_and_left",
			[]*int{utils.IntPtr(5), utils.IntPtr(3)},
			5,
			3,
			5,
		},
		{
			"two_nodes_root_and_right",
			[]*int{utils.IntPtr(5), nil, utils.IntPtr(7)},
			5,
			7,
			5,
		},
		{
			"both_in_left_subtree",
			[]*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), utils.IntPtr(3), utils.IntPtr(7), nil, nil, utils.IntPtr(1), utils.IntPtr(4)},
			1,
			7,
			5,
		},
		{
			"both_in_right_subtree",
			[]*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), nil, nil, utils.IntPtr(12), utils.IntPtr(20), utils.IntPtr(11), utils.IntPtr(13)},
			11,
			20,
			15,
		},
		{
			"p_is_parent_of_q",
			[]*int{utils.IntPtr(20), utils.IntPtr(10), utils.IntPtr(30), utils.IntPtr(5), utils.IntPtr(15)},
			10,
			15,
			10,
		},
		{
			"q_is_parent_of_p",
			[]*int{utils.IntPtr(20), utils.IntPtr(10), utils.IntPtr(30), utils.IntPtr(5), utils.IntPtr(15)},
			15,
			10,
			10,
		},
		{
			"deep_left_and_deep_right",
			[]*int{utils.IntPtr(100), utils.IntPtr(50), utils.IntPtr(200), utils.IntPtr(25), utils.IntPtr(75), utils.IntPtr(150), utils.IntPtr(250), utils.IntPtr(10), utils.IntPtr(30), utils.IntPtr(70), utils.IntPtr(80)},
			10,
			80,
			50,
		},
		{
			"negative_values",
			[]*int{utils.IntPtr(-10), utils.IntPtr(-20), utils.IntPtr(-5), utils.IntPtr(-30), nil, utils.IntPtr(-7), utils.IntPtr(-3)},
			-30,
			-3,
			-10,
		},
		{
			"mixed_positive_negative",
			[]*int{utils.IntPtr(0), utils.IntPtr(-100), utils.IntPtr(100), utils.IntPtr(-200), utils.IntPtr(-50), utils.IntPtr(50), utils.IntPtr(200)},
			-200,
			200,
			0,
		},
		{
			"large_values_near_boundary",
			[]*int{utils.IntPtr(1000000000), utils.IntPtr(-1000000000)},
			1000000000,
			-1000000000,
			1000000000,
		},
		{
			"chain_left",
			[]*int{utils.IntPtr(10), utils.IntPtr(9), nil, utils.IntPtr(8), nil, utils.IntPtr(7), nil, utils.IntPtr(6)},
			10,
			6,
			10,
		},
		{
			"chain_right",
			[]*int{utils.IntPtr(1), nil, utils.IntPtr(2), nil, utils.IntPtr(3), nil, utils.IntPtr(4), nil, utils.IntPtr(5)},
			1,
			5,
			1,
		},
		{
			"adjacent_nodes_in_chain",
			[]*int{utils.IntPtr(10), utils.IntPtr(9), nil, utils.IntPtr(8), nil, utils.IntPtr(7), nil, utils.IntPtr(6)},
			8,
			7,
			8,
		},
		{
			"balanced_tree_middle_nodes",
			[]*int{utils.IntPtr(50), utils.IntPtr(25), utils.IntPtr(75), utils.IntPtr(12), utils.IntPtr(37), utils.IntPtr(62), utils.IntPtr(87), utils.IntPtr(6), utils.IntPtr(18), utils.IntPtr(31), utils.IntPtr(43), utils.IntPtr(56), utils.IntPtr(68), utils.IntPtr(81), utils.IntPtr(93)},
			18,
			43,
			25,
		},
		{
			"p_is_root_q_is_far_left",
			[]*int{utils.IntPtr(100), utils.IntPtr(50), utils.IntPtr(200), utils.IntPtr(25), utils.IntPtr(75), nil, nil, utils.IntPtr(10), utils.IntPtr(30)},
			100,
			10,
			100,
		},
		{
			"q_is_root_p_is_far_right",
			[]*int{utils.IntPtr(100), utils.IntPtr(50), utils.IntPtr(200), nil, nil, utils.IntPtr(150), utils.IntPtr(300), nil, nil, nil, utils.IntPtr(400)},
			400,
			100,
			100,
		},
		{
			"nodes_at_same_depth_different_subtrees",
			[]*int{utils.IntPtr(10), utils.IntPtr(5), utils.IntPtr(15), utils.IntPtr(3), utils.IntPtr(7), utils.IntPtr(13), utils.IntPtr(17)},
			3,
			17,
			10,
		},
		{
			"one_node_is_direct_child_of_other_deep",
			[]*int{utils.IntPtr(100), utils.IntPtr(50), utils.IntPtr(150), utils.IntPtr(25), utils.IntPtr(75), utils.IntPtr(125), utils.IntPtr(175), nil, nil, utils.IntPtr(60), utils.IntPtr(80)},
			75,
			80,
			75,
		},
		{
			"swapped_p_and_q",
			[]*int{utils.IntPtr(6), utils.IntPtr(2), utils.IntPtr(8), utils.IntPtr(0), utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(9), nil, nil, utils.IntPtr(3), utils.IntPtr(5)},
			8,
			2,
			6,
		},
		{
			"swapped_p_and_q_same_subtree",
			[]*int{utils.IntPtr(6), utils.IntPtr(2), utils.IntPtr(8), utils.IntPtr(0), utils.IntPtr(4), utils.IntPtr(7), utils.IntPtr(9), nil, nil, utils.IntPtr(3), utils.IntPtr(5)},
			4,
			2,
			2,
		},
		{
			"min_node_count_2",
			[]*int{utils.IntPtr(2), utils.IntPtr(1)},
			2,
			1,
			2,
		},
		{
			"zig_zag_path",
			[]*int{utils.IntPtr(50), utils.IntPtr(30), utils.IntPtr(70), nil, utils.IntPtr(40), nil, utils.IntPtr(80), nil, utils.IntPtr(45)},
			45,
			80,
			50,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			root := utils.SliceToTreeNode(tst.rootVals)
			p := utils.FindNode(root, tst.pVal)
			q := utils.FindNode(root, tst.qVal)
			want := utils.FindNode(root, tst.wantVal)

			got := lowestCommonAncestor(root, p, q)
			if got != want {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, want.Val)
			}

			got = lowestCommonAncestor2(root, p, q)
			if got != want {
				t.Errorf("lowestCommonAncestor2() = %v, want %v", got.Val, want.Val)
			}

			got = lowestCommonAncestor3(root, p, q)
			if got != want {
				t.Errorf("lowestCommonAncestor3() = %v, want %v", got.Val, want.Val)
			}
		})
	}
}
