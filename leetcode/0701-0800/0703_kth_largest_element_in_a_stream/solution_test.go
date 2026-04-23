package leetcode0703

import "testing"

func TestKthLargestExamples(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		initial  []int
		addVals  []int
		expected []int
	}{
		{
			"example_1",
			3,
			[]int{4, 5, 8, 2},
			[]int{3, 5, 10, 9, 4},
			[]int{4, 5, 5, 8, 8},
		},
		{
			"example_2",
			4,
			[]int{7, 7, 7, 7, 8, 3},
			[]int{2, 10, 9, 9},
			[]int{7, 7, 7, 8},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			obj := Constructor(tst.k, tst.initial)
			for i, val := range tst.addVals {
				got := obj.Add(val)
				if got != tst.expected[i] {
					t.Errorf("Add(%d) = %d, want %d (step %d)", val, got, tst.expected[i], i+1)
				}
			}
		})
	}
}

func TestKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		initial  []int
		addVals  []int
		expected []int
	}{
		{
			"k_1_single_element",
			1,
			[]int{5},
			[]int{3, 7, 2, 9},
			[]int{5, 7, 7, 9},
		},
		{
			"k_1_empty_initial",
			1,
			[]int{},
			[]int{3, 1, 4},
			[]int{3, 3, 4},
		},
		{
			"k_equal_to_length",
			3,
			[]int{1, 2, 3},
			[]int{0, 4, 5},
			[]int{1, 2, 3},
		},
		{
			"k_equals_initial_length_plus_1",
			4,
			[]int{1, 2, 3},
			[]int{4, 5, 0},
			[]int{1, 2, 2},
		},
		{
			"all_negative_numbers",
			2,
			[]int{-5, -3, -10},
			[]int{-1, -8, -2},
			[]int{-3, -3, -2},
		},
		{
			"mixed_positive_negative",
			3,
			[]int{-1, 5, 0, -3, 10},
			[]int{-10, 7, 2, -5},
			[]int{0, 5, 5, 5},
		},
		{
			"duplicate_values",
			2,
			[]int{5, 5, 5},
			[]int{5, 4, 6, 6},
			[]int{5, 5, 5, 6},
		},
		{
			"decreasing_sequence",
			2,
			[]int{10, 9, 8, 7},
			[]int{6, 5, 11, 4},
			[]int{9, 9, 10, 10},
		},
		{
			"increasing_sequence",
			2,
			[]int{1, 2, 3, 4},
			[]int{5, 6, 0, 7},
			[]int{4, 5, 5, 6},
		},
		{
			"single_add_call",
			2,
			[]int{10, 20},
			[]int{15},
			[]int{15},
		},
		{
			"add_same_as_current_kth",
			3,
			[]int{1, 3, 5, 7, 9},
			[]int{5, 5, 5},
			[]int{5, 5, 5},
		},
		{
			"add_smaller_than_current_kth",
			3,
			[]int{10, 20, 30, 40, 50},
			[]int{5, 15, 25},
			[]int{30, 30, 30},
		},
		{
			"add_larger_than_current_kth",
			3,
			[]int{10, 20, 30, 40, 50},
			[]int{35, 45, 55},
			[]int{35, 40, 45},
		},
		{
			"zero_in_values",
			2,
			[]int{0, 0, 0},
			[]int{0, -1, 1},
			[]int{0, 0, 0},
		},
		{
			"max_value_boundary",
			2,
			[]int{10000, 9999},
			[]int{10000, 9998},
			[]int{10000, 10000},
		},
		{
			"min_value_boundary",
			2,
			[]int{-10000, -9999},
			[]int{-10000, -9998},
			[]int{-10000, -9999},
		},
		{
			"k_1_with_negatives",
			1,
			[]int{-100, -200, -50},
			[]int{-300, -10, -75},
			[]int{-50, -10, -10},
		},
		{
			"initial_already_sorted",
			3,
			[]int{1, 2, 3, 4, 5},
			[]int{0, 6, 3},
			[]int{3, 4, 4},
		},
		{
			"initial_reverse_sorted",
			3,
			[]int{5, 4, 3, 2, 1},
			[]int{6, 0, 4},
			[]int{4, 4, 4},
		},
		{
			"k_2_empty_initial",
			2,
			[]int{},
			[]int{5, 3, 7, 1, 9},
			[]int{5, 3, 5, 5, 7},
		},
		{
			"alternating_adds",
			3,
			[]int{10, 20, 30},
			[]int{5, 35, 8, 40, 3},
			[]int{10, 20, 20, 30, 30},
		},
		{
			"all_same_initial",
			5,
			[]int{7, 7, 7, 7, 7, 7, 7},
			[]int{7, 7, 6, 8},
			[]int{7, 7, 7, 7},
		},
		{
			"large_k_small_initial",
			5,
			[]int{1, 2},
			[]int{3, 4, 5, 6, 0},
			[]int{1, 1, 1, 2, 2},
		},
		{
			"value_at_min_boundary",
			2,
			[]int{-10000, 5000},
			[]int{-10000, 10000},
			[]int{-10000, 5000},
		},
		{
			"value_at_max_boundary",
			2,
			[]int{10000, -5000},
			[]int{10000, -10000},
			[]int{10000, 10000},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			obj := Constructor(tst.k, tst.initial)
			for i, val := range tst.addVals {
				got := obj.Add(val)
				if got != tst.expected[i] {
					t.Errorf("Add(%d) = %d, want %d (step %d)", val, got, tst.expected[i], i+1)
				}
			}
		})
	}
}
