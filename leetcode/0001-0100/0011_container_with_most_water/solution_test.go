package leetcode0011

import "testing"

func TestMaxAreaExamples(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			"example_1",
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			"example_2",
			[]int{1, 1},
			1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := maxArea(tst.height)
			if got != tst.want {
				t.Errorf("maxArea(%v) = %v, want %v", tst.height, got, tst.want)
			}
		})
	}
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		// 最小长度测试 (n >= 2)
		{"min_length_same_height", []int{5, 5}, 5},
		{"min_length_diff_height", []int{1, 2}, 1},
		{"min_length_zero_first", []int{0, 5}, 0},
		{"min_length_zero_second", []int{5, 0}, 0},
		{"min_length_both_zero", []int{0, 0}, 0},

		// 三元素
		{"three_elements_peak", []int{1, 10, 1}, 2},
		{"three_elements_ascending", []int{1, 2, 3}, 2},
		{"three_elements_descending", []int{3, 2, 1}, 2},
		{"three_elements_same", []int{5, 5, 5}, 10},

		// 全相同高度
		{"all_same_4", []int{3, 3, 3, 3}, 9},
		{"all_same_5", []int{7, 7, 7, 7, 7}, 28},
		{"all_same_10", []int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, 36},

		// 全零
		{"all_zeros", []int{0, 0, 0, 0}, 0},

		// 单边为零
		{"first_zero_others_same", []int{0, 5, 5, 5}, 10},
		{"last_zero_others_same", []int{5, 5, 5, 0}, 10},

		// 递增序列
		{"ascending_4", []int{1, 2, 3, 4}, 4},
		{"ascending_5", []int{1, 2, 3, 4, 5}, 6},
		{"ascending_6", []int{1, 2, 3, 4, 5, 6}, 9},

		// 递减序列
		{"descending_4", []int{4, 3, 2, 1}, 4},
		{"descending_5", []int{5, 4, 3, 2, 1}, 6},
		{"descending_6", []int{6, 5, 4, 3, 2, 1}, 9},

		// 两端高中间低
		{"v_shape", []int{10, 1, 10}, 20},
		{"v_shape_wider", []int{10, 1, 1, 1, 10}, 40},
		{"v_shape_asymmetric", []int{10, 1, 1, 5}, 15},

		// 中间高两端低
		{"peak_shape", []int{1, 10, 1}, 2},
		{"peak_shape_wider", []int{1, 1, 10, 1, 1}, 4},
		{"multiple_peaks", []int{1, 5, 1, 5, 1}, 10},

		// 最大值在不同位置
		{"max_at_ends", []int{100, 1, 1, 1, 100}, 400},
		{"max_at_left", []int{100, 50, 50, 50}, 150},
		{"max_at_right", []int{50, 50, 50, 100}, 150},
		{"max_at_middle", []int{1, 100, 100, 1}, 100},

		// 交替高低
		{"alternating", []int{1, 10, 1, 10, 1, 10}, 40},
		{"alternating_reverse", []int{10, 1, 10, 1, 10, 1}, 40},

		// 边界值测试 (height[i] <= 10^4)
		{"max_height_value", []int{10000, 10000}, 10000},
		{"max_height_spread", []int{10000, 1, 1, 1, 1, 10000}, 50000},
		{"mixed_max_min", []int{10000, 0, 10000}, 20000},

		// 较长数组
		{"long_ascending", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 25},
		{"long_descending", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 25},
		{"long_same", []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, 45},
		{"long_v_shape", []int{10, 5, 3, 1, 1, 1, 3, 5, 10}, 80},

		// 特殊模式
		{"stairs_up_down", []int{1, 2, 3, 4, 5, 4, 3, 2, 1}, 12},
		{"plateau_middle", []int{1, 5, 5, 5, 5, 5, 1}, 20},
		{"double_peak", []int{1, 10, 1, 1, 1, 10, 1}, 40},

		// 极端差异
		{"extreme_diff", []int{1, 10000}, 1},
		{"extreme_diff_wider", []int{1, 1, 1, 1, 10000}, 4},
		{"extreme_diff_reverse", []int{10000, 1, 1, 1, 1}, 4},

		// 随机模式
		{"random_1", []int{3, 1, 2, 4, 5}, 12},
		{"random_2", []int{2, 3, 10, 5, 7, 8, 9}, 36},
		{"random_3", []int{4, 3, 2, 1, 4}, 16},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := maxArea(tst.height)
			if got != tst.want {
				t.Errorf("maxArea(%v) = %v, want %v", tst.height, got, tst.want)
			}
		})
	}
}
