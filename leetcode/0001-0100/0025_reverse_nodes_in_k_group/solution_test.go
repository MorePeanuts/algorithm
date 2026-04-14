package leetcode0025

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestReverseKGroupExamples(t *testing.T) {
	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		{
			"example_1",
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{2, 1, 4, 3, 5},
		},
		{
			"example_2",
			[]int{1, 2, 3, 4, 5},
			3,
			[]int{3, 2, 1, 4, 5},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToList(tst.head)
			got := reverseKGroup(head, tst.k)
			gotSlice := utils.ListToSlice(got)
			if !utils.EqualLists(utils.SliceToList(gotSlice), utils.SliceToList(tst.want)) {
				t.Errorf("reverseKGroup(%v, %d) = %v, want %v", tst.head, tst.k, gotSlice, tst.want)
			}
		})
	}
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		{"single_node", []int{1}, 1, []int{1}},
		{"k_equals_length", []int{1, 2, 3}, 3, []int{3, 2, 1}},
		{"k_equals_1", []int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{"exact_multiple", []int{1, 2, 3, 4}, 2, []int{2, 1, 4, 3}},
		{"one_left_over", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{3, 2, 1, 6, 5, 4, 7}},
		{"two_left_over", []int{1, 2, 3, 4, 5, 6, 7, 8}, 3, []int{3, 2, 1, 6, 5, 4, 7, 8}},
		{"all_same_values", []int{5, 5, 5, 5, 5}, 2, []int{5, 5, 5, 5, 5}},
		{"descending_order", []int{5, 4, 3, 2, 1}, 2, []int{4, 5, 2, 3, 1}},
		{"ascending_order", []int{1, 2, 3, 4, 5, 6}, 3, []int{3, 2, 1, 6, 5, 4}},
		{"zero_values", []int{0, 0, 0, 0}, 2, []int{0, 0, 0, 0}},
		{"mixed_values", []int{0, 1, 0, 1, 0, 1}, 2, []int{1, 0, 1, 0, 1, 0}},
		{"k_equals_4_length_5", []int{1, 2, 3, 4, 5}, 4, []int{4, 3, 2, 1, 5}},
		{"k_equals_5_length_5", []int{1, 2, 3, 4, 5}, 5, []int{5, 4, 3, 2, 1}},
		{"two_nodes_k_2", []int{1, 2}, 2, []int{2, 1}},
		{"three_nodes_k_2", []int{1, 2, 3}, 2, []int{2, 1, 3}},
		{"max_value_nodes", []int{1000, 1000, 1000}, 2, []int{1000, 1000, 1000}},
		{"min_max_mixed", []int{0, 1000, 0, 1000}, 2, []int{1000, 0, 1000, 0}},
		{"six_nodes_k_3", []int{1, 2, 3, 4, 5, 6}, 3, []int{3, 2, 1, 6, 5, 4}},
		{"seven_nodes_k_2", []int{1, 2, 3, 4, 5, 6, 7}, 2, []int{2, 1, 4, 3, 6, 5, 7}},
		{"eight_nodes_k_4", []int{1, 2, 3, 4, 5, 6, 7, 8}, 4, []int{4, 3, 2, 1, 8, 7, 6, 5}},
		{"alternating_values", []int{1, 2, 1, 2, 1, 2, 1}, 2, []int{2, 1, 2, 1, 2, 1, 1}},
		{"large_k_small_list", []int{1, 2, 3}, 3, []int{3, 2, 1}},
		{"four_groups_of_two", []int{1, 2, 3, 4, 5, 6, 7, 8}, 2, []int{2, 1, 4, 3, 6, 5, 8, 7}},
		{"two_groups_of_three", []int{1, 2, 3, 4, 5, 6}, 3, []int{3, 2, 1, 6, 5, 4}},
		{"one_group_remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, []int{4, 3, 2, 1, 8, 7, 6, 5, 9, 10}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToList(tst.head)
			got := reverseKGroup(head, tst.k)
			gotSlice := utils.ListToSlice(got)
			if !utils.EqualLists(utils.SliceToList(gotSlice), utils.SliceToList(tst.want)) {
				t.Errorf("reverseKGroup(%v, %d) = %v, want %v", tst.head, tst.k, gotSlice, tst.want)
			}
		})
	}
}
