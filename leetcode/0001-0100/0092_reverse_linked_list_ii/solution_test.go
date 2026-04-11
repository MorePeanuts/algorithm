package leetcode0092

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestReverseBetweenExamples(t *testing.T) {
	tests := []struct {
		name  string
		head  []int
		left  int
		right int
		want  []int
	}{
		{
			"example_1",
			[]int{1, 2, 3, 4, 5},
			2,
			4,
			[]int{1, 4, 3, 2, 5},
		},
		{
			"example_2",
			[]int{5},
			1,
			1,
			[]int{5},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToList(tst.head)
			got := reverseBetween(head, tst.left, tst.right)
			gotSlice := utils.ListToSlice(got)

			if len(gotSlice) != len(tst.want) {
				t.Errorf("reverseBetween(%v, %d, %d) = %v, want %v", tst.head, tst.left, tst.right, gotSlice, tst.want)
				return
			}

			for i := range gotSlice {
				if gotSlice[i] != tst.want[i] {
					t.Errorf("reverseBetween(%v, %d, %d) = %v, want %v", tst.head, tst.left, tst.right, gotSlice, tst.want)
					return
				}
			}
		})
	}
}

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		head  []int
		left  int
		right int
		want  []int
	}{
		{
			"single_element",
			[]int{100},
			1,
			1,
			[]int{100},
		},
		{
			"reverse_entire_list",
			[]int{1, 2, 3, 4, 5},
			1,
			5,
			[]int{5, 4, 3, 2, 1},
		},
		{
			"reverse_first_two",
			[]int{1, 2, 3, 4, 5},
			1,
			2,
			[]int{2, 1, 3, 4, 5},
		},
		{
			"reverse_last_two",
			[]int{1, 2, 3, 4, 5},
			4,
			5,
			[]int{1, 2, 3, 5, 4},
		},
		{
			"reverse_middle_two",
			[]int{1, 2, 3, 4, 5},
			2,
			3,
			[]int{1, 3, 2, 4, 5},
		},
		{
			"two_elements_reverse_both",
			[]int{1, 2},
			1,
			2,
			[]int{2, 1},
		},
		{
			"two_elements_left_equals_right",
			[]int{1, 2},
			1,
			1,
			[]int{1, 2},
		},
		{
			"three_elements_reverse_all",
			[]int{1, 2, 3},
			1,
			3,
			[]int{3, 2, 1},
		},
		{
			"negative_values",
			[]int{-5, -4, -3, -2, -1},
			2,
			4,
			[]int{-5, -2, -3, -4, -1},
		},
		{
			"mixed_values",
			[]int{-100, 0, 200, -300, 400},
			1,
			5,
			[]int{400, -300, 200, 0, -100},
		},
		{
			"max_length_list",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			3,
			8,
			[]int{1, 2, 8, 7, 6, 5, 4, 3, 9, 10},
		},
		{
			"reverse_first_element",
			[]int{1, 2, 3, 4, 5},
			1,
			1,
			[]int{1, 2, 3, 4, 5},
		},
		{
			"reverse_last_element",
			[]int{1, 2, 3, 4, 5},
			5,
			5,
			[]int{1, 2, 3, 4, 5},
		},
		{
			"adjacent_elements_beginning",
			[]int{10, 20, 30, 40, 50},
			1,
			3,
			[]int{30, 20, 10, 40, 50},
		},
		{
			"adjacent_elements_end",
			[]int{10, 20, 30, 40, 50},
			3,
			5,
			[]int{10, 20, 50, 40, 30},
		},
		{
			"all_same_values",
			[]int{7, 7, 7, 7, 7},
			2,
			4,
			[]int{7, 7, 7, 7, 7},
		},
		{
			"alternating_values",
			[]int{1, -1, 1, -1, 1, -1},
			2,
			5,
			[]int{1, 1, -1, 1, -1, -1},
		},
		{
			"reverse_three_in_middle",
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			5,
			[]int{1, 2, 5, 4, 3, 6, 7},
		},
		{
			"reverse_four_elements",
			[]int{1, 2, 3, 4, 5, 6},
			2,
			5,
			[]int{1, 5, 4, 3, 2, 6},
		},
		{
			"minimum_node_value",
			[]int{-500, -400, -300, -200, -100},
			1,
			3,
			[]int{-300, -400, -500, -200, -100},
		},
		{
			"maximum_node_value",
			[]int{100, 200, 300, 400, 500},
			3,
			5,
			[]int{100, 200, 500, 400, 300},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head := utils.SliceToList(tst.head)
			got := reverseBetween(head, tst.left, tst.right)
			gotSlice := utils.ListToSlice(got)

			if len(gotSlice) != len(tst.want) {
				t.Errorf("reverseBetween(%v, %d, %d) = %v, want %v", tst.head, tst.left, tst.right, gotSlice, tst.want)
				return
			}

			for i := range gotSlice {
				if gotSlice[i] != tst.want[i] {
					t.Errorf("reverseBetween(%v, %d, %d) = %v, want %v", tst.head, tst.left, tst.right, gotSlice, tst.want)
					return
				}
			}
		})
	}
}
