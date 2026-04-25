package leetcode0075

import (
	"slices"
	"testing"
)

func TestSortColorsExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			"example_1",
			[]int{2, 0, 2, 1, 1, 0},
			[]int{0, 0, 1, 1, 2, 2},
		},
		{
			"example_2",
			[]int{2, 0, 1},
			[]int{0, 1, 2},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			numsCopy := make([]int, len(tst.nums))
			copy(numsCopy, tst.nums)
			sortColors(numsCopy)
			if !slices.Equal(numsCopy, tst.want) {
				t.Errorf("sortColors(%v) = %v, want %v", tst.nums, numsCopy, tst.want)
			}
		})
	}
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"single_element", []int{0}, []int{0}},
		{"single_element_1", []int{1}, []int{1}},
		{"single_element_2", []int{2}, []int{2}},
		{"two_elements_same", []int{0, 0}, []int{0, 0}},
		{"two_elements_0_1", []int{1, 0}, []int{0, 1}},
		{"two_elements_0_2", []int{2, 0}, []int{0, 2}},
		{"two_elements_1_2", []int{2, 1}, []int{1, 2}},
		{"all_zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"all_ones", []int{1, 1, 1}, []int{1, 1, 1}},
		{"all_twos", []int{2, 2, 2}, []int{2, 2, 2}},
		{"already_sorted", []int{0, 0, 1, 1, 2, 2}, []int{0, 0, 1, 1, 2, 2}},
		{"reverse_sorted", []int{2, 2, 1, 1, 0, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"only_0_and_1", []int{1, 0, 1, 0, 1}, []int{0, 0, 1, 1, 1}},
		{"only_0_and_2", []int{2, 0, 2, 0, 2}, []int{0, 0, 2, 2, 2}},
		{"only_1_and_2", []int{2, 1, 2, 1, 2}, []int{1, 1, 2, 2, 2}},
		{"starts_with_2_ends_with_0", []int{2, 1, 0, 2, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"zeros_in_middle", []int{1, 1, 0, 0, 2, 2}, []int{0, 0, 1, 1, 2, 2}},
		{"twos_in_middle", []int{0, 0, 2, 2, 1, 1}, []int{0, 0, 1, 1, 2, 2}},
		{"alternating_colors", []int{0, 1, 2, 0, 1, 2}, []int{0, 0, 1, 1, 2, 2}},
		{"alternating_reverse", []int{2, 1, 0, 2, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"one_zero_rest_ones", []int{1, 1, 0, 1, 1}, []int{0, 1, 1, 1, 1}},
		{"one_one_rest_zeros", []int{0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 1}},
		{"one_two_rest_ones", []int{1, 1, 2, 1, 1}, []int{1, 1, 1, 1, 2}},
		{"one_zero_one_two", []int{1, 0, 1, 2, 1}, []int{0, 1, 1, 1, 2}},
		{"three_elements_2_1_0", []int{2, 1, 0}, []int{0, 1, 2}},
		{"three_elements_2_0_2", []int{2, 0, 2}, []int{0, 2, 2}},
		{"three_elements_1_2_0", []int{1, 2, 0}, []int{0, 1, 2}},
		{"four_elements_2_2_0_0", []int{2, 2, 0, 0}, []int{0, 0, 2, 2}},
		{"four_elements_1_0_2_1", []int{1, 0, 2, 1}, []int{0, 1, 1, 2}},
		{"five_elements_2_0_2_0_2", []int{2, 0, 2, 0, 2}, []int{0, 0, 2, 2, 2}},
		{"all_colors_in_middle", []int{1, 0, 2, 1, 0, 2, 1, 0}, []int{0, 0, 0, 1, 1, 1, 2, 2}},
		{"leading_ones_then_zeros_then_twos", []int{1, 1, 0, 0, 2, 2}, []int{0, 0, 1, 1, 2, 2}},
		{"leading_twos_then_ones_then_zeros", []int{2, 2, 1, 1, 0, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"single_zero_single_one_single_two", []int{2, 0, 1}, []int{0, 1, 2}},
		{"two_zeros_one_one_one_two", []int{2, 0, 1, 0}, []int{0, 0, 1, 2}},
		{"one_zero_two_ones_one_two", []int{2, 0, 1, 1}, []int{0, 1, 1, 2}},
		{"one_zero_one_one_two_twos", []int{2, 0, 1, 2}, []int{0, 1, 2, 2}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			numsCopy := make([]int, len(tst.nums))
			copy(numsCopy, tst.nums)
			sortColors(numsCopy)
			if !slices.Equal(numsCopy, tst.want) {
				t.Errorf("sortColors(%v) = %v, want %v", tst.nums, numsCopy, tst.want)
			}
		})
	}
}
