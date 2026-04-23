package leetcode0215

import "testing"

func TestFindKthLargestExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			"example_1",
			[]int{3, 2, 1, 5, 6, 4},
			2,
			5,
		},
		{
			"example_2",
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			4,
			4,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := findKthLargest(tst.nums, tst.k)
			if got != tst.want {
				t.Errorf("findKthLargest(%v, %d) = %d, want %d", tst.nums, tst.k, got, tst.want)
			}

			got = findKthLargest2(tst.nums, tst.k)
			if got != tst.want {
				t.Errorf("findKthLargest2(%v, %d) = %d, want %d", tst.nums, tst.k, got, tst.want)
			}
		})
	}
}

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"single_element", []int{1}, 1, 1},
		{"two_elements_k1", []int{1, 2}, 1, 2},
		{"two_elements_k2", []int{1, 2}, 2, 1},
		{"all_same", []int{5, 5, 5, 5}, 2, 5},
		{"all_same_max_length", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 10, 1},
		{"negative_numbers", []int{-1, -2, -3, -4, -5}, 3, -3},
		{"mixed_positive_negative", []int{-5, 10, -3, 8, -1, 0}, 2, 8},
		{"contains_zero", []int{0, 0, 0, 0}, 1, 0},
		{"k_equals_length", []int{1, 2, 3, 4, 5}, 5, 1},
		{"k_equals_1", []int{5, 3, 9, 1, 7}, 1, 9},
		{"sorted_ascending", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 6},
		{"sorted_descending", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 5, 6},
		{"alternating_high_low", []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 7, 4},
		{"max_value_element", []int{10000, 10000, 10000}, 1, 10000},
		{"min_value_element", []int{-10000, -10000, -10000}, 3, -10000},
		{"peak_in_middle", []int{1, 3, 5, 7, 9, 8, 6, 4, 2}, 1, 9},
		{"valley_in_middle", []int{9, 7, 5, 3, 1, 2, 4, 6, 8}, 9, 1},
		{"all_distinct", []int{10, 20, 30, 40, 50, 60}, 3, 40},
		{"duplicates_at_start", []int{5, 5, 5, 1, 2, 3, 4}, 4, 4},
		{"duplicates_at_end", []int{1, 2, 3, 4, 5, 5, 5}, 4, 4},
		{"duplicates_in_middle", []int{1, 3, 5, 5, 5, 7, 9}, 3, 5},
		{"single_large_element", []int{1, 1, 1, 1, 10000, 1, 1, 1, 1}, 1, 10000},
		{"single_small_element", []int{9999, 9999, 9999, 9999, -10000, 9999, 9999}, 7, -10000},
		{"three_elements", []int{3, 1, 2}, 2, 2},
		{"three_elements_sorted", []int{1, 2, 3}, 2, 2},
		{"three_elements_reversed", []int{3, 2, 1}, 2, 2},
		{"even_length_array", []int{4, 3, 2, 1}, 2, 3},
		{"odd_length_array", []int{5, 4, 3, 2, 1}, 3, 3},
		{"same_as_example1_shuffled", []int{6, 5, 4, 3, 2, 1}, 2, 5},
		{"same_as_example2_shuffled", []int{6, 5, 5, 4, 3, 3, 2, 2, 1}, 4, 4},
		{"k_middle_of_long_array", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 8, 8},
		{"all_positive", []int{100, 200, 300, 400, 500}, 3, 300},
		{"all_negative", []int{-100, -200, -300, -400, -500}, 3, -300},
		{"zero_in_middle", []int{-5, -4, 0, 4, 5}, 3, 0},
		{"k1_in_large_array", []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10, 11, 13, 15, 12, 14}, 1, 15},
		{"k_last_in_large_array", []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 15, 1},
		{"alternating_duplicates", []int{1, 2, 1, 2, 1, 2, 1, 2}, 5, 1},
		{"ascending_duplicates", []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 6, 3},
		{"descending_duplicates", []int{5, 5, 4, 4, 3, 3, 2, 2, 1, 1}, 6, 3},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := findKthLargest(tst.nums, tst.k)
			if got != tst.want {
				t.Errorf("findKthLargest(%v, %d) = %d, want %d", tst.nums, tst.k, got, tst.want)
			}

			got = findKthLargest2(tst.nums, tst.k)
			if got != tst.want {
				t.Errorf("findKthLargest2(%v, %d) = %d, want %d", tst.nums, tst.k, got, tst.want)
			}
		})
	}
}
