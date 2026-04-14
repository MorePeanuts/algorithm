package leetcode0023

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestMergeKListsExamples(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{
			"example_1",
			[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			"example_2",
			[][]int{},
			[]int{},
		},
		{
			"example_3",
			[][]int{{}},
			[]int{},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			// Convert input slices to linked lists
			inputLists := make([]*utils.ListNode, len(tst.lists))
			for i, list := range tst.lists {
				inputLists[i] = utils.SliceToList(list)
			}

			got := mergeKLists(inputLists)
			gotSlice := utils.ListToSlice(got)
			if len(gotSlice) != len(tst.want) {
				t.Errorf("mergeKLists() = %v, want %v", gotSlice, tst.want)
				return
			}
			for i := range gotSlice {
				if gotSlice[i] != tst.want[i] {
					t.Errorf("mergeKLists() = %v, want %v", gotSlice, tst.want)
					return
				}
			}
		})
	}
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{"single_empty_list", [][]int{{}}, []int{}},
		{"single_list_with_one_element", [][]int{{5}}, []int{5}},
		{"single_list_with_multiple_elements", [][]int{{1, 2, 3, 4, 5}}, []int{1, 2, 3, 4, 5}},
		{"two_empty_lists", [][]int{{}, {}}, []int{}},
		{"two_lists_one_empty", [][]int{{1, 2, 3}, {}}, []int{1, 2, 3}},
		{"two_lists_one_empty_reversed", [][]int{{}, {1, 2, 3}}, []int{1, 2, 3}},
		{"two_sorted_lists_no_overlap", [][]int{{1, 2, 3}, {4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"two_sorted_lists_with_overlap", [][]int{{1, 3, 5}, {2, 4, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"two_sorted_lists_duplicate_elements", [][]int{{1, 1, 2}, {1, 2, 2}}, []int{1, 1, 1, 2, 2, 2}},
		{"three_lists_all_empty", [][]int{{}, {}, {}}, []int{}},
		{"three_lists_two_empty", [][]int{{1, 2, 3}, {}, {}}, []int{1, 2, 3}},
		{"three_lists_with_negative_numbers", [][]int{{-5, -3, 0}, {-4, -2, 1}, {-1, 2, 3}}, []int{-5, -4, -3, -2, -1, 0, 1, 2, 3}},
		{"lists_with_single_element_each", [][]int{{5}, {3}, {1}, {2}, {4}}, []int{1, 2, 3, 4, 5}},
		{"lists_with_same_elements", [][]int{{2, 2, 2}, {2, 2}, {2}}, []int{2, 2, 2, 2, 2, 2}},
		{"lists_with_descending_single_elements", [][]int{{10}, {9}, {8}, {7}, {6}, {5}, {4}, {3}, {2}, {1}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"four_lists_mixed_lengths", [][]int{{1, 4, 7, 10}, {2, 5, 8}, {3, 6, 9, 11, 12}, {}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
		{"lists_with_max_min_values", [][]int{{-10000, 10000}, {-9999, 9999}, {0}}, []int{-10000, -9999, 0, 9999, 10000}},
		{"alternating_high_low", [][]int{{1, 100, 101}, {2, 99, 102}, {3, 98, 103}}, []int{1, 2, 3, 98, 99, 100, 101, 102, 103}},
		{"one_long_list_and_many_short", [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {11}, {12}, {13}, {14}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
		{"all_lists_start_with_same_value", [][]int{{1, 5, 9}, {1, 6, 10}, {1, 7, 11}, {1, 8, 12}}, []int{1, 1, 1, 1, 5, 6, 7, 8, 9, 10, 11, 12}},
		{"k_equals_1", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"k_equals_2", [][]int{{1, 3}, {2, 4}}, []int{1, 2, 3, 4}},
		{"k_equals_5", [][]int{{1}, {2}, {3}, {4}, {5}}, []int{1, 2, 3, 4, 5}},
		{"lists_with_zero_and_negatives", [][]int{{-3, -1, 0}, {-2, 0, 2}, {-1, 1, 3}}, []int{-3, -2, -1, -1, 0, 0, 1, 2, 3}},
		{"longest_possible_single_list", [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			// Convert input slices to linked lists
			inputLists := make([]*utils.ListNode, len(tst.lists))
			for i, list := range tst.lists {
				inputLists[i] = utils.SliceToList(list)
			}

			got := mergeKLists(inputLists)
			gotSlice := utils.ListToSlice(got)
			if len(gotSlice) != len(tst.want) {
				t.Errorf("mergeKLists() = %v, want %v", gotSlice, tst.want)
				return
			}
			for i := range gotSlice {
				if gotSlice[i] != tst.want[i] {
					t.Errorf("mergeKLists() = %v, want %v", gotSlice, tst.want)
					return
				}
			}
		})
	}
}
