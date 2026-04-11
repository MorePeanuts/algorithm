package leetcode0019

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestRemoveNthFromEndExamples(t *testing.T) {
	tests := []struct {
		name string
		head []int
		n    int
		want []int
	}{
		{
			"example_1",
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{1, 2, 3, 5},
		},
		{
			"example_2",
			[]int{1},
			1,
			[]int{},
		},
		{
			"example_3",
			[]int{1, 2},
			1,
			[]int{1},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToList(tst.head)
			got := removeNthFromEnd(head, tst.n)
			gotSlice := utils.ListToSlice(got)
			if !utils.EqualLists(utils.SliceToList(gotSlice), utils.SliceToList(tst.want)) {
				t.Errorf("removeNthFromEnd(%v, %d) = %v, want %v", tst.head, tst.n, gotSlice, tst.want)
			}
		})
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		head []int
		n    int
		want []int
	}{
		{"single_element", []int{1}, 1, []int{}},
		{"two_elements_remove_first", []int{1, 2}, 2, []int{2}},
		{"two_elements_remove_last", []int{1, 2}, 1, []int{1}},
		{"three_elements_remove_middle", []int{1, 2, 3}, 2, []int{1, 3}},
		{"three_elements_remove_first", []int{1, 2, 3}, 3, []int{2, 3}},
		{"three_elements_remove_last", []int{1, 2, 3}, 1, []int{1, 2}},
		{"four_elements_remove_second_from_end", []int{1, 2, 3, 4}, 2, []int{1, 2, 4}},
		{"four_elements_remove_third_from_end", []int{1, 2, 3, 4}, 3, []int{1, 3, 4}},
		{"all_same_values", []int{5, 5, 5, 5}, 2, []int{5, 5, 5}},
		{"increasing_sequence", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, []int{1, 2, 3, 4, 5, 7, 8, 9, 10}},
		{"remove_first_of_long_list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"remove_last_of_long_list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"remove_middle_of_long_list", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 8, []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15}},
		{"zero_values", []int{0, 0, 0, 0, 0}, 3, []int{0, 0, 0, 0}},
		{"alternating_values", []int{1, 0, 1, 0, 1, 0}, 4, []int{1, 0, 0, 1, 0}},
		{"max_length_remove_middle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, 15, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}},
		{"max_length_remove_first", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, 30, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}},
		{"max_length_remove_last", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, 1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},
		{"descending_sequence", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 6, []int{10, 9, 8, 7, 5, 4, 3, 2, 1}},
		{"remove_second_from_single_element", []int{100}, 1, []int{}},
		{"all_max_values", []int{100, 100, 100, 100, 100}, 1, []int{100, 100, 100, 100}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToList(tst.head)
			got := removeNthFromEnd(head, tst.n)
			gotSlice := utils.ListToSlice(got)
			if !utils.EqualLists(utils.SliceToList(gotSlice), utils.SliceToList(tst.want)) {
				t.Errorf("removeNthFromEnd(%v, %d) = %v, want %v", tst.head, tst.n, gotSlice, tst.want)
			}
		})
	}
}
