package leetcode0239

import (
	"slices"
	"testing"
)

func TestMaxSlidingWindowExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			"example_1",
			[]int{1, 3, -1, -3, 5, 3, 6, 7},
			3,
			[]int{3, 3, 5, 5, 6, 7},
		},
		{
			"example_2",
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := maxSlidingWindow(tst.nums, tst.k)
			if !slices.Equal(got, tst.want) {
				t.Errorf("maxSlidingWindow(%v, %v) = %v, want %v", tst.nums, tst.k, got, tst.want)
			}
		})
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		// Edge cases
		{"single_element", []int{5}, 1, []int{5}},
		{"k_equals_length", []int{1, 2, 3, 4, 5}, 5, []int{5}},
		{"k_1", []int{10, 9, 8, 7, 6}, 1, []int{10, 9, 8, 7, 6}},
		{"k_2_small_array", []int{3, 1}, 2, []int{3}},

		// All same elements
		{"all_same_positive", []int{2, 2, 2, 2}, 2, []int{2, 2, 2}},
		{"all_same_negative", []int{-5, -5, -5}, 2, []int{-5, -5}},
		{"all_zero", []int{0, 0, 0, 0, 0}, 3, []int{0, 0, 0}},

		// Strictly increasing
		{"strictly_increasing_k_2", []int{1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5}},
		{"strictly_increasing_k_3", []int{1, 2, 3, 4, 5, 6}, 3, []int{3, 4, 5, 6}},
		{"strictly_increasing_k_4", []int{10, 20, 30, 40, 50}, 4, []int{40, 50}},

		// Strictly decreasing
		{"strictly_decreasing_k_2", []int{5, 4, 3, 2, 1}, 2, []int{5, 4, 3, 2}},
		{"strictly_decreasing_k_3", []int{6, 5, 4, 3, 2, 1}, 3, []int{6, 5, 4, 3}},
		{"strictly_decreasing_k_4", []int{50, 40, 30, 20, 10}, 4, []int{50, 40}},

		// Peak patterns
		{"peak_in_middle", []int{1, 3, 5, 4, 2}, 3, []int{5, 5, 5}},
		{"peak_at_start", []int{10, 5, 3, 4, 6}, 3, []int{10, 5, 6}},
		{"peak_at_end", []int{2, 4, 3, 5, 8}, 3, []int{4, 5, 8}},
		{"valley_pattern", []int{5, 4, 3, 2, 1, 2, 3, 4, 5}, 3, []int{5, 4, 3, 2, 3, 4, 5}},

		// Negative numbers
		{"all_negatives", []int{-1, -3, -2, -5, -4}, 2, []int{-1, -2, -2, -4}},
		{"mix_pos_neg_k_2", []int{-5, 3, -2, 4, -1}, 2, []int{3, 3, 4, 4}},
		{"mix_pos_neg_k_3", []int{-5, 3, -2, 4, -1, 6}, 3, []int{3, 4, 4, 6}},
		{"negatives_with_zeros", []int{-2, 0, -1, -3, 0}, 2, []int{0, 0, -1, 0}},

		// Duplicate maxima
		{"duplicate_maxima", []int{5, 3, 5, 2, 5}, 2, []int{5, 5, 5, 5}},
		{"duplicate_maxima_k_3", []int{7, 7, 6, 7, 7}, 3, []int{7, 7, 7}},
		{"max_moves_out_window", []int{1, 5, 3, 2, 4, 6}, 3, []int{5, 5, 4, 6}},

		// Alternating patterns
		{"alternating_high_low", []int{10, 1, 10, 1, 10, 1}, 2, []int{10, 10, 10, 10, 10}},
		{"alternating_low_high_k_3", []int{1, 10, 1, 10, 1, 10, 1}, 3, []int{10, 10, 10, 10, 10}},

		// Longer sequences
		{"length_10_k_3", []int{9, 5, 7, 3, 8, 2, 6, 4, 1, 10}, 3, []int{9, 7, 8, 8, 8, 6, 6, 10}},
		{"length_10_k_5", []int{1, 2, 3, 4, 5, 4, 3, 2, 1}, 5, []int{5, 5, 5, 5, 5}},

		// Edge values from constraints
		{"min_value", []int{-10000, -10000, -10000}, 2, []int{-10000, -10000}},
		{"max_value", []int{10000, 10000, 10000}, 2, []int{10000, 10000}},
		{"min_max_mix", []int{-10000, 10000, -10000, 10000}, 2, []int{10000, 10000, 10000}},

		// Window slides past maximum
		{"max_at_start_then_small", []int{10, 1, 2, 3, 4, 5}, 3, []int{10, 3, 4, 5}},
		{"max_just_leaves_window", []int{5, 4, 3, 6, 2, 1, 7}, 3, []int{5, 6, 6, 6, 7}},

		// Random patterns
		{"random_1", []int{3, 1, 4, 1, 5, 9, 2, 6}, 3, []int{4, 4, 5, 9, 9, 9}},
		{"random_2", []int{8, 5, 10, 7, 9, 4, 15, 12, 3}, 4, []int{10, 10, 10, 15, 15, 15}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := maxSlidingWindow(tst.nums, tst.k)
			if !slices.Equal(got, tst.want) {
				t.Errorf("maxSlidingWindow(%v, %v) = %v, want %v", tst.nums, tst.k, got, tst.want)
			}
		})
	}
}
