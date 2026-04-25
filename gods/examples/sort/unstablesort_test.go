package sort

import (
	"testing"
)

func TestShellSortExamples(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"single_element", []int{5}},
		{"two_elements", []int{2, 1}},
		{"already_sorted", []int{1, 2, 3, 4, 5}},
		{"reverse_sorted", []int{5, 4, 3, 2, 1}},
		{"with_duplicates", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}},
		{"all_zeros", []int{0, 0, 0, 0}},
		{"with_zero", []int{0, 5, 3, 0, 2}},
		{"consecutive", []int{10, 7, 8, 9, 1, 5}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			original := make([]int, len(tst.arr))
			copy(original, tst.arr)
			arr := make([]int, len(tst.arr))
			copy(arr, tst.arr)
			ShellSort(arr)
			validateSortedInPlace(t, "ShellSort", original, arr)
		})
	}
}

func TestShellSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"empty", []int{}},
		{"min_length", []int{1}},
		{"min_length_zero", []int{0}},
		{"same_elements", []int{5, 5, 5, 5}},
		{"same_large", []int{100, 100, 100}},
		{"only_zeros", []int{0, 0, 0}},
		{"negative_numbers", []int{-5, -3, -7, -1}},
		{"mix_positive_negative", []int{3, -1, 4, -2, 5, -3}},
		{"consecutive_asc", []int{0, 1, 2, 3, 4, 5}},
		{"consecutive_desc", []int{5, 4, 3, 2, 1, 0}},
		{"consecutive_mixed", []int{3, 0, 5, 2, 4, 1}},
		{"pattern_1212", []int{1, 2, 1, 2, 1, 2}},
		{"pattern_123123", []int{1, 2, 3, 1, 2, 3}},
		{"single_peak", []int{1, 3, 5, 7, 5, 3, 1}},
		{"valley", []int{7, 5, 3, 1, 3, 5, 7}},
		{"starts_with_zero", []int{0, 10, 20, 30}},
		{"ends_with_zero", []int{30, 20, 10, 0}},
		{"zero_in_middle", []int{10, 20, 0, 30, 40}},
		{"large_numbers", []int{1000, 500, 100, 50, 10, 0}},
		{"hundreds", []int{300, 100, 200, 500, 400}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			original := make([]int, len(tst.arr))
			copy(original, tst.arr)
			arr := make([]int, len(tst.arr))
			copy(arr, tst.arr)
			ShellSort(arr)
			validateSortedInPlace(t, "ShellSort", original, arr)
		})
	}
}

func TestHeapSortExamples(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"single_element", []int{5}},
		{"two_elements", []int{2, 1}},
		{"already_sorted", []int{1, 2, 3, 4, 5}},
		{"reverse_sorted", []int{5, 4, 3, 2, 1}},
		{"with_duplicates", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}},
		{"all_zeros", []int{0, 0, 0, 0}},
		{"with_zero", []int{0, 5, 3, 0, 2}},
		{"consecutive", []int{10, 7, 8, 9, 1, 5}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			original := make([]int, len(tst.arr))
			copy(original, tst.arr)
			arr := make([]int, len(tst.arr))
			copy(arr, tst.arr)
			HeapSort(arr)
			validateSortedInPlace(t, "HeapSort", original, arr)
		})
	}
}

func TestHeapSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"empty", []int{}},
		{"min_length", []int{1}},
		{"min_length_zero", []int{0}},
		{"same_elements", []int{5, 5, 5, 5}},
		{"same_large", []int{100, 100, 100}},
		{"only_zeros", []int{0, 0, 0}},
		{"negative_numbers", []int{-5, -3, -7, -1}},
		{"mix_positive_negative", []int{3, -1, 4, -2, 5, -3}},
		{"consecutive_asc", []int{0, 1, 2, 3, 4, 5}},
		{"consecutive_desc", []int{5, 4, 3, 2, 1, 0}},
		{"consecutive_mixed", []int{3, 0, 5, 2, 4, 1}},
		{"pattern_1212", []int{1, 2, 1, 2, 1, 2}},
		{"pattern_123123", []int{1, 2, 3, 1, 2, 3}},
		{"single_peak", []int{1, 3, 5, 7, 5, 3, 1}},
		{"valley", []int{7, 5, 3, 1, 3, 5, 7}},
		{"starts_with_zero", []int{0, 10, 20, 30}},
		{"ends_with_zero", []int{30, 20, 10, 0}},
		{"zero_in_middle", []int{10, 20, 0, 30, 40}},
		{"large_numbers", []int{1000, 500, 100, 50, 10, 0}},
		{"hundreds", []int{300, 100, 200, 500, 400}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			original := make([]int, len(tst.arr))
			copy(original, tst.arr)
			arr := make([]int, len(tst.arr))
			copy(arr, tst.arr)
			HeapSort(arr)
			validateSortedInPlace(t, "HeapSort", original, arr)
		})
	}
}

