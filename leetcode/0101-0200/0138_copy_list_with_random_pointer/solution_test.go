package leetcode0138

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestCopyRandomListExamples(t *testing.T) {
	// Helper to create int pointers
	pint := func(i int) *int { return &i }

	tests := []struct {
		name  string
		input []utils.NodeWithRandomInput
	}{
		{
			"example_1",
			[]utils.NodeWithRandomInput{
				{Val: 7, RandomIndex: nil},
				{Val: 13, RandomIndex: pint(0)},
				{Val: 11, RandomIndex: pint(4)},
				{Val: 10, RandomIndex: pint(2)},
				{Val: 1, RandomIndex: pint(0)},
			},
		},
		{
			"example_2",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: pint(1)},
				{Val: 2, RandomIndex: pint(1)},
			},
		},
		{
			"example_3",
			[]utils.NodeWithRandomInput{
				{Val: 3, RandomIndex: nil},
				{Val: 3, RandomIndex: pint(0)},
				{Val: 3, RandomIndex: nil},
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			original := utils.SliceToNodeWithRandom(tst.input)

			// Test copyRandomList
			got := copyRandomList(original)
			if !utils.EqualNodeWithRandom(original, got) {
				t.Errorf("copyRandomList() did not produce a correct deep copy")
			}

			// Test copyRandomList2
			got = copyRandomList2(original)
			if !utils.EqualNodeWithRandom(original, got) {
				t.Errorf("copyRandomList2() did not produce a correct deep copy")
			}

			// Test copyRandomList3
			got = copyRandomList3(original)
			if !utils.EqualNodeWithRandom(original, got) {
				t.Errorf("copyRandomList3() did not produce a correct deep copy")
			}
		})
	}
}

