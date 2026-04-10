package leetcode0021

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestMergeTwoListsExamples(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{
			"example_1",
			[]int{1, 2, 4},
			[]int{1, 3, 4},
			[]int{1, 1, 2, 3, 4, 4},
		},
		{
			"example_2",
			[]int{},
			[]int{},
			[]int{},
		},
		{
			"example_3",
			[]int{},
			[]int{0},
			[]int{0},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			l1 := utils.SliceToList(tst.list1)
			l2 := utils.SliceToList(tst.list2)
			want := utils.SliceToList(tst.want)

			got := mergeTwoLists(l1, l2)
			if !utils.EqualLists(got, want) {
				t.Errorf("mergeTwoLists(%v, %v) = %v, want %v", tst.list1, tst.list2, utils.ListToSlice(got), tst.want)
			}
		})
	}
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{"single_element_both", []int{1}, []int{2}, []int{1, 2}},
		{"single_element_first", []int{5}, []int{}, []int{5}},
		{"single_element_second", []int{}, []int{5}, []int{5}},
		{"all_same_elements", []int{2, 2, 2}, []int{2, 2}, []int{2, 2, 2, 2, 2}},
		{"first_list_all_smaller", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"second_list_all_smaller", []int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
		{"interleaved", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"duplicates_at_boundary", []int{1, 4, 5}, []int{1, 2, 6}, []int{1, 1, 2, 4, 5, 6}},
		{"negative_values", []int{-5, -3, 0}, []int{-4, -2, 1}, []int{-5, -4, -3, -2, 0, 1}},
		{"all_negative", []int{-10, -5, -1}, []int{-8, -3}, []int{-10, -8, -5, -3, -1}},
		{"max_value_100", []int{100}, []int{100}, []int{100, 100}},
		{"min_value_neg100", []int{-100}, []int{-100}, []int{-100, -100}},
		{"first_longer_than_second", []int{1, 2, 3, 4, 5}, []int{6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"second_longer_than_first", []int{1, 2}, []int{3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"alternating_duplicates", []int{1, 1, 3, 3}, []int{2, 2, 4, 4}, []int{1, 1, 2, 2, 3, 3, 4, 4}},
		{"one_element_in_middle", []int{1, 5}, []int{3}, []int{1, 3, 5}},
		{"strictly_increasing", []int{1, 4, 7}, []int{2, 5, 8}, []int{1, 2, 4, 5, 7, 8}},
		{"plateaus", []int{2, 2, 3, 3, 3}, []int{1, 1, 1, 2, 2}, []int{1, 1, 1, 2, 2, 2, 2, 3, 3, 3}},
		{"max_length_50_each", makeRange(1, 50), makeRange(51, 100), makeRange(1, 100)},
		{"max_length_interleaved", makeOddRange(1, 99), makeEvenRange(2, 100), makeRange(1, 100)},
		{"descending_start", []int{10, 20, 30}, []int{5, 15, 25, 35}, []int{5, 10, 15, 20, 25, 30, 35}},
		{"single_vs_multiple", []int{5}, []int{1, 2, 3, 4, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"zero_in_middle", []int{-2, 0, 2}, []int{-1, 0, 1}, []int{-2, -1, 0, 0, 1, 2}},
		{"all_zero", []int{0, 0, 0}, []int{0, 0}, []int{0, 0, 0, 0, 0}},
		{"mixed_signs", []int{-5, 10, 15}, []int{-10, 5, 20}, []int{-10, -5, 5, 10, 15, 20}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			l1 := utils.SliceToList(tst.list1)
			l2 := utils.SliceToList(tst.list2)
			want := utils.SliceToList(tst.want)

			got := mergeTwoLists(l1, l2)
			if !utils.EqualLists(got, want) {
				t.Errorf("mergeTwoLists(%v, %v) = %v, want %v", tst.list1, tst.list2, utils.ListToSlice(got), tst.want)
			}
		})
	}
}

func makeRange(start, end int) []int {
	result := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		result = append(result, i)
	}
	return result
}

func makeOddRange(start, end int) []int {
	result := make([]int, 0)
	for i := start; i <= end; i++ {
		if i%2 == 1 {
			result = append(result, i)
		}
	}
	return result
}

func makeEvenRange(start, end int) []int {
	result := make([]int, 0)
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	return result
}
