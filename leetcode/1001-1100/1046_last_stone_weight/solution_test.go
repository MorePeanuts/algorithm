// Package leetcode1046 solves LeetCode 1046. Last Stone Weight
package leetcode1046

import "testing"

func TestLastStoneWeightExamples(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{
			"example_1",
			[]int{2, 7, 4, 1, 8, 1},
			1,
		},
		{
			"example_2",
			[]int{1},
			1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			// Make a copy to avoid modifying the original test data
			stonesCopy := make([]int, len(tst.stones))
			copy(stonesCopy, tst.stones)

			got := lastStoneWeight(stonesCopy)
			if got != tst.want {
				t.Errorf("lastStoneWeight(%v) = %v, want %v", tst.stones, got, tst.want)
			}
		})
	}
}

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{"single_element", []int{5}, 5},
		{"two_equal_elements", []int{5, 5}, 0},
		{"two_unequal_elements", []int{10, 5}, 5},
		{"all_same_elements_even", []int{3, 3, 3, 3}, 0},
		{"all_same_elements_odd", []int{3, 3, 3, 3, 3}, 3},
		{"max_single_element", []int{1000}, 1000},
		{"two_max_elements", []int{1000, 1000}, 0},
		{"three_elements", []int{8, 6, 4}, 2},
		{"four_elements_pairwise", []int{10, 10, 5, 5}, 0},
		{"pairwise_same", []int{2, 2, 4, 4, 6, 6}, 0},
		{"three_sevens", []int{7, 7, 7}, 7},
		{"five_nines", []int{9, 9, 9, 9, 9}, 9},
		{"empty_after_smash", []int{2, 2}, 0},
		{"one_left_after_smash", []int{5, 3}, 2},
		{"small_array_1", []int{1, 3}, 2},
		{"small_array_2", []int{1, 2, 3}, 0},
		{"small_array_3", []int{2, 4, 6}, 0},
		{"max_length_all_ones", makeMaxLengthSlice(1), 0},
		{"max_length_all_twos", makeMaxLengthSlice(2), 0},
		{"max_length_one_extra", append(makeMaxLengthSlice(1)[:29], 2), 1},
		{"only_two_large", []int{1000, 999}, 1},
		{"multiple_smash_steps", []int{5, 5, 5}, 5},
		{"smash_creates_new_pairs", []int{4, 4, 4, 4, 4}, 4},
		{"simple_chain", []int{10, 5, 5}, 0},
		{"two_pairs", []int{1, 1, 2, 2}, 0},
		{"one_large_small_remainder", []int{100, 1, 1}, 98},
		{"even_number_same", []int{5, 5, 5, 5}, 0},
		{"odd_number_same", []int{5, 5, 5, 5, 5}, 5},
		{"all_ones_odd", []int{1, 1, 1}, 1},
		{"all_ones_even", []int{1, 1, 1, 1}, 0},
		{"mixed_sizes", []int{10, 20, 10, 20}, 0},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			// Make a copy to avoid modifying the original test data
			stonesCopy := make([]int, len(tst.stones))
			copy(stonesCopy, tst.stones)

			got := lastStoneWeight(stonesCopy)
			if got != tst.want {
				t.Errorf("lastStoneWeight(%v) = %v, want %v", tst.stones, got, tst.want)
			}
		})
	}
}

func makeMaxLengthSlice(val int) []int {
	stones := make([]int, 30)
	for i := range stones {
		stones[i] = val
	}
	return stones
}

