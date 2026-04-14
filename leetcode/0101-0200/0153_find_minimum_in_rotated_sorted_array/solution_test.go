// Package leetcode0153 solves LeetCode 153. Find Minimum in Rotated Sorted Array
package leetcode0153

import "testing"

func TestFindMinExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"example_1",
			[]int{3, 4, 5, 1, 2},
			1,
		},
		{
			"example_2",
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
		},
		{
			"example_3",
			[]int{11, 13, 15, 17},
			11,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := findMin(tst.nums)
			if got != tst.want {
				t.Errorf("findMin(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"single_element", []int{1}, 1},
		{"single_element_negative", []int{-5000}, -5000},
		{"single_element_max", []int{5000}, 5000},
		{"two_elements_rotated_once", []int{2, 1}, 1},
		{"two_elements_not_rotated", []int{1, 2}, 1},
		{"three_elements_rotated_once", []int{3, 1, 2}, 1},
		{"three_elements_rotated_twice", []int{2, 3, 1}, 1},
		{"min_at_end", []int{2, 3, 4, 5, 1}, 1},
		{"min_at_beginning", []int{1, 2, 3, 4, 5}, 1},
		{"min_in_middle_left", []int{4, 5, 1, 2, 3}, 1},
		{"min_in_middle_right", []int{3, 4, 5, 1, 2}, 1},
		{"all_negative_rotated", []int{-1, -2, -3, -4, -5}, -5},
		{"all_negative_not_rotated", []int{-5, -4, -3, -2, -1}, -5},
		{"mixed_positive_negative", []int{0, 1, 2, -3, -2, -1}, -3},
		{"rotated_n_minus_1_times", []int{2, 3, 4, 5, 6, 7, 1}, 1},
		{"rotated_n_times", []int{1, 2, 3, 4, 5, 6, 7}, 1},
		{"descending_after_rotating", []int{5, 4, 3, 2, 1}, 1},
		{"large_range_values", []int{-5000, 5000, -4999, 4999, -4998}, -5000},
		{"min_is_max_negative", []int{-5000, -4999, -4998, 0, 1, 2}, -5000},
		{"min_is_smallest_positive", []int{100, 200, 300, 400, 1}, 1},
		{"even_length_array", []int{5, 6, 7, 8, 1, 2, 3, 4}, 1},
		{"odd_length_array", []int{6, 7, 8, 1, 2, 3, 4, 5, 6}, 1},
		{"rotated_at_middle_even", []int{4, 5, 6, 7, 0, 1, 2, 3}, 0},
		{"rotated_at_middle_odd", []int{5, 6, 7, 0, 1, 2, 3, 4, 5}, 0},
		{"left_part_all_larger_than_right", []int{1000, 2000, 3000, 1, 2, 3, 4, 5}, 1},
		{"only_one_rotation", []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}, 1},
		{"almost_sorted_min_at_end", []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 1}, 1},
		{"almost_sorted_min_at_start", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1},
		{"consecutive_decreasing", []int{5, 1, 2, 3, 4}, 1},
		{"min_right_before_middle", []int{5, 6, 1, 2, 3, 4}, 1},
		{"min_right_after_middle", []int{4, 5, 6, 1, 2, 3}, 1},
		{"array_with_zero", []int{1, 2, 3, 4, 0}, 0},
		{"zero_in_middle", []int{3, 4, 0, 1, 2}, 0},
		{"zero_at_start", []int{0, 1, 2, 3, 4}, 0},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := findMin(tst.nums)
			if got != tst.want {
				t.Errorf("findMin(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}
