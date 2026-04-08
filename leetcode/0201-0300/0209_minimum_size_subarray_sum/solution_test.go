package leetcode0209

import "testing"

func TestMinSubArrayLenExamples(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "example_1",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			name:   "example_2",
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		{
			name:   "example_3",
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := minSubArrayLen(tst.target, tst.nums)
			if got != tst.want {
				t.Errorf("minSubArrayLen(%d, %v) = %d, want %d", tst.target, tst.nums, got, tst.want)
			}
			got = minSubArrayLen2(tst.target, tst.nums)
			if got != tst.want {
				t.Errorf("minSubArrayLen(%d, %v) = %d, want %d", tst.target, tst.nums, got, tst.want)
			}
		})
	}
}

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		// Single element cases
		{"single_element_equal", 5, []int{5}, 1},
		{"single_element_greater", 3, []int{5}, 1},
		{"single_element_less", 10, []int{5}, 0},

		// Two elements
		{"two_elements_first_enough", 3, []int{5, 1}, 1},
		{"two_elements_second_enough", 3, []int{1, 5}, 1},
		{"two_elements_sum_needed", 5, []int{3, 3}, 2},
		{"two_elements_sum_not_enough", 10, []int{3, 3}, 0},

		// Target equals sum of all elements
		{"sum_equals_target", 10, []int{2, 3, 1, 4}, 4},
		{"sum_less_than_target", 20, []int{2, 3, 1, 4}, 0},

		// First element is enough
		{"first_element_enough", 5, []int{10, 1, 2, 3, 4}, 1},

		// Last element is enough
		{"last_element_enough", 5, []int{1, 2, 3, 4, 10}, 1},

		// Middle element is enough
		{"middle_element_enough", 5, []int{1, 2, 10, 3, 4}, 1},

		// Optimal window at different positions
		{"optimal_at_start", 5, []int{3, 3, 1, 1, 1, 1}, 2},
		{"optimal_at_end", 5, []int{1, 1, 1, 1, 3, 3}, 2},
		{"optimal_in_middle", 5, []int{1, 1, 3, 3, 1, 1}, 2},

		// All elements same
		{"all_ones_target_3", 3, []int{1, 1, 1, 1, 1}, 3},
		{"all_twos_target_5", 5, []int{2, 2, 2, 2, 2}, 3},
		{"all_fives_target_5", 5, []int{5, 5, 5, 5, 5}, 1},
		{"all_fives_target_10", 10, []int{5, 5, 5, 5, 5}, 2},

		// Increasing sequence
		{"increasing_seq", 9, []int{1, 2, 3, 4, 5}, 2},
		{"increasing_need_all", 15, []int{1, 2, 3, 4, 5}, 5},

		// Decreasing sequence
		{"decreasing_seq", 9, []int{5, 4, 3, 2, 1}, 2},
		{"decreasing_need_all", 15, []int{5, 4, 3, 2, 1}, 5},

		// Large single element
		{"large_single_satisfies", 100, []int{1, 2, 100, 3, 4}, 1},
		{"large_single_not_enough", 150, []int{1, 2, 100, 3, 4}, 0},

		// Edge: target is 1 (minimum constraint)
		{"target_1_any_element", 1, []int{5, 3, 2, 1}, 1},
		{"target_1_all_ones", 1, []int{1, 1, 1, 1}, 1},

		// Multiple valid windows, find minimum
		{"multiple_windows", 6, []int{2, 2, 2, 3, 3, 2, 2, 2}, 2},

		// Boundary at exact target
		{"exact_target_sum", 6, []int{1, 2, 3, 4, 5}, 2},

		// Larger arrays
		{"larger_array_scattered", 15, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2},
		{"larger_array_need_many", 45, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
		{"larger_array_impossible", 100, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},

		// Special patterns
		{"alternating_1_10", 15, []int{1, 10, 1, 10, 1, 10}, 3},
		{"spike_in_middle", 50, []int{1, 1, 1, 50, 1, 1, 1}, 1},
		{"two_spikes", 30, []int{1, 20, 1, 1, 20, 1}, 4},

		// Contiguous subarray vs scattered
		{"contiguous_better", 10, []int{1, 1, 5, 6, 1, 1}, 2},

		// Window shrinking scenarios
		{"shrink_from_left", 7, []int{1, 1, 1, 4, 4}, 2},
		{"shrink_from_right", 7, []int{4, 4, 1, 1, 1}, 2},

		// Near boundary values
		{"small_elements_big_target", 50, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8},
		{"all_max_value_10000", 20000, []int{10000, 10000}, 2},
		{"all_max_value_single", 10000, []int{10000}, 1},

		// Subarray starts from index 0
		{"starts_from_zero", 6, []int{3, 4, 1, 1, 1}, 2},

		// Subarray ends at last index
		{"ends_at_last", 6, []int{1, 1, 1, 3, 4}, 2},

		// Only one valid subarray
		{"only_one_valid", 10, []int{1, 1, 1, 10, 1, 1}, 1},

		{"leetcode_case18", 213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}, 8},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := minSubArrayLen(tst.target, tst.nums)
			if got != tst.want {
				t.Errorf("minSubArrayLen(%d, %v) = %d, want %d", tst.target, tst.nums, got, tst.want)
			}
			got = minSubArrayLen2(tst.target, tst.nums)
			if got != tst.want {
				t.Errorf("minSubArrayLen(%d, %v) = %d, want %d", tst.target, tst.nums, got, tst.want)
			}
		})
	}
}
