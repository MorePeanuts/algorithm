package leetcode0347

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestTopKFrequentExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "example_1",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "example_2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "example_3",
			nums: []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:    2,
			want: []int{1, 2},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := topKFrequent(tst.nums, tst.k)
			if !utils.MatchIntSlice(got, tst.want) {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v", tst.nums, tst.k, got, tst.want)
			}
		})
	}
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		// Single element
		{name: "single_element", nums: []int{5}, k: 1, want: []int{5}},
		{name: "single_negative", nums: []int{-1}, k: 1, want: []int{-1}},
		{name: "single_zero", nums: []int{0}, k: 1, want: []int{0}},

		// All same elements
		{name: "all_same_two", nums: []int{7, 7}, k: 1, want: []int{7}},
		{name: "all_same_five", nums: []int{3, 3, 3, 3, 3}, k: 1, want: []int{3}},

		// Two distinct elements
		{name: "two_distinct_k1", nums: []int{1, 1, 2}, k: 1, want: []int{1}},
		{name: "two_distinct_k2", nums: []int{1, 1, 2}, k: 2, want: []int{1, 2}},
		{name: "two_distinct_equal_freq", nums: []int{1, 2, 1, 2}, k: 2, want: []int{1, 2}},

		// Three distinct elements
		{name: "three_distinct_k1", nums: []int{1, 1, 1, 2, 2, 3}, k: 1, want: []int{1}},
		{name: "three_distinct_k2", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{name: "three_distinct_k3", nums: []int{1, 1, 1, 2, 2, 3}, k: 3, want: []int{1, 2, 3}},

		// Negative numbers
		{name: "negative_nums_k1", nums: []int{-1, -1, -2, -2, -2, -3}, k: 1, want: []int{-2}},
		{name: "negative_nums_k2", nums: []int{-1, -1, -2, -2, -2, -3}, k: 2, want: []int{-2, -1}},
		{name: "all_negative", nums: []int{-5, -5, -5, -3, -3, -1}, k: 2, want: []int{-5, -3}},

		// Mixed positive and negative
		{name: "mixed_pos_neg_k1", nums: []int{-1, -1, 2, 2, 2, 3}, k: 1, want: []int{2}},
		{name: "mixed_pos_neg_k2", nums: []int{-1, -1, 2, 2, 2, 3}, k: 2, want: []int{2, -1}},
		{name: "mixed_with_zero", nums: []int{0, 0, 0, 1, 1, -1}, k: 1, want: []int{0}},
		{name: "mixed_with_zero_k2", nums: []int{0, 0, 0, 1, 1, -1}, k: 2, want: []int{0, 1}},

		// Edge values (within constraints: -10^4 to 10^4)
		{name: "max_value", nums: []int{10000, 10000, 10000, 1}, k: 1, want: []int{10000}},
		{name: "min_value", nums: []int{-10000, -10000, -10000, 1}, k: 1, want: []int{-10000}},
		{name: "max_min_values", nums: []int{10000, 10000, -10000, -10000, -10000}, k: 1, want: []int{-10000}},
		{name: "max_min_values_k2", nums: []int{10000, 10000, -10000, -10000, -10000}, k: 2, want: []int{-10000, 10000}},

		// K equals number of distinct elements
		{name: "k_equals_distinct_count", nums: []int{1, 2, 3, 4, 5}, k: 5, want: []int{1, 2, 3, 4, 5}},
		{name: "k_equals_distinct_varied_freq", nums: []int{1, 1, 1, 2, 2, 3}, k: 3, want: []int{1, 2, 3}},

		// Larger frequency differences
		{name: "large_freq_diff", nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, k: 1, want: []int{1}},
		{name: "large_freq_diff_k2", nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, k: 2, want: []int{1, 2}},

		// Multiple elements with same frequency
		{name: "same_freq_all", nums: []int{1, 2, 3, 4}, k: 4, want: []int{1, 2, 3, 4}},
		{name: "same_freq_partial", nums: []int{1, 1, 2, 2, 3, 3, 4}, k: 3, want: []int{1, 2, 3}},

		// Consecutive numbers
		{name: "consecutive_nums", nums: []int{1, 2, 3, 1, 2, 1}, k: 2, want: []int{1, 2}},

		// Scattered duplicates
		{name: "scattered_duplicates", nums: []int{1, 2, 1, 3, 1, 2, 4, 2, 5}, k: 2, want: []int{1, 2}},

		// Medium size arrays
		{name: "medium_array_k1", nums: []int{4, 1, -1, 2, -1, 2, 3, 3, 3, 3}, k: 1, want: []int{3}},
		{name: "medium_array_k3", nums: []int{4, 1, -1, 2, -1, 2, 3, 3, 3, 3}, k: 3, want: []int{3, 2, -1}},

		// All elements appear once except one
		{name: "one_repeated", nums: []int{1, 2, 3, 4, 5, 5}, k: 1, want: []int{5}},
		{name: "first_repeated", nums: []int{1, 1, 2, 3, 4, 5}, k: 1, want: []int{1}},

		// Long sequences of same element
		{name: "long_sequence", nums: []int{9, 9, 9, 9, 9, 9, 9, 9, 1, 2}, k: 1, want: []int{9}},

		// Various k values on same array
		{name: "varied_k_array_k1", nums: []int{5, 5, 5, 4, 4, 3, 2, 1}, k: 1, want: []int{5}},
		{name: "varied_k_array_k2", nums: []int{5, 5, 5, 4, 4, 3, 2, 1}, k: 2, want: []int{5, 4}},
		{name: "varied_k_array_k5", nums: []int{5, 5, 5, 4, 4, 3, 2, 1}, k: 5, want: []int{5, 4, 3, 2, 1}},

		// Zeros mixed in
		{name: "zeros_dominant", nums: []int{0, 0, 0, 0, 1, 2, 3}, k: 1, want: []int{0}},
		{name: "zeros_not_dominant", nums: []int{0, 1, 1, 1, 2, 2}, k: 1, want: []int{1}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := topKFrequent(tst.nums, tst.k)
			if !utils.MatchIntSlice(got, tst.want) {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v", tst.nums, tst.k, got, tst.want)
			}
		})
	}
}
