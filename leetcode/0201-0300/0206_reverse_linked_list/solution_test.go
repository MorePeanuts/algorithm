// Package leetcode0206 solves LeetCode 206. Reverse Linked List
package leetcode0206

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestReverseListExamples(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{
			"example_1",
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
		{
			"example_2",
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			"example_3",
			[]int{},
			[]int{},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToList(tst.head)
			got := reverseList(head)
			gotSlice := utils.ListToSlice(got)
			if len(gotSlice) != len(tst.want) {
				t.Errorf("reverseList(%v) = %v, want %v", tst.head, gotSlice, tst.want)
				return
			}
			for i := range gotSlice {
				if gotSlice[i] != tst.want[i] {
					t.Errorf("reverseList(%v) = %v, want %v", tst.head, gotSlice, tst.want)
					return
				}
			}
		})
	}
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{"single_node", []int{1}, []int{1}},
		{"two_nodes", []int{1, 2}, []int{2, 1}},
		{"three_nodes", []int{1, 2, 3}, []int{3, 2, 1}},
		{"all_same_values", []int{5, 5, 5}, []int{5, 5, 5}},
		{"alternating_values", []int{1, 2, 1, 2}, []int{2, 1, 2, 1}},
		{"negative_values", []int{-1, -2, -3}, []int{-3, -2, -1}},
		{"mixed_positive_negative", []int{-1, 2, -3, 4}, []int{4, -3, 2, -1}},
		{"zero_in_list", []int{0, 1, 0}, []int{0, 1, 0}},
		{"descending_order", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"single_max_value", []int{5000}, []int{5000}},
		{"single_min_value", []int{-5000}, []int{-5000}},
		{"max_and_min", []int{5000, -5000}, []int{-5000, 5000}},
		{"six_nodes", []int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
		{"seven_nodes", []int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1}},
		{"increasing_then_decreasing", []int{1, 3, 5, 3, 1}, []int{1, 3, 5, 3, 1}},
		{"all_zero", []int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{"sequence_with_gaps", []int{1, 3, 5, 7}, []int{7, 5, 3, 1}},
		{"palindrome_length_5", []int{1, 2, 3, 2, 1}, []int{1, 2, 3, 2, 1}},
		{"palindrome_length_4", []int{1, 2, 2, 1}, []int{1, 2, 2, 1}},
		{"duplicates_at_start", []int{1, 1, 2, 3}, []int{3, 2, 1, 1}},
		{"duplicates_at_end", []int{1, 2, 3, 3}, []int{3, 3, 2, 1}},
		{"alternating_signs", []int{1, -1, 1, -1, 1}, []int{1, -1, 1, -1, 1}},
		{"large_negative_sequence", []int{-100, -200, -300, -400}, []int{-400, -300, -200, -100}},
		{"mixed_large_values", []int{5000, -5000, 1000, -1000}, []int{-1000, 1000, -5000, 5000}},
		{"ten_nodes", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToList(tst.head)
			got := reverseList(head)
			gotSlice := utils.ListToSlice(got)
			if len(gotSlice) != len(tst.want) {
				t.Errorf("reverseList(%v) = %v, want %v", tst.head, gotSlice, tst.want)
				return
			}
			for i := range gotSlice {
				if gotSlice[i] != tst.want[i] {
					t.Errorf("reverseList(%v) = %v, want %v", tst.head, gotSlice, tst.want)
					return
				}
			}
		})
	}
}
