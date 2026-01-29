package leetcode0001

import "testing"

// validateTwoSum 验证 twoSum 的结果是否满足题目条件
func validateTwoSum(t *testing.T, fnName string, nums []int, target int, got []int) {
	t.Helper()

	// 检查返回长度
	if len(got) != 2 {
		t.Errorf("%s(%v, %v) returned %v, expected 2 elements", fnName, nums, target, got)
		return
	}

	i, j := got[0], got[1]

	// 检查索引有效性
	if i < 0 || i >= len(nums) || j < 0 || j >= len(nums) {
		t.Errorf("%s(%v, %v) = %v, indices out of range", fnName, nums, target, got)
		return
	}

	// 检查索引不同
	if i == j {
		t.Errorf("%s(%v, %v) = %v, indices must be different", fnName, nums, target, got)
		return
	}

	// 检查和是否等于 target
	if nums[i]+nums[j] != target {
		t.Errorf("%s(%v, %v) = %v, but nums[%d]+nums[%d] = %d+%d = %d != %d",
			fnName, nums, target, got, i, j, nums[i], nums[j], nums[i]+nums[j], target)
	}
}

func TestTwoSumExamples(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
	}{
		{"example_1", []int{2, 7, 11, 15}, 9},
		{"example_2", []int{3, 2, 4}, 6},
		{"example_3", []int{3, 3}, 6},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := twoSum(tst.nums, tst.target)
			validateTwoSum(t, "twoSum", tst.nums, tst.target, got)
		})
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
	}{
		// 最小长度边界
		{"min_length", []int{1, 2}, 3},
		{"min_length_negative", []int{-1, 1}, 0},

		// 相同元素
		{"same_elements", []int{5, 5}, 10},
		{"same_negative_elements", []int{-3, -3}, -6},
		{"zeros", []int{0, 0}, 0},

		// target 为 0
		{"target_zero_opposite", []int{-5, 5}, 0},
		{"target_zero_with_zeros", []int{0, 1, 0}, 0},

		// 负数情况
		{"all_negative", []int{-2, -7, -11, -15}, -9},
		{"negative_target", []int{1, -2, 3, -4}, -1},
		{"mixed_signs", []int{-3, 4, 3, 90}, 0},

		// 大数值边界
		{"large_positive", []int{1000000000, 999999999}, 1999999999},
		{"large_negative", []int{-1000000000, -999999999}, -1999999999},
		{"large_mixed", []int{-1000000000, 1000000000}, 0},

		// 答案在数组不同位置
		{"answer_at_start", []int{1, 2, 5, 6, 7}, 3},
		{"answer_at_end", []int{1, 2, 5, 6, 7}, 13},
		{"answer_at_middle", []int{1, 2, 5, 6, 7}, 7},
		{"answer_span_array", []int{1, 2, 5, 6, 8}, 9},

		// 多元素数组
		{"longer_array", []int{1, 3, 5, 7, 9, 11, 13, 15}, 20},
		{"sequential", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 19},

		// 包含重复元素
		{"duplicates_not_answer", []int{1, 1, 1, 2, 3}, 5},
		{"many_duplicates", []int{2, 2, 2, 2, 3}, 5},

		// 稀疏数组（答案元素距离远）
		{"sparse_answer", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 2}, 3},

		// 特殊 target 值
		{"target_equals_single_element", []int{5, -5, 10}, 5},
		{"target_large_positive", []int{1, 1000000000, 999999999}, 1999999999},
		{"target_large_negative", []int{1, -1000000000, -999999999}, -1999999999},

		// 首尾配对
		{"first_last_pair", []int{10, 5, 3, 7, 20}, 30},

		// 连续配对
		{"consecutive_pair", []int{5, 10, 15, 20, 25}, 35},

		// 单一可能配对
		{"unique_pair", []int{2, 7, 11, 15, 19, 23}, 26},

		// 元素值与索引特殊关系
		{"value_equals_index", []int{0, 1, 2, 3, 4}, 7},

		// 三元素数组
		{"three_elements_first_two", []int{1, 2, 10}, 3},
		{"three_elements_last_two", []int{10, 1, 2}, 3},
		{"three_elements_first_last", []int{1, 10, 2}, 3},

		// 交替正负
		{"alternating_signs", []int{-1, 2, -3, 4, -5, 6}, 1},
		{"alternating_find_zero", []int{-1, 2, -3, 4, -5, 6}, -4},

		// 较长数组中间配对
		{"mid_array_pair", []int{100, 200, 300, 5, 10, 400, 500}, 15},

		// 小数值
		{"small_values", []int{-1, 0, 1}, 0},
		{"tiny_difference", []int{0, 1}, 1},

		// 对称数组
		{"symmetric_array", []int{1, 2, 3, 2, 1}, 4},

		// 递减数组
		{"decreasing_array", []int{10, 8, 6, 4, 2}, 12},

		// 含零的各种情况
		{"zero_with_positive", []int{0, 5, 3}, 5},
		{"zero_with_negative", []int{0, -5, 3}, -5},
		{"multiple_zeros", []int{0, 1, 0, 2}, 0},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := twoSum(tst.nums, tst.target)
			validateTwoSum(t, "twoSum", tst.nums, tst.target, got)
		})
	}
}
