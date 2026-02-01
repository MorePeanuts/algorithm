package leetcode0238

import (
	"slices"
	"testing"
)

func TestProductExceptSelfExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			"example_1",
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			"example_2",
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := productExceptSelf(tst.nums)
			if !slices.Equal(got, tst.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tst.nums, got, tst.want)
			}

			got = productExceptSelf2(tst.nums)
			if !slices.Equal(got, tst.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tst.nums, got, tst.want)
			}

			got = productExceptSelf3(tst.nums)
			if !slices.Equal(got, tst.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		// 最小长度测试
		{"min_length_positive", []int{1, 2}, []int{2, 1}},
		{"min_length_negative", []int{-1, -2}, []int{-2, -1}},
		{"min_length_mixed", []int{-3, 5}, []int{5, -3}},
		{"min_length_with_zero", []int{0, 5}, []int{5, 0}},
		{"min_length_both_zero", []int{0, 0}, []int{0, 0}},

		// 全相同元素
		{"all_ones", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"all_twos", []int{2, 2, 2}, []int{4, 4, 4}},
		{"all_negative_ones", []int{-1, -1, -1, -1}, []int{-1, -1, -1, -1}},
		{"all_zeros", []int{0, 0, 0}, []int{0, 0, 0}},

		// 包含零的情况
		{"single_zero", []int{1, 2, 0, 4}, []int{0, 0, 8, 0}},
		{"two_zeros", []int{1, 0, 3, 0}, []int{0, 0, 0, 0}},
		{"zero_at_start", []int{0, 1, 2, 3}, []int{6, 0, 0, 0}},
		{"zero_at_end", []int{1, 2, 3, 0}, []int{0, 0, 0, 6}},
		{"zeros_adjacent", []int{1, 0, 0, 4}, []int{0, 0, 0, 0}},

		// 包含负数的情况
		{"single_negative", []int{1, -2, 3, 4}, []int{-24, 12, -8, -6}},
		{"two_negatives", []int{-1, -2, 3, 4}, []int{-24, -12, 8, 6}},
		{"all_negatives", []int{-1, -2, -3, -4}, []int{-24, -12, -8, -6}},
		{"negative_with_zero", []int{-1, 0, -3, -4}, []int{0, -12, 0, 0}},

		// 边界值测试 (-30 到 30)
		{"max_value", []int{30, 1, 1}, []int{1, 30, 30}},
		{"min_value", []int{-30, 1, 1}, []int{1, -30, -30}},
		{"max_min_values", []int{30, -30, 1}, []int{-30, 30, -900}},
		{"extreme_values", []int{30, 30, -30, -30}, []int{27000, 27000, -27000, -27000}},

		// 三元素测试
		{"three_elements_positive", []int{1, 2, 3}, []int{6, 3, 2}},
		{"three_elements_mixed", []int{-1, 2, -3}, []int{-6, 3, -2}},
		{"three_elements_with_zero", []int{1, 0, 3}, []int{0, 3, 0}},

		// 递增/递减序列
		{"increasing_sequence", []int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{"decreasing_sequence", []int{5, 4, 3, 2, 1}, []int{24, 30, 40, 60, 120}},

		// 对称序列
		{"symmetric_positive", []int{1, 2, 3, 2, 1}, []int{12, 6, 4, 6, 12}},
		{"symmetric_with_zero", []int{1, 2, 0, 2, 1}, []int{0, 0, 4, 0, 0}},

		// 特殊模式
		{"alternating_sign", []int{1, -1, 1, -1, 1}, []int{1, -1, 1, -1, 1}},
		{"one_and_minus_one", []int{1, -1}, []int{-1, 1}},
		{"powers_of_two", []int{2, 4, 8}, []int{32, 16, 8}},

		// 稍大规模测试
		{"six_elements", []int{1, 2, 3, 4, 5, 6}, []int{720, 360, 240, 180, 144, 120}},
		{"seven_elements_with_zero", []int{1, 2, 3, 0, 5, 6, 7}, []int{0, 0, 0, 1260, 0, 0, 0}},

		// 乘积为1的情况
		{"product_is_one", []int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
		{"product_is_negative_one", []int{-1, 1, 1, 1}, []int{1, -1, -1, -1}},

		// 较大乘积测试（在32位范围内）
		{"large_product", []int{10, 10, 10, 10}, []int{1000, 1000, 1000, 1000}},
		{"mixed_large_values", []int{30, 20, 10, 1}, []int{200, 300, 600, 6000}},

		// 单个1或-1的测试
		{"single_one_rest_same", []int{1, 5, 5, 5}, []int{125, 25, 25, 25}},
		{"single_negative_one", []int{-1, 5, 5, 5}, []int{125, -25, -25, -25}},

		// 更多边界情况
		{"two_large_one_small", []int{30, 30, 1}, []int{30, 30, 900}},
		{"negative_large_values", []int{-30, -30, -1}, []int{30, 30, 900}},
		{"mixed_extremes", []int{-30, 0, 30}, []int{0, -900, 0}},

		// 连续相同值
		{"five_twos", []int{2, 2, 2, 2, 2}, []int{16, 16, 16, 16, 16}},
		{"five_threes", []int{3, 3, 3, 3, 3}, []int{81, 81, 81, 81, 81}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := productExceptSelf(tst.nums)
			if !slices.Equal(got, tst.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tst.nums, got, tst.want)
			}

			got = productExceptSelf2(tst.nums)
			if !slices.Equal(got, tst.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tst.nums, got, tst.want)
			}

			got = productExceptSelf3(tst.nums)
			if !slices.Equal(got, tst.want) {
				t.Errorf("productExceptSelf(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}
