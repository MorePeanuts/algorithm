package leetcode0739

import (
	"slices"
	"testing"
)

func TestDailyTemperaturesExamples(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			"example_1",
			[]int{73, 74, 75, 71, 69, 72, 76, 73},
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			"example_2",
			[]int{30, 40, 50, 60},
			[]int{1, 1, 1, 0},
		},
		{
			"example_3",
			[]int{30, 60, 90},
			[]int{1, 1, 0},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := dailyTemperatures(tst.temperatures)
			if !slices.Equal(got, tst.want) {
				t.Errorf("dailyTemperatures(%v) = %v, want %v", tst.temperatures, got, tst.want)
			}

			got = dailyTemperatures2(tst.temperatures)
			if !slices.Equal(got, tst.want) {
				t.Errorf("dailyTemperatures(%v) = %v, want %v", tst.temperatures, got, tst.want)
			}
		})
	}
}

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		// 单元素测试
		{"single_element", []int{50}, []int{0}},
		{"single_element_min_temp", []int{30}, []int{0}},
		{"single_element_max_temp", []int{100}, []int{0}},

		// 两元素测试
		{"two_elements_increasing", []int{30, 40}, []int{1, 0}},
		{"two_elements_decreasing", []int{40, 30}, []int{0, 0}},
		{"two_elements_same", []int{50, 50}, []int{0, 0}},

		// 三元素测试
		{"three_increasing", []int{30, 40, 50}, []int{1, 1, 0}},
		{"three_decreasing", []int{50, 40, 30}, []int{0, 0, 0}},
		{"three_same", []int{50, 50, 50}, []int{0, 0, 0}},
		{"three_v_shape", []int{50, 30, 60}, []int{2, 1, 0}},
		{"three_peak", []int{30, 60, 40}, []int{1, 0, 0}},

		// 全相同元素
		{"all_same_small", []int{30, 30, 30, 30}, []int{0, 0, 0, 0}},
		{"all_same_large", []int{100, 100, 100, 100, 100}, []int{0, 0, 0, 0, 0}},

		// 严格递增
		{"strictly_increasing_4", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"strictly_increasing_5", []int{30, 35, 40, 45, 50}, []int{1, 1, 1, 1, 0}},
		{"strictly_increasing_long", []int{30, 31, 32, 33, 34, 35, 36, 37, 38, 39}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 0}},

		// 严格递减
		{"strictly_decreasing_4", []int{60, 50, 40, 30}, []int{0, 0, 0, 0}},
		{"strictly_decreasing_5", []int{100, 90, 80, 70, 60}, []int{0, 0, 0, 0, 0}},
		{"strictly_decreasing_long", []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},

		// 交替模式
		{"alternating_up_down", []int{30, 50, 30, 50, 30}, []int{1, 0, 1, 0, 0}},
		{"alternating_down_up", []int{50, 30, 50, 30, 50}, []int{0, 1, 0, 1, 0}},

		// 边界温度值测试
		{"min_max_boundary", []int{30, 100}, []int{1, 0}},
		{"max_min_boundary", []int{100, 30}, []int{0, 0}},
		{"all_min_temp", []int{30, 30, 30}, []int{0, 0, 0}},
		{"all_max_temp", []int{100, 100, 100}, []int{0, 0, 0}},
		{"min_to_max_range", []int{30, 65, 100}, []int{1, 1, 0}},

		// 复杂模式
		{"valley_then_peak", []int{50, 40, 30, 40, 50, 60}, []int{5, 3, 1, 1, 1, 0}},
		{"peak_then_valley", []int{60, 50, 40, 30, 40, 50}, []int{0, 0, 3, 1, 1, 0}},
		{"multiple_peaks", []int{30, 60, 40, 70, 50, 80}, []int{1, 2, 1, 2, 1, 0}},
		{"multiple_valleys", []int{60, 30, 50, 30, 40, 30}, []int{0, 1, 0, 1, 0, 0}},

		// 跳跃查找测试
		{"skip_one", []int{50, 40, 60}, []int{2, 1, 0}},
		{"skip_two", []int{50, 40, 30, 60}, []int{3, 2, 1, 0}},
		{"skip_many", []int{70, 60, 50, 40, 30, 80}, []int{5, 4, 3, 2, 1, 0}},

		// 等温平台测试
		{"plateau_then_rise", []int{50, 50, 50, 60}, []int{3, 2, 1, 0}},
		{"rise_plateau_rise", []int{40, 50, 50, 50, 60}, []int{1, 3, 2, 1, 0}},
		{"plateau_at_end", []int{40, 50, 60, 60, 60}, []int{1, 1, 0, 0, 0}},

		// 长序列测试
		{"long_plateau", []int{50, 50, 50, 50, 50, 50, 50, 50, 60}, []int{8, 7, 6, 5, 4, 3, 2, 1, 0}},
		{"sawtooth_pattern", []int{30, 40, 30, 40, 30, 40, 30, 40}, []int{1, 0, 1, 0, 1, 0, 1, 0}},

		// 特殊边缘情况
		{"last_is_warmest", []int{30, 40, 50, 60, 100}, []int{1, 1, 1, 1, 0}},
		{"first_is_warmest", []int{100, 90, 80, 70, 60}, []int{0, 0, 0, 0, 0}},
		{"warmest_in_middle", []int{30, 40, 100, 50, 40}, []int{1, 1, 0, 0, 0}},

		// 接近边界长度
		{"length_10", []int{73, 74, 75, 71, 69, 72, 76, 73, 80, 70}, []int{1, 1, 4, 2, 1, 1, 2, 1, 0, 0}},

		// 更复杂的模式
		{"zigzag", []int{30, 50, 40, 60, 50, 70, 60, 80}, []int{1, 2, 1, 2, 1, 2, 1, 0}},
		{"stairs_up_down", []int{30, 40, 50, 60, 50, 40, 30}, []int{1, 1, 1, 0, 0, 0, 0}},
		{"stairs_down_up", []int{60, 50, 40, 30, 40, 50, 60}, []int{0, 5, 3, 1, 1, 1, 0}},

		// 重复子模式
		{"repeated_pattern", []int{30, 40, 30, 40, 30, 40}, []int{1, 0, 1, 0, 1, 0}},
		{"repeated_triple", []int{30, 40, 50, 30, 40, 50, 30, 40, 50}, []int{1, 1, 0, 1, 1, 0, 1, 1, 0}},

		// 最后几个元素相同
		{"same_at_end", []int{40, 50, 60, 60, 60}, []int{1, 1, 0, 0, 0}},
		{"same_at_start", []int{50, 50, 50, 60, 70}, []int{3, 2, 1, 1, 0}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := dailyTemperatures(tst.temperatures)
			if !slices.Equal(got, tst.want) {
				t.Errorf("dailyTemperatures(%v) = %v, want %v", tst.temperatures, got, tst.want)
			}

			got = dailyTemperatures2(tst.temperatures)
			if !slices.Equal(got, tst.want) {
				t.Errorf("dailyTemperatures(%v) = %v, want %v", tst.temperatures, got, tst.want)
			}
		})
	}
}
