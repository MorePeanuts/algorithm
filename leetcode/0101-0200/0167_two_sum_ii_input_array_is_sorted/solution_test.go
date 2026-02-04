package leetcode0167

import (
	"reflect"
	"testing"
)

func TestTwoSumExamples(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			"example_1",
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 2},
		},
		{
			"example_2",
			[]int{2, 3, 4},
			6,
			[]int{1, 3},
		},
		{
			"example_3",
			[]int{-1, 0},
			-1,
			[]int{1, 2},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := twoSum(tst.numbers, tst.target)
			if !reflect.DeepEqual(got, tst.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tst.numbers, tst.target, got, tst.want)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		// Minimum length (2 elements)
		{"min_length_positive", []int{1, 2}, 3, []int{1, 2}},
		{"min_length_negative", []int{-2, -1}, -3, []int{1, 2}},
		{"min_length_mixed", []int{-5, 5}, 0, []int{1, 2}},
		{"min_length_zeros", []int{0, 0}, 0, []int{1, 2}},

		// Boundary values for numbers
		{"max_value_pair", []int{999, 1000}, 1999, []int{1, 2}},
		{"min_value_pair", []int{-1000, -999}, -1999, []int{1, 2}},
		{"extreme_range", []int{-1000, 1000}, 0, []int{1, 2}},

		// Target at boundaries
		{"target_max", []int{0, 1000}, 1000, []int{1, 2}},
		{"target_min", []int{-1000, 0}, -1000, []int{1, 2}},
		{"target_zero_positive", []int{-500, 500}, 0, []int{1, 2}},

		// First and last elements
		{"first_last_3_elements", []int{1, 5, 9}, 10, []int{1, 3}},
		{"first_last_5_elements", []int{1, 2, 3, 4, 10}, 11, []int{1, 5}},
		{"first_last_large", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, []int{1, 10}},

		// Adjacent elements
		{"adjacent_beginning", []int{1, 2, 5, 8}, 3, []int{1, 2}},
		{"adjacent_middle", []int{1, 3, 4, 8}, 7, []int{2, 3}},
		{"adjacent_end", []int{1, 5, 8, 9}, 17, []int{3, 4}},

		// Elements far apart
		{"far_apart_small", []int{1, 2, 3, 4, 5}, 6, []int{1, 5}},
		{"far_apart_medium", []int{1, 3, 5, 7, 9, 11}, 12, []int{1, 6}},

		// Duplicates in array
		{"duplicates_at_start", []int{1, 1, 3, 5}, 2, []int{1, 2}},
		{"duplicates_at_end", []int{1, 3, 5, 5}, 10, []int{3, 4}},
		{"many_duplicates", []int{1, 2, 2, 2, 2, 3}, 5, []int{2, 6}},

		// Negative numbers
		{"all_negative", []int{-8, -5, -3, -1}, -6, []int{2, 4}},

		// Zero in array
		{"zero_at_start", []int{0, 1, 2, 3}, 3, []int{1, 4}},
		{"zero_at_end", []int{-3, -2, -1, 0}, -3, []int{1, 4}},
		{"multiple_zeros", []int{-1, 0, 0, 0, 1}, 0, []int{1, 5}},

		// Larger arrays
		{"ten_elements", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 19, []int{9, 10}},
		{"sequential_positive", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 23, []int{11, 12}},
		{"sequential_mixed", []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}, 0, []int{1, 11}},

		// Sparse arrays (large gaps)
		{"sparse_positive", []int{1, 100, 200, 300}, 101, []int{1, 2}},
		{"sparse_mixed", []int{-500, -100, 0, 100, 500}, 0, []int{1, 5}},

		// Single pair possible
		{"only_pair_works", []int{1, 2, 4, 8, 16}, 3, []int{1, 2}},
		{"last_pair_only", []int{1, 10, 100, 890, 900}, 1790, []int{4, 5}},

		// Target equals one element doubled (duplicates needed)
		{"double_element", []int{3, 3, 6, 9}, 6, []int{1, 2}},

		// Large target requiring last elements
		{"large_target", []int{100, 200, 300, 400, 500}, 900, []int{4, 5}},
		{"small_target_negative", []int{-500, -400, -300, -200, -100}, -900, []int{1, 2}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := twoSum(tst.numbers, tst.target)
			if !reflect.DeepEqual(got, tst.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tst.numbers, tst.target, got, tst.want)
			}
		})
	}
}
