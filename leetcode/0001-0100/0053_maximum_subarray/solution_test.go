package leetcode0053

import "testing"

func TestMaxSubArrayExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			"example_1",
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			"example_2",
			[]int{1},
			1,
		},
		{
			"example_3",
			[]int{5, 4, -1, 7, 8},
			23,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			if got := maxSubArray(tst.nums); got != tst.want {
				t.Errorf("maxSubArray(%v) = %v, want %v", tst.nums, got, tst.want)
			}
			if got := maxSubArray2(tst.nums); got != tst.want {
				t.Errorf("maxSubArray2(%v) = %v, want %v", tst.nums, got, tst.want)
			}
			if got := maxSubArray3(tst.nums); got != tst.want {
				t.Errorf("maxSubArray3(%v) = %v, want %v", tst.nums, got, tst.want)
			}
			if got := maxSubArray4(tst.nums); got != tst.want {
				t.Errorf("maxSubArray4(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// 单元素测试
		{"single_positive", []int{5}, 5},
		{"single_negative", []int{-5}, -5},
		{"single_zero", []int{0}, 0},

		// 全正数
		{"all_positive", []int{1, 2, 3, 4, 5}, 15},
		{"all_positive_two", []int{3, 7}, 10},

		// 全负数
		{"all_negative", []int{-1, -2, -3, -4, -5}, -1},
		{"all_negative_two", []int{-5, -3}, -3},
		{"all_negative_three", []int{-2, -1}, -1},

		// 全零
		{"all_zeros", []int{0, 0, 0}, 0},
		{"zeros_with_positive", []int{0, 0, 1, 0, 0}, 1},
		{"zeros_with_negative", []int{0, 0, -1, 0, 0}, 0},

		// 交替正负
		{"alternating_pos_neg", []int{1, -1, 1, -1, 1}, 1},
		{"alternating_neg_pos", []int{-1, 2, -1, 2, -1}, 3},
		{"alternating_large", []int{10, -5, 10, -5, 10}, 20},

		// 最大子数组在开头
		{"max_at_start", []int{5, 4, -10, 1, 2}, 9},
		{"max_at_start_single", []int{100, -200, 1, 2, 3}, 100},

		// 最大子数组在结尾
		{"max_at_end", []int{-1, -2, 5, 4, 3}, 12},
		{"max_at_end_single", []int{-10, -20, 100}, 100},

		// 最大子数组在中间
		{"max_at_middle", []int{-5, 3, 4, 5, -10}, 12},
		{"max_at_middle_with_neg", []int{-2, 1, 3, -1, 2, -5}, 5},

		// 整个数组是最大子数组
		{"entire_array", []int{1, 2, 3, 4}, 10},
		{"entire_array_with_neg", []int{2, -1, 3, -1, 4}, 7},

		// 包含大数值
		{"large_values", []int{10000, -10000, 10000}, 10000},
		{"large_positive", []int{10000, 10000}, 20000},
		{"large_negative", []int{-10000, -10000}, -10000},

		// 边界值
		{"min_value", []int{-10000}, -10000},
		{"max_value", []int{10000}, 10000},
		{"min_max_values", []int{-10000, 10000}, 10000},
		{"max_min_values", []int{10000, -10000}, 10000},

		// 连续相同元素
		{"all_same_positive", []int{3, 3, 3, 3}, 12},
		{"all_same_negative", []int{-2, -2, -2}, -2},

		// 子数组需要跳过负数
		{"skip_negative", []int{1, 2, -10, 4, 5}, 9},
		{"skip_large_negative", []int{5, -100, 6}, 6},

		// 复杂混合
		{"complex_1", []int{-2, -3, 4, -1, -2, 1, 5, -3}, 7},
		{"complex_2", []int{8, -19, 5, -4, 20}, 21},
		{"complex_3", []int{-1, 3, -2, 5, -6, 2, 4}, 6},
		{"complex_4", []int{1, -1, 1, -1, 1, -1, 1}, 1},

		// 两个独立的正数区间
		{"two_regions", []int{5, -100, 6}, 6},
		{"two_regions_first_larger", []int{10, -100, 5}, 10},
		{"two_equal_regions", []int{5, -100, 5}, 5},

		// 长数组测试
		{"longer_array_1", []int{1, 2, 3, -6, 1, 2, 3, 4, -10, 5}, 10},
		{"longer_array_2", []int{-1, -2, 1, 2, 3, 4, 5, -1, -2}, 15},

		// 前缀和后缀测试
		{"prefix_sum", []int{3, -1, -1, -1, 10}, 10},
		{"suffix_sum", []int{10, -1, -1, -1, 3}, 10},

		// 负数中含零
		{"negatives_with_zero", []int{-3, -2, 0, -1}, 0},

		// 特殊模式
		{"valley_pattern", []int{5, -10, 5}, 5},
		{"peak_pattern", []int{-5, 10, -5}, 10},
		{"double_peak", []int{5, -1, 5, -1, 5}, 13},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			if got := maxSubArray(tst.nums); got != tst.want {
				t.Errorf("maxSubArray(%v) = %v, want %v", tst.nums, got, tst.want)
			}
			if got := maxSubArray2(tst.nums); got != tst.want {
				t.Errorf("maxSubArray2(%v) = %v, want %v", tst.nums, got, tst.want)
			}
			if got := maxSubArray3(tst.nums); got != tst.want {
				t.Errorf("maxSubArray3(%v) = %v, want %v", tst.nums, got, tst.want)
			}
			if got := maxSubArray4(tst.nums); got != tst.want {
				t.Errorf("maxSubArray4(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}
