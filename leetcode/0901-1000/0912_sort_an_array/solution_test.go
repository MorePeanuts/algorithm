package leetcode0912

import (
	"slices"
	"testing"
)

func TestSortArrayExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			"example_1",
			[]int{5, 2, 3, 1},
			[]int{1, 2, 3, 5},
		},
		{
			"example_2",
			[]int{5, 1, 1, 2, 0, 0},
			[]int{0, 0, 1, 1, 2, 5},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			input1 := slices.Clone(tst.nums)
			got1 := sortArray(input1)
			if !slices.Equal(got1, tst.want) {
				t.Errorf("sortArray(%v) = %v, want %v", tst.nums, got1, tst.want)
			}

			input2 := slices.Clone(tst.nums)
			got2 := sortArray2(input2)
			if !slices.Equal(got2, tst.want) {
				t.Errorf("sortArray2(%v) = %v, want %v", tst.nums, got2, tst.want)
			}

			input3 := slices.Clone(tst.nums)
			got3 := sortArray3(input3)
			if !slices.Equal(got3, tst.want) {
				t.Errorf("sortArray3(%v) = %v, want %v", tst.nums, got3, tst.want)
			}
		})
	}
}

func TestSortArray(t *testing.T) {
	const maxVal = 50000
	const minVal = -50000

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"single_element", []int{1}, []int{1}},
		{"two_elements_sorted", []int{1, 2}, []int{1, 2}},
		{"two_elements_reverse", []int{2, 1}, []int{1, 2}},
		{"all_same", []int{3, 3, 3, 3}, []int{3, 3, 3, 3}},
		{"already_sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse_sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"all_negatives", []int{-1, -3, -2, -5}, []int{-5, -3, -2, -1}},
		{"mixed_pos_neg", []int{-1, 3, -2, 5, 0}, []int{-2, -1, 0, 3, 5}},
		{"single_negative", []int{-7}, []int{-7}},
		{"with_zero", []int{0, -1, 1}, []int{-1, 0, 1}},
		{"all_zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"duplicates_at_ends", []int{5, 2, 3, 5}, []int{2, 3, 5, 5}},
		{"max_value", []int{maxVal, 1, 2}, []int{1, 2, maxVal}},
		{"min_value", []int{minVal, 1, 2}, []int{minVal, 1, 2}},
		{"min_max_together", []int{maxVal, minVal, 0}, []int{minVal, 0, maxVal}},
		{"alternating", []int{1, -1, 1, -1, 1}, []int{-1, -1, 1, 1, 1}},
		{"many_duplicates", []int{2, 2, 2, 1, 1, 1, 3, 3, 3}, []int{1, 1, 1, 2, 2, 2, 3, 3, 3}},
		{"two_unique_values", []int{1, 2, 1, 2, 1}, []int{1, 1, 1, 2, 2}},
		{"all_same_except_one", []int{5, 5, 1, 5, 5}, []int{1, 5, 5, 5, 5}},
		{"neg_same_abs_value", []int{-3, 3, -1, 1, 0}, []int{-3, -1, 0, 1, 3}},
		{"large_range", []int{1000, -1000, 500, -500, 0}, []int{-1000, -500, 0, 500, 1000}},
		{"mostly_sorted_one_swap", []int{1, 2, 4, 3, 5}, []int{1, 2, 3, 4, 5}},
		{"insertion_sort_boundary", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"exactly_11_elements", []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{"power_of_two_length", []int{8, 4, 2, 1, 7, 3, 5, 6}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{"odd_length", []int{3, 1, 2}, []int{1, 2, 3}},
		{"even_length", []int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{"single_zero", []int{0}, []int{0}},
		{"descending_all_equal_range", []int{5, 5, 3, 3, 1, 1}, []int{1, 1, 3, 3, 5, 5}},
		{"ascending_all_equal_range", []int{1, 1, 3, 3, 5, 5}, []int{1, 1, 3, 3, 5, 5}},
		{"boundary_only", []int{minVal, maxVal, minVal, maxVal}, []int{minVal, minVal, maxVal, maxVal}},
		{"large_all_same", makeLargeSameSlice(1000, 42), makeLargeSameSlice(1000, 42)},
		{"large_reverse_sorted", makeReverseSlice(1000), makeSortedSlice(1000)},
		{"large_already_sorted", makeSortedSlice(1000), makeSortedSlice(1000)},
		{"leetcode_case", []int{-74, 48, -20, 2, 10, -84, -5, -9, 11, -24, -91, 2, -71, 64, 63, 80, 28, -30, -58, -11, -44, -87, -22, 54, -74, -10, -55, -28, -46, 29, 10, 50, -72, 34, 26, 25, 8, 51, 13, 30, 35, -8, 50, 65, -6, 16, -2, 21, -78, 35, -13, 14, 23, -3, 26, -90, 86, 25, -56, 91, -13, 92, -25, 37, 57, -20, -69, 98, 95, 45, 47, 29, 86, -28, 73, -44, -46, 65, -84, -96, -24, -12, 72, -68, 93, 57, 92, 52, -45, -2, 85, -63, 56, 55, 12, -85, 77, -39}, []int{-96, -91, -90, -87, -85, -84, -84, -78, -74, -74, -72, -71, -69, -68, -63, -58, -56, -55, -46, -46, -45, -44, -44, -39, -30, -28, -28, -25, -24, -24, -22, -20, -20, -13, -13, -12, -11, -10, -9, -8, -6, -5, -3, -2, -2, 2, 2, 8, 10, 10, 11, 12, 13, 14, 16, 21, 23, 25, 25, 26, 26, 28, 29, 29, 30, 34, 35, 35, 37, 45, 47, 48, 50, 50, 51, 52, 54, 55, 56, 57, 57, 63, 64, 65, 65, 72, 73, 77, 80, 85, 86, 86, 91, 92, 92, 93, 95, 98}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			input1 := slices.Clone(tst.nums)
			got1 := sortArray(input1)
			if !slices.Equal(got1, tst.want) {
				t.Errorf("sortArray() = %v, want %v", got1, tst.want)
			}

			input2 := slices.Clone(tst.nums)
			got2 := sortArray2(input2)
			if !slices.Equal(got2, tst.want) {
				t.Errorf("sortArray2() = %v, want %v", got2, tst.want)
			}

			input3 := slices.Clone(tst.nums)
			got3 := sortArray3(input3)
			if !slices.Equal(got3, tst.want) {
				t.Errorf("sortArray3() = %v, want %v", got3, tst.want)
			}
		})
	}
}

// makeSortedSlice returns [1, 2, ..., n].
func makeSortedSlice(n int) []int {
	s := make([]int, n)
	for i := range n {
		s[i] = i + 1
	}
	return s
}

// makeReverseSlice returns [n, n-1, ..., 1].
func makeReverseSlice(n int) []int {
	s := make([]int, n)
	for i := range n {
		s[i] = n - i
	}
	return s
}

// makeLargeSameSlice returns a slice of length n filled with val.
func makeLargeSameSlice(n, val int) []int {
	s := make([]int, n)
	for i := range n {
		s[i] = val
	}
	return s
}
