package leetcode0287

import "testing"

func TestFindDuplicateExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"example_1",
			[]int{1, 3, 4, 2, 2},
			2,
		},
		{
			"example_2",
			[]int{3, 1, 3, 4, 2},
			3,
		},
		{
			"example_3",
			[]int{3, 3, 3, 3, 3},
			3,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := findDuplicate(tst.nums)
			if got != tst.want {
				t.Errorf("findDuplicate(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"minimum_n_1", []int{1, 1}, 1},
		{"minimum_n_2_duplicate_at_start", []int{1, 1, 2}, 1},
		{"minimum_n_2_duplicate_at_end", []int{1, 2, 2}, 2},
		{"duplicate_twice", []int{1, 2, 3, 4, 2}, 2},
		{"duplicate_three_times", []int{1, 2, 2, 3, 2}, 2},
		{"duplicate_at_beginning", []int{2, 2, 3, 4, 1}, 2},
		{"duplicate_at_end", []int{1, 3, 4, 2, 2}, 2},
		{"large_numbers_n_10", []int{9, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}, 9},
		{"all_in_order_except_duplicate", []int{1, 2, 3, 4, 5, 5, 6, 7}, 5},
		{"reverse_order_with_duplicate", []int{7, 6, 5, 4, 3, 2, 1, 4}, 4},
		{"duplicate_in_middle", []int{1, 5, 3, 4, 5, 2}, 5},
		{"duplicate_followed_by_itself", []int{1, 2, 3, 3, 4, 5}, 3},
		{"only_two_distinct_numbers", []int{1, 2, 2, 2}, 2},
		{"cycle_detection_case_1", []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}, 9},
		{"cycle_detection_case_2", []int{2, 6, 4, 1, 3, 1, 5}, 1},
		{"cycle_detection_case_3", []int{4, 2, 1, 3, 2}, 2},
		{"single_duplicate_in_large_range", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 10}, 10},
		{"n_5_duplicate_1", []int{1, 2, 3, 4, 5, 1}, 1},
		{"n_5_duplicate_5", []int{1, 2, 3, 4, 5, 5}, 5},
		{"n_5_duplicate_3", []int{1, 2, 3, 4, 5, 3}, 3},
		{"zigzag_pattern", []int{1, 3, 2, 5, 4, 5, 6}, 5},
		{"duplicate_at_even_indices", []int{3, 1, 3, 2, 3, 4}, 3},
		{"duplicate_at_odd_indices", []int{1, 4, 2, 4, 3, 4}, 4},
		{"long_chain_before_cycle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 5}, 5},
		{"immediate_cycle", []int{1, 1, 2, 3, 4, 5, 6}, 1},
		{"two_pairs_but_only_one_duplicate", []int{1, 2, 2, 3, 4, 5, 6}, 2},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := findDuplicate(tst.nums)
			if got != tst.want {
				t.Errorf("findDuplicate(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}