func TestQuickSortExamples(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"single_element", []int{5}},
		{"two_elements", []int{2, 1}},
		{"already_sorted", []int{1, 2, 3, 4, 5}},
		{"reverse_sorted", []int{5, 4, 3, 2, 1}},
		{"with_duplicates", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}},
		{"all_zeros", []int{0, 0, 0, 0}},
		{"with_zero", []int{0, 5, 3, 0, 2}},
		{"consecutive", []int{10, 7, 8, 9, 1, 5}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			original := make([]int, len(tst.arr))
			copy(original, tst.arr)
			arr := make([]int, len(tst.arr))
			copy(arr, tst.arr)
			QuickSort(arr)
			validateSortedInPlace(t, "QuickSort", original, arr)
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"empty", []int{}},
		{"min_length", []int{1}},
		{"min_length_zero", []int{0}},
		{"same_elements", []int{5, 5, 5, 5}},
		{"same_large", []int{100, 100, 100}},
		{"only_zeros", []int{0, 0, 0}},
		{"negative_numbers", []int{-5, -3, -7, -1}},
		{"mix_positive_negative", []int{3, -1, 4, -2, 5, -3}},
		{"consecutive_asc", []int{0, 1, 2, 3, 4, 5}},
		{"consecutive_desc", []int{5, 4, 3, 2, 1, 0}},
		{"consecutive_mixed", []int{3, 0, 5, 2, 4, 1}},
		{"pattern_1212", []int{1, 2, 1, 2, 1, 2}},
		{"pattern_123123", []int{1, 2, 3, 1, 2, 3}},
		{"single_peak", []int{1, 3, 5, 7, 5, 3, 1}},
		{"valley", []int{7, 5, 3, 1, 3, 5, 7}},
		{"starts_with_zero", []int{0, 10, 20, 30}},
		{"ends_with_zero", []int{30, 20, 10, 0}},
		{"zero_in_middle", []int{10, 20, 0, 30, 40}},
		{"large_numbers", []int{1000, 500, 100, 50, 10, 0}},
		{"hundreds", []int{300, 100, 200, 500, 400}},
		{"below_cutoff", []int{5, 4, 3, 2, 1, 0, 9, 8, 7}},
		{"at_cutoff", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{"just_above_cutoff", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			original := make([]int, len(tst.arr))
			copy(original, tst.arr)
			arr := make([]int, len(tst.arr))
			copy(arr, tst.arr)
			QuickSort(arr)
			validateSortedInPlace(t, "QuickSort", original, arr)
		})
	}
}

func TestQuickSort3WayExamples(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"single_element", []int{5}},
		{"two_elements", []int{2, 1}},
		{"already_sorted", []int{1, 2, 3, 4, 5}},
		{"reverse_sorted", []int{5, 4, 3, 2, 1}},
		{"with_duplicates", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}},
		{"all_zeros", []int{0, 0, 0, 0}},
		{"with_zero", []int{0, 5, 3, 0, 2}},
		{"consecutive", []int{10, 7, 8, 9, 1, 5}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			original := make([]int, len(tst.arr))
			copy(original, tst.arr)
			arr := make([]int, len(tst.arr))
			copy(arr, tst.arr)
			QuickSort3Way(arr)
			validateSortedInPlace(t, "QuickSort3Way", original, arr)
		})
	}
}

func TestQuickSort3Way(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"empty", []int{}},
		{"min_length", []int{1}},
		{"min_length_zero", []int{0}},
		{"same_elements", []int{5, 5, 5, 5}},
		{"same_large", []int{100, 100, 100}},
		{"only_zeros", []int{0, 0, 0}},
		{"negative_numbers", []int{-5, -3, -7, -1}},
		{"mix_positive_negative", []int{3, -1, 4, -2, 5, -3}},
		{"consecutive_asc", []int{0, 1, 2, 3, 4, 5}},
		{"consecutive_desc", []int{5, 4, 3, 2, 1, 0}},
		{"consecutive_mixed", []int{3, 0, 5, 2, 4, 1}},
		{"pattern_1212", []int{1, 2, 1, 2, 1, 2}},
		{"pattern_123123", []int{1, 2, 3, 1, 2, 3}},
		{"single_peak", []int{1, 3, 5, 7, 5, 3, 1}},
		{"valley", []int{7, 5, 3, 1, 3, 5, 7}},
		{"starts_with_zero", []int{0, 10, 20, 30}},
		{"ends_with_zero", []int{30, 20, 10, 0}},
		{"zero_in_middle", []int{10, 20, 0, 30, 40}},
		{"large_numbers", []int{1000, 500, 100, 50, 10, 0}},
		{"hundreds", []int{300, 100, 200, 500, 400}},
		{"all_same", []int{7, 7, 7, 7, 7, 7, 7}},
		{"many_duplicates", []int{2, 1, 2, 1, 2, 1, 2, 1, 2, 1}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			original := make([]int, len(tst.arr))
			copy(original, tst.arr)
			arr := make([]int, len(tst.arr))
			copy(arr, tst.arr)
			QuickSort3Way(arr)
			validateSortedInPlace(t, "QuickSort3Way", original, arr)
		})
	}
}