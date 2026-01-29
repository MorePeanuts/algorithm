package leetcode0217

import "testing"

func TestContainsDuplicateExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			"example_1",
			[]int{1, 2, 3, 1},
			true,
		},
		{
			"example_2",
			[]int{1, 2, 3, 4},
			false,
		},
		{
			"example_3",
			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			true,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := containsDuplicate(tst.nums)
			if got != tst.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		// 单元素测试
		{"single_element", []int{1}, false},
		{"single_zero", []int{0}, false},
		{"single_negative", []int{-1}, false},
		{"single_max_value", []int{1000000000}, false},
		{"single_min_value", []int{-1000000000}, false},

		// 两元素测试
		{"two_same", []int{1, 1}, true},
		{"two_different", []int{1, 2}, false},
		{"two_opposite", []int{-1, 1}, false},
		{"two_zeros", []int{0, 0}, true},

		// 全相同元素
		{"all_same_small", []int{5, 5, 5}, true},
		{"all_same_zeros", []int{0, 0, 0, 0}, true},
		{"all_same_negative", []int{-1, -1, -1}, true},

		// 全不同元素
		{"all_different_small", []int{1, 2, 3, 4, 5}, false},
		{"all_different_negative", []int{-5, -4, -3, -2, -1}, false},
		{"all_different_mixed", []int{-2, -1, 0, 1, 2}, false},

		// 边界值测试
		{"max_values_same", []int{1000000000, 1000000000}, true},
		{"min_values_same", []int{-1000000000, -1000000000}, true},
		{"max_and_min", []int{-1000000000, 1000000000}, false},
		{"extremes_with_zero", []int{-1000000000, 0, 1000000000}, false},
		{"extremes_with_duplicate", []int{-1000000000, 0, 1000000000, 0}, true},

		// 重复在首位
		{"duplicate_at_start", []int{1, 1, 2, 3, 4}, true},
		{"duplicate_at_end", []int{1, 2, 3, 4, 4}, true},
		{"duplicate_at_boundaries", []int{1, 2, 3, 4, 1}, true},

		// 重复在中间
		{"duplicate_in_middle", []int{1, 2, 2, 3, 4}, true},
		{"multiple_duplicates", []int{1, 2, 1, 2, 1}, true},

		// 连续数字
		{"consecutive_ascending", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
		{"consecutive_descending", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
		{"consecutive_with_repeat", []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}, true},

		// 特殊模式
		{"alternating", []int{1, 2, 1, 2, 1, 2}, true},
		{"almost_unique", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}, true},
		{"unique_large_range", []int{-100, -50, 0, 50, 100}, false},

		// 零相关
		{"zeros_mixed", []int{0, 1, 0, 2}, true},
		{"single_zero_in_array", []int{1, 2, 0, 3, 4}, false},
		{"multiple_zeros", []int{0, 0, 0}, true},

		// 负数相关
		{"negative_duplicate", []int{-5, -3, -5, -1}, true},
		{"negative_all_different", []int{-10, -20, -30, -40}, false},
		{"negative_positive_same_abs", []int{-5, 5, -5}, true},

		// 较大数组
		{"medium_no_duplicate", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, false},
		{"medium_with_duplicate", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 1}, true},
		{"medium_last_duplicate", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 20}, true},

		// 乱序测试
		{"shuffled_no_duplicate", []int{7, 3, 9, 1, 5, 2, 8, 4, 6, 10}, false},
		{"shuffled_with_duplicate", []int{7, 3, 9, 1, 5, 2, 8, 4, 6, 3}, true},

		// 大数值范围
		{"sparse_values", []int{1, 1000, 1000000, 1000000000}, false},
		{"sparse_with_duplicate", []int{1, 1000, 1000000, 1000000000, 1000}, true},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := containsDuplicate(tst.nums)
			if got != tst.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}
