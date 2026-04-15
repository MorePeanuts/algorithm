package leetcode0704

import "testing"

func TestSearchExamples(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			"example_1",
			[]int{-1, 0, 3, 5, 9, 12},
			9,
			4,
		},
		{
			"example_2",
			[]int{-1, 0, 3, 5, 9, 12},
			2,
			-1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := search(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %d) = %d, want %d", tst.nums, tst.target, got, tst.want)
			}

			got = search2(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %d) = %d, want %d", tst.nums, tst.target, got, tst.want)
			}

			got = search3(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %d) = %d, want %d", tst.nums, tst.target, got, tst.want)
			}

			got = search4(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %d) = %d, want %d", tst.nums, tst.target, got, tst.want)
			}

			got = search5(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %d) = %d, want %d", tst.nums, tst.target, got, tst.want)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		// 单元素数组
		{"single_element_found", []int{5}, 5, 0},
		{"single_element_not_found_smaller", []int{5}, 3, -1},
		{"single_element_not_found_larger", []int{5}, 7, -1},

		// 双元素数组
		{"two_elements_found_first", []int{1, 3}, 1, 0},
		{"two_elements_found_second", []int{1, 3}, 3, 1},
		{"two_elements_not_found_smaller", []int{1, 3}, 0, -1},
		{"two_elements_not_found_middle", []int{1, 3}, 2, -1},
		{"two_elements_not_found_larger", []int{1, 3}, 5, -1},

		// 三元素数组
		{"three_elements_found_first", []int{1, 2, 3}, 1, 0},
		{"three_elements_found_middle", []int{1, 2, 3}, 2, 1},
		{"three_elements_found_last", []int{1, 2, 3}, 3, 2},
		{"three_elements_not_found", []int{1, 2, 3}, 4, -1},

		// 边界值测试（根据提示约束）
		{"min_value_found", []int{-9999, 0, 100}, -9999, 0},
		{"max_value_found", []int{-100, 0, 9999}, 9999, 2},
		{"target_smaller_than_min", []int{-9999, 0, 100}, -10000, -1},
		{"target_larger_than_max", []int{-100, 0, 9999}, 10000, -1},

		// 负数测试
		{"all_negative_found", []int{-100, -50, -10, -1}, -50, 1},
		{"all_negative_not_found", []int{-100, -50, -10, -1}, -25, -1},
		{"negative_to_positive_found_zero", []int{-5, -3, 0, 2, 4}, 0, 2},
		{"negative_to_positive_found_negative", []int{-5, -3, 0, 2, 4}, -3, 1},
		{"negative_to_positive_found_positive", []int{-5, -3, 0, 2, 4}, 4, 4},

		// 连续整数
		{"consecutive_found_first", []int{1, 2, 3, 4, 5}, 1, 0},
		{"consecutive_found_last", []int{1, 2, 3, 4, 5}, 5, 4},
		{"consecutive_found_middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"consecutive_not_found_gap", []int{1, 3, 5, 7, 9}, 4, -1},

		// 偶数长度数组
		{"even_length_found_left_half", []int{1, 2, 3, 4, 5, 6}, 2, 1},
		{"even_length_found_right_half", []int{1, 2, 3, 4, 5, 6}, 5, 4},
		{"even_length_not_found", []int{1, 2, 3, 4, 5, 6}, 7, -1},

		// 奇数长度数组
		{"odd_length_found_exact_middle", []int{1, 2, 3, 4, 5, 6, 7}, 4, 3},
		{"odd_length_found_left", []int{1, 2, 3, 4, 5, 6, 7}, 2, 1},
		{"odd_length_found_right", []int{1, 2, 3, 4, 5, 6, 7}, 6, 5},

		// 大间隔数组
		{"large_gaps_found", []int{1, 100, 1000, 10000}, 100, 1},
		{"large_gaps_not_found", []int{1, 100, 1000, 10000}, 500, -1},

		// 较大规模数据
		{"medium_array_found_start", generateSequence(0, 99), 0, 0},
		{"medium_array_found_end", generateSequence(0, 99), 99, 99},
		{"medium_array_found_middle", generateSequence(0, 99), 50, 50},
		{"medium_array_not_found", generateSequence(0, 99), 100, -1},

		// 大规模数据
		{"large_array_found_start", generateSequence(0, 9999), 0, 0},
		{"large_array_found_end", generateSequence(0, 9999), 9999, 9999},
		{"large_array_found_quarter", generateSequence(0, 9999), 2500, 2500},
		{"large_array_found_three_quarter", generateSequence(0, 9999), 7500, 7500},
		{"large_array_not_found", generateSequence(0, 9999), 10000, -1},

		// 特殊模式
		{"powers_of_two_found", []int{1, 2, 4, 8, 16, 32, 64, 128}, 16, 4},
		{"powers_of_two_not_found", []int{1, 2, 4, 8, 16, 32, 64, 128}, 10, -1},

		// 包含零
		{"contains_zero_found", []int{-10, -5, 0, 5, 10}, 0, 2},
		{"contains_zero_search_negative", []int{-10, -5, 0, 5, 10}, -5, 1},
		{"contains_zero_search_positive", []int{-10, -5, 0, 5, 10}, 10, 4},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := search(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %d) = %d, want %d", tst.nums, tst.target, got, tst.want)
			}

			got = search2(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %d) = %d, want %d", tst.nums, tst.target, got, tst.want)
			}

			got = search3(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %d) = %d, want %d", tst.nums, tst.target, got, tst.want)
			}

			got = search4(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %d) = %d, want %d", tst.nums, tst.target, got, tst.want)
			}

			got = search5(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %d) = %d, want %d", tst.nums, tst.target, got, tst.want)
			}
		})
	}
}

// generateSequence 生成从 start 到 end 的连续整数切片
func generateSequence(start, end int) []int {
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = start + i
	}
	return result
}
