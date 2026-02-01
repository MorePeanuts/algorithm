package leetcode0128

import "testing"

func TestLongestConsecutiveExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"example_1",
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
		{
			"example_2",
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			9,
		},
		{
			"example_3",
			[]int{1, 0, 1, 2},
			3,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := longestConsecutive(tst.nums)
			if got != tst.want {
				t.Errorf("longestConsecutive(%v) = %v, want %v", tst.nums, got, tst.want)
			}

			got = longestConsecutive2(tst.nums)
			if got != tst.want {
				t.Errorf("longestConsecutive(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// Empty and single element
		{"empty_array", []int{}, 0},
		{"single_element", []int{5}, 1},
		{"single_zero", []int{0}, 1},
		{"single_negative", []int{-1}, 1},

		// Two elements
		{"two_consecutive", []int{1, 2}, 2},
		{"two_consecutive_reverse", []int{2, 1}, 2},
		{"two_non_consecutive", []int{1, 3}, 1},
		{"two_same", []int{5, 5}, 1},

		// All same elements
		{"all_same_three", []int{7, 7, 7}, 1},
		{"all_same_five", []int{0, 0, 0, 0, 0}, 1},

		// Consecutive sequences
		{"consecutive_ascending", []int{1, 2, 3, 4, 5}, 5},
		{"consecutive_descending", []int{5, 4, 3, 2, 1}, 5},
		{"consecutive_shuffled", []int{3, 1, 4, 2, 5}, 5},

		// Multiple sequences
		{"two_sequences", []int{1, 2, 10, 11, 12}, 3},
		{"three_sequences", []int{1, 2, 10, 11, 20, 21, 22}, 3},
		{"equal_length_sequences", []int{1, 2, 3, 10, 11, 12}, 3},

		// With duplicates
		{"duplicates_in_sequence", []int{1, 2, 2, 3}, 3},
		{"many_duplicates", []int{1, 1, 1, 2, 2, 3, 3, 3}, 3},
		{"duplicates_scattered", []int{1, 3, 1, 2, 3, 2}, 3},

		// Negative numbers
		{"all_negative", []int{-3, -2, -1}, 3},
		{"negative_to_positive", []int{-2, -1, 0, 1, 2}, 5},
		{"negative_sequence", []int{-5, -4, -3, -2, -1}, 5},
		{"mixed_negative_positive", []int{-1, 0, 1, 5, 6, 7, 8}, 4},

		// Crossing zero
		{"cross_zero", []int{-1, 0, 1}, 3},
		{"cross_zero_longer", []int{-3, -2, -1, 0, 1, 2, 3}, 7},

		// Large gaps
		{"large_gap", []int{1, 100, 200, 300}, 1},
		{"sequence_with_outliers", []int{1, 2, 3, 1000, 2000}, 3},

		// Edge values
		{"large_positive", []int{999999999, 1000000000}, 2},
		{"large_negative", []int{-1000000000, -999999999}, 2},
		{"extreme_values", []int{-1000000000, 0, 1000000000}, 1},

		// Complex patterns
		{"interleaved", []int{1, 3, 5, 2, 4, 6}, 6},
		{"pyramid_pattern", []int{5, 4, 6, 3, 7, 2, 8, 1, 9}, 9},
		{"zigzag", []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 10},

		// Longer sequences
		{"sequence_10", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 10},
		{"sequence_with_noise", []int{1, 100, 2, 200, 3, 300, 4, 400, 5}, 5},

		// Isolated elements among sequences
		{"isolated_among_sequences", []int{1, 2, 3, 100, 5, 6, 7, 8}, 4},

		// Multiple duplicates with sequences
		{"complex_duplicates", []int{1, 2, 0, 1, 2, 3, 2, 1, 0}, 4},

		// Only two numbers far apart
		{"far_apart_two", []int{-1000000000, 1000000000}, 1},

		// Sequence starting from zero
		{"from_zero", []int{0, 1, 2, 3, 4}, 5},
		{"to_zero", []int{-4, -3, -2, -1, 0}, 5},

		// Three element variations
		{"three_consecutive", []int{1, 2, 3}, 3},
		{"three_with_gap", []int{1, 2, 4}, 2},
		{"three_isolated", []int{1, 5, 9}, 1},

		// Duplicates at boundaries
		{"duplicate_at_start", []int{1, 1, 2, 3}, 3},
		{"duplicate_at_end", []int{1, 2, 3, 3}, 3},
		{"duplicate_in_middle", []int{1, 2, 2, 3}, 3},

		// Long sequence with many duplicates
		{"long_with_duplicates", []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 5},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := longestConsecutive(tst.nums)
			if got != tst.want {
				t.Errorf("longestConsecutive(%v) = %v, want %v", tst.nums, got, tst.want)
			}

			got = longestConsecutive2(tst.nums)
			if got != tst.want {
				t.Errorf("longestConsecutive(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}
