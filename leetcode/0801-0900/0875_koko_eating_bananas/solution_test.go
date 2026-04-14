// Package leetcode0875 solves LeetCode 875. Koko Eating Bananas
package leetcode0875

import "testing"

func TestMinEatingSpeedExamples(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{
			"example_1",
			[]int{3, 6, 7, 11},
			8,
			4,
		},
		{
			"example_2",
			[]int{30, 11, 23, 4, 20},
			5,
			30,
		},
		{
			"example_3",
			[]int{30, 11, 23, 4, 20},
			6,
			23,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			// Create a copy of piles to avoid modifying the original since the solution sorts it
			pilesCopy := make([]int, len(tst.piles))
			copy(pilesCopy, tst.piles)

			got := minEatingSpeed(pilesCopy, tst.h)
			if got != tst.want {
				t.Errorf("minEatingSpeed(%v, %d) = %d, want %d", tst.piles, tst.h, got, tst.want)
			}
		})
	}
}

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		// Edge cases
		{"single_pile_single_banana", []int{1}, 1, 1},
		{"single_pile_multiple_bananas", []int{5}, 1, 5},
		{"single_pile_exact_hours", []int{10}, 5, 2},
		{"single_pile_more_hours_than_bananas", []int{5}, 10, 1},
		{"all_ones", []int{1, 1, 1, 1, 1}, 5, 1},
		{"all_ones_more_hours", []int{1, 1, 1, 1, 1}, 10, 1},
		{"all_same_large", []int{100, 100, 100}, 3, 100},
		{"all_same_large_more_hours", []int{100, 100, 100}, 6, 50},

		// Two piles
		{"two_piles_equal", []int{3, 3}, 2, 3},
		{"two_piles_equal_more_hours", []int{3, 3}, 3, 3},
		{"two_piles_unequal", []int{5, 8}, 2, 8},
		{"two_piles_unequal_more_hours", []int{5, 8}, 3, 5},
		{"two_piles_one_large", []int{1, 100}, 2, 100},
		{"two_piles_one_large_more_hours", []int{1, 100}, 101, 1},

		// Small arrays
		{"small_array_1", []int{1, 2, 3, 4, 5}, 5, 5},
		{"small_array_2", []int{1, 2, 3, 4, 5}, 7, 3},
		{"small_array_3", []int{1, 2, 3, 4, 5}, 15, 1},

		// Speed boundary tests
		{"speed_exact_divisible", []int{4, 4, 4}, 3, 4},
		{"speed_with_remainder", []int{5, 5, 5}, 3, 5},
		{"speed_with_remainder_2", []int{5, 5, 5}, 4, 5},
		{"speed_with_mixed_remainders", []int{3, 5, 7}, 5, 4},

		// Large value in piles
		{"single_large_value", []int{1000000000}, 1, 1000000000},
		{"single_large_value_many_hours", []int{1000000000}, 1000000000, 1},
		{"multiple_large_values", []int{1000000000, 1000000000}, 2, 1000000000},

		// h equals piles.length (must eat at max pile speed)
		{"h_equals_piles_length_1", []int{1, 2, 3, 4, 5}, 5, 5},
		{"h_equals_piles_length_2", []int{10, 20, 30}, 3, 30},
		{"h_equals_piles_length_3", []int{5, 1, 3}, 3, 5},

		// h just slightly larger than piles.length
		{"h_just_larger_1", []int{3, 6, 7, 11}, 5, 7},
		{"h_just_larger_2", []int{30, 11, 23, 4, 20}, 7, 20},
		{"h_just_larger_3", []int{1, 2, 3, 4, 5, 6}, 7, 5},

		// Already sorted piles
		{"already_sorted_asc", []int{1, 2, 3, 4, 5}, 7, 3},
		{"already_sorted_desc", []int{5, 4, 3, 2, 1}, 7, 3},

		// One pile much larger than others
		{"one_giant_pile", []int{1, 1, 1, 100}, 4, 100},
		{"one_giant_pile_more_hours", []int{1, 1, 1, 100}, 5, 50},
		{"one_giant_pile_many_hours", []int{1, 1, 1, 100}, 103, 1},

		// Alternating small and large
		{"alternating_small_large", []int{1, 100, 1, 100, 1, 100}, 6, 100},
		{"alternating_small_large_more_hours", []int{1, 100, 1, 100, 1, 100}, 9, 50},

		// Various sizes mixed
		{"mixed_sizes_1", []int{5, 8, 3, 12, 7, 10}, 10, 6},
		{"mixed_sizes_2", []int{9, 3, 8, 4, 7, 5}, 8, 7},
		{"mixed_sizes_3", []int{15, 1, 1, 1, 1, 1}, 7, 8},
		{"mixed_sizes_4", []int{15, 1, 1, 1, 1, 1}, 8, 5},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			// Create a copy of piles to avoid modifying the original since the solution sorts it
			pilesCopy := make([]int, len(tst.piles))
			copy(pilesCopy, tst.piles)

			got := minEatingSpeed(pilesCopy, tst.h)
			if got != tst.want {
				t.Errorf("minEatingSpeed(%v, %d) = %d, want %d", tst.piles, tst.h, got, tst.want)
			}
		})
	}
}
