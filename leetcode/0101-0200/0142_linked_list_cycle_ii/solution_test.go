package leetcode0142

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

// getNodeAtIndex returns the node at the given index in the linked list.
// If the list is shorter than index+1 nodes, returns nil.
func getNodeAtIndex(head *utils.ListNode, index int) *utils.ListNode {
	if index < 0 {
		return nil
	}
	curr := head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}
	return curr
}

func TestDetectCycleExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		pos  int
	}{
		{
			"example_1",
			[]int{3, 2, 0, -4},
			1,
		},
		{
			"example_2",
			[]int{1, 2},
			0,
		},
		{
			"example_3",
			[]int{1},
			-1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToCyclicList(tst.nums, tst.pos)
			want := getNodeAtIndex(head, tst.pos)

			got := detectCycle(head)
			if got != want {
				t.Errorf("detectCycle() = %v, want %v", got, want)
			}

			got2 := detectCycle2(head)
			if got2 != want {
				t.Errorf("detectCycle2() = %v, want %v", got2, want)
			}
		})
	}
}

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		pos  int
	}{
		{"empty_list", []int{}, -1},
		{"single_node_no_cycle", []int{1}, -1},
		{"single_node_with_cycle", []int{1}, 0},
		{"two_nodes_no_cycle", []int{1, 2}, -1},
		{"two_nodes_cycle_at_0", []int{1, 2}, 0},
		{"two_nodes_cycle_at_1", []int{1, 2}, 1},
		{"three_nodes_no_cycle", []int{1, 2, 3}, -1},
		{"three_nodes_cycle_at_0", []int{1, 2, 3}, 0},
		{"three_nodes_cycle_at_1", []int{1, 2, 3}, 1},
		{"three_nodes_cycle_at_2", []int{1, 2, 3}, 2},
		{"long_list_no_cycle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1},
		{"long_list_cycle_at_start", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0},
		{"long_list_cycle_at_middle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4},
		{"long_list_cycle_at_end", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9},
		{"all_same_values_no_cycle", []int{5, 5, 5, 5, 5}, -1},
		{"all_same_values_cycle_at_middle", []int{5, 5, 5, 5, 5}, 2},
		{"negative_values_no_cycle", []int{-1, -2, -3, -4, -5}, -1},
		{"negative_values_cycle_at_1", []int{-1, -2, -3, -4, -5}, 1},
		{"mixed_values_no_cycle", []int{-100, 0, 100, -50, 50}, -1},
		{"mixed_values_cycle_at_3", []int{-100, 0, 100, -50, 50}, 3},
		{"cycle_length_1", []int{1, 2, 3}, 2},
		{"cycle_length_2", []int{1, 2, 3, 4}, 2},
		{"cycle_length_5", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{"list_with_max_constraint", make([]int, 100), -1},
		{"list_with_max_constraint_and_cycle", make([]int, 100), 50},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToCyclicList(tst.nums, tst.pos)
			want := getNodeAtIndex(head, tst.pos)

			got := detectCycle(head)
			if got != want {
				t.Errorf("detectCycle() = %v, want %v", got, want)
			}

			got2 := detectCycle2(head)
			if got2 != want {
				t.Errorf("detectCycle2() = %v, want %v", got2, want)
			}
		})
	}
}
