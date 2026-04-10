// Package leetcode0141 solves LeetCode 141. Linked List Cycle
package leetcode0141

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestHasCycleExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		pos  int
		want bool
	}{
		{
			"example_1",
			[]int{3, 2, 0, -4},
			1,
			true,
		},
		{
			"example_2",
			[]int{1, 2},
			0,
			true,
		},
		{
			"example_3",
			[]int{1},
			-1,
			false,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToCyclicList(tst.nums, tst.pos)
			got := hasCycle(head)
			if got != tst.want {
				t.Errorf("hasCycle() = %v, want %v", got, tst.want)
			}
		})
	}
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		pos  int
		want bool
	}{
		{"empty_list", []int{}, -1, false},
		{"single_node_no_cycle", []int{1}, -1, false},
		{"single_node_with_cycle", []int{1}, 0, true},
		{"two_nodes_no_cycle", []int{1, 2}, -1, false},
		{"two_nodes_cycle_at_0", []int{1, 2}, 0, true},
		{"two_nodes_cycle_at_1", []int{1, 2}, 1, true},
		{"three_nodes_no_cycle", []int{1, 2, 3}, -1, false},
		{"three_nodes_cycle_at_0", []int{1, 2, 3}, 0, true},
		{"three_nodes_cycle_at_1", []int{1, 2, 3}, 1, true},
		{"three_nodes_cycle_at_2", []int{1, 2, 3}, 2, true},
		{"cycle_at_middle", []int{1, 2, 3, 4, 5}, 2, true},
		{"cycle_at_tail", []int{1, 2, 3, 4, 5}, 4, true},
		{"all_same_values_no_cycle", []int{5, 5, 5, 5}, -1, false},
		{"all_same_values_with_cycle", []int{5, 5, 5, 5}, 1, true},
		{"negative_values_no_cycle", []int{-1, -2, -3, -4}, -1, false},
		{"negative_values_with_cycle", []int{-1, -2, -3, -4}, 1, true},
		{"mixed_values_no_cycle", []int{-100, 0, 50, -50, 100}, -1, false},
		{"mixed_values_with_cycle", []int{-100, 0, 50, -50, 100}, 3, true},
		{"long_list_no_cycle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1, false},
		{"long_list_cycle_at_start", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, true},
		{"long_list_cycle_at_middle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, true},
		{"long_list_cycle_at_end", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, true},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToCyclicList(tst.nums, tst.pos)
			got := hasCycle(head)
			if got != tst.want {
				t.Errorf("hasCycle() = %v, want %v", got, tst.want)
			}
		})
	}
}
