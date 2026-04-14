package leetcode0033

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
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		},
		{
			"example_2",
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
			-1,
		},
		{
			"example_3",
			[]int{1},
			0,
			-1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := search(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %v) = %v, want %v", tst.nums, tst.target, got, tst.want)
			}

			got2 := search2(tst.nums, tst.target)
			if got2 != tst.want {
				t.Errorf("search2(%v, %v) = %v, want %v", tst.nums, tst.target, got2, tst.want)
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
		{"single_element_found", []int{5}, 5, 0},
		{"single_element_not_found", []int{5}, 3, -1},
		{"not_rotated_found_beginning", []int{1, 2, 3, 4, 5}, 1, 0},
		{"not_rotated_found_middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"not_rotated_found_end", []int{1, 2, 3, 4, 5}, 5, 4},
		{"not_rotated_not_found_smaller", []int{1, 2, 3, 4, 5}, 0, -1},
		{"not_rotated_not_found_larger", []int{1, 2, 3, 4, 5}, 6, -1},
		{"rotated_at_1_found", []int{2, 3, 4, 5, 1}, 1, 4},
		{"rotated_at_1_not_found", []int{2, 3, 4, 5, 1}, 6, -1},
		{"rotated_at_middle_left_part_found", []int{4, 5, 6, 7, 0, 1, 2}, 5, 1},
		{"rotated_at_middle_right_part_found", []int{4, 5, 6, 7, 0, 1, 2}, 1, 5},
		{"rotated_at_middle_first_element", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{"rotated_at_middle_last_element", []int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
		{"rotated_at_middle_min_element", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"rotated_at_middle_max_element", []int{4, 5, 6, 7, 0, 1, 2}, 7, 3},
		{"rotated_at_len_minus_1", []int{5, 1, 2, 3, 4}, 1, 1},
		{"rotated_at_len_minus_1_found_left", []int{5, 1, 2, 3, 4}, 5, 0},
		{"rotated_at_len_minus_1_found_right", []int{5, 1, 2, 3, 4}, 3, 3},
		{"two_elements_rotated_found_first", []int{2, 1}, 2, 0},
		{"two_elements_rotated_found_second", []int{2, 1}, 1, 1},
		{"two_elements_not_rotated_found_first", []int{1, 2}, 1, 0},
		{"two_elements_not_rotated_found_second", []int{1, 2}, 2, 1},
		{"two_elements_rotated_not_found", []int{2, 1}, 3, -1},
		{"negative_numbers_rotated", []int{-3, -2, -1, -5, -4}, -5, 3},
		{"negative_numbers_found", []int{-3, -2, -1, -5, -4}, -2, 1},
		{"mixed_numbers_rotated", []int{100, 200, 300, -100, 0, 50}, -100, 3},
		{"mixed_numbers_found_zero", []int{100, 200, 300, -100, 0, 50}, 0, 4},
		{"max_length_edge_not_rotated", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, 9},
		{"max_length_edge_rotated", []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}, 1, 6},
		{"target_smaller_than_all", []int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 0, -1},
		{"target_larger_than_all", []int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 10, -1},
		{"target_in_gap_between_parts", []int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 4, 8},
		{"target_in_gap_not_found", []int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 4, 8},
		{"rotated_at_2", []int{3, 4, 5, 1, 2}, 1, 3},
		{"rotated_at_3", []int{4, 5, 1, 2, 3}, 1, 2},
		{"all_negative_rotated", []int{-2, -1, -5, -4, -3}, -3, 4},
		{"all_negative_not_rotated", []int{-5, -4, -3, -2, -1}, -3, 2},
		{"zero_in_right_part", []int{5, 6, 7, 0, 1, 2, 3, 4}, 0, 3},
		{"zero_in_left_part", []int{0, 1, 2, 3, 4, 5, -5, -4, -3, -2, -1}, 0, 0},
		{"target_at_rotation_point", []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}, 1, 5},
		{"target_before_rotation_point", []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}, 10, 4},
		{"min_val_neg_10000", []int{0, 1, 2, 3, -10000}, -10000, 4},
		{"max_val_10000", []int{10000, -10000, -5000, 0, 5000}, 10000, 0},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := search(tst.nums, tst.target)
			if got != tst.want {
				t.Errorf("search(%v, %v) = %v, want %v", tst.nums, tst.target, got, tst.want)
			}

			got2 := search2(tst.nums, tst.target)
			if got2 != tst.want {
				t.Errorf("search2(%v, %v) = %v, want %v", tst.nums, tst.target, got2, tst.want)
			}
		})
	}
}
