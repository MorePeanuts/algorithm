package leetcode0015

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestThreeSumExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			"example_1",
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			"example_2",
			[]int{0, 1, 1},
			[][]int{},
		},
		{
			"example_3",
			[]int{0, 0, 0},
			[][]int{{0, 0, 0}},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := threeSum(tst.nums)
			if !utils.MatchTwo2dIntSlice(got, tst.want) {
				t.Errorf("threeSum(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		// Minimum length (3 elements)
		{"min_length_no_result", []int{1, 2, 3}, [][]int{}},
		{"min_length_with_result", []int{-1, 0, 1}, [][]int{{-1, 0, 1}}},
		{"min_length_all_zeros", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"min_length_all_positive", []int{1, 1, 1}, [][]int{}},
		{"min_length_all_negative", []int{-1, -1, -1}, [][]int{}},

		// All same elements
		{"all_same_zero", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"all_same_positive", []int{5, 5, 5, 5}, [][]int{}},
		{"all_same_negative", []int{-3, -3, -3, -3}, [][]int{}},

		// Two zeros with other elements
		{"two_zeros_with_zero", []int{0, 0, 0, 1}, [][]int{{0, 0, 0}}},
		{"two_zeros_no_match", []int{0, 0, 1, 2}, [][]int{}},

		// No valid triplet
		{"no_triplet_all_positive", []int{1, 2, 3, 4, 5}, [][]int{}},
		{"no_triplet_all_negative", []int{-5, -4, -3, -2, -1}, [][]int{}},
		{"no_triplet_mixed", []int{-4, -3, 1, 2}, [][]int{{-3, 1, 2}}},

		// Single triplet
		{"single_triplet_simple", []int{-2, 0, 2}, [][]int{{-2, 0, 2}}},
		{"single_triplet_with_extras", []int{-2, 0, 2, 5, 6}, [][]int{{-2, 0, 2}}},
		{"single_triplet_negative_heavy", []int{-5, -3, 8, 10}, [][]int{{-5, -3, 8}}},

		// Multiple triplets
		{"multiple_triplets_1", []int{-2, -1, 0, 1, 2, 3}, [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},
		{"multiple_triplets_2", []int{-4, -2, -1, 0, 1, 2, 3, 4}, [][]int{{-4, 0, 4}, {-4, 1, 3}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},

		// Duplicate handling
		{"duplicates_in_input_1", []int{-1, -1, 0, 1, 1}, [][]int{{-1, 0, 1}}},
		{"duplicates_in_input_2", []int{-2, -2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{"duplicates_multiple_results", []int{-2, -1, -1, 0, 1, 2, 2}, [][]int{{-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}},
		{"many_duplicates", []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}, [][]int{{-1, 0, 1}, {0, 0, 0}}},

		// Large values
		{"large_positive_values", []int{100000, -50000, -50000}, [][]int{{-50000, -50000, 100000}}},
		{"large_negative_values", []int{-100000, 50000, 50000}, [][]int{{-100000, 50000, 50000}}},
		{"large_mixed_values", []int{-100000, 0, 100000}, [][]int{{-100000, 0, 100000}}},

		// Zero sum with different combinations
		{"zero_sum_neg_neg_pos", []int{-3, -2, 5}, [][]int{{-3, -2, 5}}},
		{"zero_sum_neg_pos_pos", []int{-5, 2, 3}, [][]int{{-5, 2, 3}}},

		// Edge cases with zeros
		{"zeros_at_start", []int{0, 0, 0, -1, 1}, [][]int{{-1, 0, 1}, {0, 0, 0}}},
		{"zeros_at_end", []int{-1, 1, 0, 0, 0}, [][]int{{-1, 0, 1}, {0, 0, 0}}},
		{"zeros_in_middle", []int{-1, 0, 0, 0, 1}, [][]int{{-1, 0, 1}, {0, 0, 0}}},

		// Sorted vs unsorted input
		{"sorted_ascending", []int{-3, -2, -1, 0, 1, 2, 3}, [][]int{{-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},
		{"sorted_descending", []int{3, 2, 1, 0, -1, -2, -3}, [][]int{{-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},
		{"random_order", []int{1, -2, 3, 0, -1, 2, -3}, [][]int{{-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},

		// Specific patterns
		{"arithmetic_sequence", []int{-6, -4, -2, 0, 2, 4, 6}, [][]int{{-6, 0, 6}, {-6, 2, 4}, {-4, -2, 6}, {-4, 0, 4}, {-2, 0, 2}}},
		{"consecutive_integers", []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, [][]int{
			{-5, 0, 5},
			{-5, 1, 4},
			{-5, 2, 3},
			{-4, -1, 5},
			{-4, 0, 4},
			{-4, 1, 3},
			{-3, -2, 5},
			{-3, -1, 4},
			{-3, 0, 3},
			{-3, 1, 2},
			{-2, -1, 3},
			{-2, 0, 2},
			{-1, 0, 1},
		}},

		// More complex cases
		{"complex_case_1", []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}, [][]int{
			{-4, 0, 4},
			{-4, 1, 3},
			{-3, -1, 4},
			{-3, 0, 3},
			{-3, 1, 2},
			{-2, -1, 3},
			{-2, 0, 2},
			{-1, -1, 2},
			{-1, 0, 1},
		}},

		// Boundary values mixed
		{"boundary_mix_1", []int{-100000, 50000, 50000, 0}, [][]int{{-100000, 50000, 50000}}},
		{"boundary_mix_2", []int{100000, -50000, -50000, 0}, [][]int{{-50000, -50000, 100000}}},

		// Four elements
		{"four_elements_one_result", []int{-1, 0, 1, 2}, [][]int{{-1, 0, 1}}},
		{"four_elements_no_triplet", []int{-2, -1, 1, 2}, [][]int{}},
		{"four_elements_no_result", []int{1, 2, 3, 4}, [][]int{}},
		{"four_zeros", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},

		// Five elements
		{"five_elements_multiple", []int{-2, -1, 0, 1, 2}, [][]int{{-2, 0, 2}, {-1, 0, 1}}},

		// Special patterns
		{"alternating_signs", []int{-1, 1, -2, 2, -3, 3}, [][]int{{-3, 1, 2}, {-2, -1, 3}}},
		{"with_duplicates_pattern", []int{-4, -4, -2, -2, 0, 2, 2, 4, 4}, [][]int{{-4, 0, 4}, {-4, 2, 2}, {-2, 0, 2}, {-2, -2, 4}}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := threeSum(tst.nums)
			if !utils.MatchTwo2dIntSlice(got, tst.want) {
				t.Errorf("threeSum(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}
