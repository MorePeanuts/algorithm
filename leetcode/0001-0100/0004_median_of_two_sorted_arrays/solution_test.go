package leetcode0004

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestFindMedianSortedArraysExamples(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			"example_1",
			[]int{1, 3},
			[]int{2},
			2.0,
		},
		{
			"example_2",
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := findMedianSortedArrays(tst.nums1, tst.nums2)
			if !utils.Float64AlmostEqual(got, tst.want) {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v, want %v", tst.nums1, tst.nums2, got, tst.want)
			}
		})
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{"nums1_empty", []int{}, []int{1}, 1.0},
		{"nums2_empty", []int{1}, []int{}, 1.0},
		{"single_element_each", []int{1}, []int{2}, 1.5},
		{"both_single_same", []int{5}, []int{5}, 5.0},
		{"nums1_all_before_nums2", []int{1, 2, 3}, []int{4, 5, 6}, 3.5},
		{"nums2_all_before_nums1", []int{4, 5, 6}, []int{1, 2, 3}, 3.5},
		{"interleaved_odd_total", []int{1, 3, 5}, []int{2, 4, 6, 7}, 4.0},
		{"interleaved_even_total", []int{1, 4, 5}, []int{2, 3, 6}, 3.5},
		{"one_element_many", []int{1}, []int{2, 3, 4, 5, 6}, 3.5},
		{"many_elements_one", []int{2, 3, 4, 5, 6}, []int{1}, 3.5},
		{"duplicate_values", []int{1, 1, 1}, []int{1, 1, 1}, 1.0},
		{"negative_numbers", []int{-5, -3, -1}, []int{-4, -2, 0}, -2.5},
		{"mixed_positive_negative", []int{-10, -5, 0, 5, 10}, []int{-7, -3, 3, 7}, 0.0},
		{"large_values", []int{1000000, 2000000}, []int{1500000, 2500000}, 1750000.0},
		{"min_values", []int{-1000000, -999999}, []int{-999998, -999997}, -999998.5},
		{"nums1_len_1000", make([]int, 1000), []int{0}, 0.0},
		{"nums2_len_1000", []int{0}, make([]int, 1000), 0.0},
		{"nums1_len_2_nums2_len_3", []int{1, 5}, []int{2, 3, 4}, 3.0},
		{"nums1_len_3_nums2_len_2", []int{2, 3, 4}, []int{1, 5}, 3.0},
		{"one_at_start_one_at_end", []int{1}, []int{2, 3, 4, 5, 100}, 3.5},
		{"even_length_overlap", []int{2, 4, 6, 8}, []int{1, 3, 5, 7, 9, 10}, 5.5},
		{"odd_length_overlap", []int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9}, 5.5},
		{"nums1_nums2_adjacent", []int{1, 2, 3}, []int{4, 5}, 3.0},
		{"nums2_nums1_adjacent", []int{4, 5}, []int{1, 2, 3}, 3.0},
		{"all_nums1_smaller_even", []int{1, 2, 3, 4}, []int{5, 6, 7, 8}, 4.5},
		{"all_nums1_smaller_odd", []int{1, 2, 3}, []int{4, 5, 6, 7}, 4.0},
		{"alternating_values", []int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}, 5.5},
		{"same_array", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 3.0},
		{"one_greater_one_less", []int{10, 20}, []int{1, 30}, 15.0},
		{"nums1_single_min", []int{-1000000}, []int{0, 1, 2, 3}, 1.0},
		{"nums2_single_max", []int{0, 1, 2, 3}, []int{1000000}, 2.0},
		{"leetcode_example", []int{-10, -9, -8}, []int{1, 2}, -8.0},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := findMedianSortedArrays(tst.nums1, tst.nums2)
			if !utils.Float64AlmostEqual(got, tst.want) {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v, want %v", tst.nums1, tst.nums2, got, tst.want)
			}
		})
	}
}