func TestCopyRandomList(t *testing.T) {
	// Helper to create int pointers
	pint := func(i int) *int { return &i }

	tests := []struct {
		name  string
		input []utils.NodeWithRandomInput
	}{
		{
			"empty_list",
			nil,
		},
		{
			"single_node_nil_random",
			[]utils.NodeWithRandomInput{
				{Val: 5, RandomIndex: nil},
			},
		},
		{
			"single_node_self_random",
			[]utils.NodeWithRandomInput{
				{Val: 42, RandomIndex: pint(0)},
			},
		},
		{
			"two_nodes_forward_random",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: pint(1)},
				{Val: 2, RandomIndex: nil},
			},
		},
		{
			"two_nodes_backward_random",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: nil},
				{Val: 2, RandomIndex: pint(0)},
			},
		},
		{
			"two_nodes_circular_random",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: pint(1)},
				{Val: 2, RandomIndex: pint(0)},
			},
		},
		{
			"three_nodes_all_self_random",
			[]utils.NodeWithRandomInput{
				{Val: 10, RandomIndex: pint(0)},
				{Val: 20, RandomIndex: pint(1)},
				{Val: 30, RandomIndex: pint(2)},
			},
		},
		{
			"three_nodes_chain_random",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: pint(1)},
				{Val: 2, RandomIndex: pint(2)},
				{Val: 3, RandomIndex: nil},
			},
		},
		{
			"three_nodes_reverse_chain_random",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: nil},
				{Val: 2, RandomIndex: pint(0)},
				{Val: 3, RandomIndex: pint(1)},
			},
		},
		{
			"all_nodes_point_to_first",
			[]utils.NodeWithRandomInput{
				{Val: 100, RandomIndex: pint(0)},
				{Val: 200, RandomIndex: pint(0)},
				{Val: 300, RandomIndex: pint(0)},
				{Val: 400, RandomIndex: pint(0)},
			},
		},
		{
			"all_nodes_point_to_last",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: pint(4)},
				{Val: 2, RandomIndex: pint(4)},
				{Val: 3, RandomIndex: pint(4)},
				{Val: 4, RandomIndex: pint(4)},
				{Val: 5, RandomIndex: pint(4)},
			},
		},
		{
			"alternating_nil_random",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: nil},
				{Val: 2, RandomIndex: pint(0)},
				{Val: 3, RandomIndex: nil},
				{Val: 4, RandomIndex: pint(2)},
				{Val: 5, RandomIndex: nil},
				{Val: 6, RandomIndex: pint(4)},
			},
		},
		{
			"cross_pointers",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: pint(3)},
				{Val: 2, RandomIndex: pint(2)},
				{Val: 3, RandomIndex: pint(1)},
				{Val: 4, RandomIndex: pint(0)},
			},
		},
		{
			"negative_values",
			[]utils.NodeWithRandomInput{
				{Val: -1, RandomIndex: pint(2)},
				{Val: -5, RandomIndex: nil},
				{Val: -10, RandomIndex: pint(0)},
			},
		},
		{
			"mixed_positive_negative",
			[]utils.NodeWithRandomInput{
				{Val: -100, RandomIndex: pint(3)},
				{Val: 200, RandomIndex: pint(0)},
				{Val: -300, RandomIndex: pint(4)},
				{Val: 400, RandomIndex: pint(1)},
				{Val: -500, RandomIndex: pint(2)},
			},
		},
		{
			"all_same_values",
			[]utils.NodeWithRandomInput{
				{Val: 7, RandomIndex: pint(2)},
				{Val: 7, RandomIndex: pint(0)},
				{Val: 7, RandomIndex: pint(4)},
				{Val: 7, RandomIndex: pint(1)},
				{Val: 7, RandomIndex: pint(3)},
			},
		},
		{
			"long_chain_10_nodes",
			[]utils.NodeWithRandomInput{
				{Val: 0, RandomIndex: pint(9)},
				{Val: 1, RandomIndex: pint(8)},
				{Val: 2, RandomIndex: pint(7)},
				{Val: 3, RandomIndex: pint(6)},
				{Val: 4, RandomIndex: pint(5)},
				{Val: 5, RandomIndex: pint(4)},
				{Val: 6, RandomIndex: pint(3)},
				{Val: 7, RandomIndex: pint(2)},
				{Val: 8, RandomIndex: pint(1)},
				{Val: 9, RandomIndex: pint(0)},
			},
		},
		{
			"max_value_boundary",
			[]utils.NodeWithRandomInput{
				{Val: 10000, RandomIndex: pint(1)},
				{Val: -10000, RandomIndex: pint(0)},
			},
		},
		{
			"all_random_nil",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: nil},
				{Val: 2, RandomIndex: nil},
				{Val: 3, RandomIndex: nil},
				{Val: 4, RandomIndex: nil},
				{Val: 5, RandomIndex: nil},
			},
		},
		{
			"skip_one_random",
			[]utils.NodeWithRandomInput{
				{Val: 10, RandomIndex: pint(2)},
				{Val: 20, RandomIndex: pint(3)},
				{Val: 30, RandomIndex: pint(4)},
				{Val: 40, RandomIndex: pint(0)},
				{Val: 50, RandomIndex: pint(1)},
			},
		},
		{
			"consecutive_pairs",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: pint(1)},
				{Val: 2, RandomIndex: pint(0)},
				{Val: 3, RandomIndex: pint(3)},
				{Val: 4, RandomIndex: pint(2)},
				{Val: 5, RandomIndex: pint(5)},
				{Val: 6, RandomIndex: pint(4)},
			},
		},
		{
			"first_half_points_to_second_half",
			[]utils.NodeWithRandomInput{
				{Val: 1, RandomIndex: pint(5)},
				{Val: 2, RandomIndex: pint(6)},
				{Val: 3, RandomIndex: pint(7)},
				{Val: 4, RandomIndex: pint(8)},
				{Val: 5, RandomIndex: pint(9)},
				{Val: 6, RandomIndex: pint(0)},
				{Val: 7, RandomIndex: pint(1)},
				{Val: 8, RandomIndex: pint(2)},
				{Val: 9, RandomIndex: pint(3)},
				{Val: 10, RandomIndex: pint(4)},
			},
		},
		{
			"random_5_node_pattern",
			[]utils.NodeWithRandomInput{
				{Val: 100, RandomIndex: pint(3)},
				{Val: 200, RandomIndex: pint(3)},
				{Val: 300, RandomIndex: pint(3)},
				{Val: 400, RandomIndex: pint(3)},
				{Val: 500, RandomIndex: pint(3)},
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			original := utils.SliceToNodeWithRandom(tst.input)

			// Test copyRandomList
			got := copyRandomList(original)
			if !utils.EqualNodeWithRandom(original, got) {
				t.Errorf("copyRandomList() did not produce a correct deep copy for test: %s", tst.name)
			}

			// Test copyRandomList2
			got = copyRandomList2(original)
			if !utils.EqualNodeWithRandom(original, got) {
				t.Errorf("copyRandomList2() did not produce a correct deep copy for test: %s", tst.name)
			}

			// Test copyRandomList3
			got = copyRandomList3(original)
			if !utils.EqualNodeWithRandom(original, got) {
				t.Errorf("copyRandomList3() did not produce a correct deep copy for test: %s", tst.name)
			}
		})
	}
}
