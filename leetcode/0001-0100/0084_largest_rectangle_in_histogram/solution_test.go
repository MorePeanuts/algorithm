package leetcode0084

import "testing"

func TestLargestRectangleAreaExamples(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			"example_1",
			[]int{2, 1, 5, 6, 2, 3},
			10,
		},
		{
			"example_2",
			[]int{2, 4},
			4,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := largestRectangleArea(tst.heights)
			if got != tst.want {
				t.Errorf("largestRectangleArea(%v) = %v, want %v", tst.heights, got, tst.want)
			}
		})
	}
}

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		// Single element cases
		{"single_element_zero", []int{0}, 0},
		{"single_element_one", []int{1}, 1},
		{"single_element_max", []int{10000}, 10000},

		// Two elements cases
		{"two_equal", []int{3, 3}, 6},
		{"two_ascending", []int{1, 2}, 2},
		{"two_descending", []int{2, 1}, 2},
		{"two_with_zero", []int{0, 5}, 5},
		{"two_both_zero", []int{0, 0}, 0},

		// Three elements cases
		{"three_ascending", []int{1, 2, 3}, 4},
		{"three_descending", []int{3, 2, 1}, 4},
		{"three_peak", []int{1, 3, 1}, 3},
		{"three_valley", []int{3, 1, 3}, 3},
		{"three_equal", []int{2, 2, 2}, 6},

		// All same height
		{"all_same_small", []int{5, 5, 5, 5}, 20},
		{"all_same_large", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 10},
		{"all_zeros", []int{0, 0, 0, 0}, 0},

		// Ascending sequences
		{"ascending_sequence", []int{1, 2, 3, 4, 5}, 9},
		{"ascending_from_zero", []int{0, 1, 2, 3, 4}, 6},

		// Descending sequences
		{"descending_sequence", []int{5, 4, 3, 2, 1}, 9},
		{"descending_to_zero", []int{4, 3, 2, 1, 0}, 6},

		// Peak patterns
		{"single_peak", []int{1, 2, 3, 2, 1}, 6},
		{"sharp_peak", []int{1, 5, 1}, 5},
		{"wide_peak", []int{1, 2, 3, 3, 3, 2, 1}, 10},

		// Valley patterns
		{"single_valley", []int{3, 2, 1, 2, 3}, 5},
		{"deep_valley", []int{5, 1, 5}, 5},
		{"wide_valley", []int{4, 3, 2, 2, 2, 3, 4}, 14},

		// Multiple peaks
		{"two_peaks", []int{1, 3, 1, 3, 1}, 5},
		{"unequal_peaks", []int{1, 5, 1, 3, 1}, 5},

		// Plateau patterns
		{"left_plateau", []int{3, 3, 3, 1, 2}, 9},
		{"right_plateau", []int{2, 1, 3, 3, 3}, 9},
		{"middle_plateau", []int{1, 3, 3, 3, 1}, 9},

		// Step patterns
		{"step_up", []int{1, 1, 2, 2, 3, 3}, 8},
		{"step_down", []int{3, 3, 2, 2, 1, 1}, 8},

		// Zigzag patterns
		{"zigzag", []int{1, 3, 2, 4, 3, 5}, 10},
		{"reverse_zigzag", []int{5, 3, 4, 2, 3, 1}, 10},

		// Edge cases with zeros
		{"zeros_at_ends", []int{0, 3, 3, 3, 0}, 9},
		{"zeros_in_middle", []int{3, 0, 3}, 3},
		{"alternating_zeros", []int{1, 0, 2, 0, 3}, 3},

		// Large values
		{"max_height_single", []int{10000}, 10000},
		{"max_heights", []int{10000, 10000, 10000}, 30000},
		{"mixed_with_max", []int{1, 10000, 1}, 10000},

		// Classic histogram cases
		{"classic_1", []int{2, 1, 2}, 3},
		{"classic_2", []int{4, 2, 0, 3, 2, 5}, 6},
		{"classic_3", []int{1, 2, 3, 4, 5, 3, 3, 2}, 15},

		// Longer sequences
		{"long_ascending", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 30},
		{"long_descending", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 30},
		{"long_uniform", []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}, 30},
		{"long_alternating", []int{1, 5, 1, 5, 1, 5, 1, 5, 1, 5}, 10},

		// Complex patterns
		{"mountain_range", []int{1, 3, 2, 4, 3, 5, 4, 6, 5, 4}, 21},
		{"staircase_up_down", []int{1, 2, 3, 4, 5, 4, 3, 2, 1}, 15},
		{"double_peak_valley", []int{1, 4, 2, 4, 1}, 6},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := largestRectangleArea(tst.heights)
			if got != tst.want {
				t.Errorf("largestRectangleArea(%v) = %v, want %v", tst.heights, got, tst.want)
			}
		})
	}
}
