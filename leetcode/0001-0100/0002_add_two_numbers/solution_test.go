package leetcode0002

import (
	"reflect"
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestAddTwoNumbersExamples(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "example_1",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8}, // 342 + 465 = 807
		},
		{
			name: "example_2",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			name: "example_3",
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1}, // 9999999 + 9999 = 10009998
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			l1 := utils.SliceToList(tst.l1)
			l2 := utils.SliceToList(tst.l2)
			got := addTwoNumbers(l1, l2)
			gotSlice := utils.ListToSlice(got)
			if !reflect.DeepEqual(gotSlice, tst.want) {
				t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", tst.l1, tst.l2, gotSlice, tst.want)
			}
		})
	}
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		// Single digit cases
		{"single_digit_no_carry", []int{1}, []int{2}, []int{3}},
		{"single_digit_with_carry", []int{5}, []int{7}, []int{2, 1}},
		{"single_digit_zero", []int{0}, []int{0}, []int{0}},
		{"single_digit_nine", []int{9}, []int{9}, []int{8, 1}},

		// Different lengths - l1 longer
		{"l1_longer_no_carry", []int{1, 2, 3}, []int{4}, []int{5, 2, 3}},
		{"l1_longer_with_carry", []int{9, 9, 9}, []int{1}, []int{0, 0, 0, 1}},
		{"l1_longer_partial_carry", []int{1, 9, 9}, []int{9}, []int{0, 0, 0, 1}},

		// Different lengths - l2 longer
		{"l2_longer_no_carry", []int{4}, []int{1, 2, 3}, []int{5, 2, 3}},
		{"l2_longer_with_carry", []int{1}, []int{9, 9, 9}, []int{0, 0, 0, 1}},
		{"l2_longer_partial_carry", []int{9}, []int{1, 9, 9}, []int{0, 0, 0, 1}},

		// Equal lengths
		{"equal_len_no_carry", []int{1, 2, 3}, []int{4, 5, 6}, []int{5, 7, 9}},
		{"equal_len_all_carry", []int{9, 9, 9}, []int{9, 9, 9}, []int{8, 9, 9, 1}},
		{"equal_len_some_carry", []int{1, 9, 1}, []int{9, 1, 9}, []int{0, 1, 1, 1}},

		// All zeros
		{"all_zeros_single", []int{0}, []int{0}, []int{0}},
		{"zero_plus_number", []int{0}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"number_plus_zero", []int{1, 2, 3}, []int{0}, []int{1, 2, 3}},

		// All nines (maximum carry propagation)
		{"all_nines_same_len", []int{9, 9}, []int{9, 9}, []int{8, 9, 1}},
		{"all_nines_diff_len", []int{9, 9, 9, 9}, []int{9}, []int{8, 0, 0, 0, 1}},
		{"nine_plus_one", []int{9}, []int{1}, []int{0, 1}},
		{"nines_plus_one", []int{9, 9, 9, 9, 9}, []int{1}, []int{0, 0, 0, 0, 0, 1}},

		// Carry propagation through multiple digits
		{"carry_chain", []int{1, 9, 9, 9}, []int{9}, []int{0, 0, 0, 0, 1}},
		{"long_carry_chain", []int{9, 9, 9, 9, 9, 9}, []int{1}, []int{0, 0, 0, 0, 0, 0, 1}},

		// Large numbers (near constraint limits)
		{"large_same_len", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}},

		// Mixed digit values
		{"mixed_values_1", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"mixed_values_2", []int{1, 8}, []int{0}, []int{1, 8}},
		{"mixed_values_3", []int{5}, []int{5}, []int{0, 1}},

		// Edge cases with carry at different positions
		{"carry_at_start", []int{9, 0, 0}, []int{1, 0, 0}, []int{0, 1, 0}},
		{"carry_at_middle", []int{0, 9, 0}, []int{0, 1, 0}, []int{0, 0, 1}},
		{"carry_at_end", []int{0, 0, 9}, []int{0, 0, 1}, []int{0, 0, 0, 1}},

		// Two digit numbers
		{"two_digit_no_carry", []int{1, 2}, []int{3, 4}, []int{4, 6}},
		{"two_digit_first_carry", []int{8, 2}, []int{5, 4}, []int{3, 7}},
		{"two_digit_second_carry", []int{1, 9}, []int{3, 9}, []int{4, 8, 1}},
		{"two_digit_both_carry", []int{9, 9}, []int{1, 1}, []int{0, 1, 1}},

		// Asymmetric cases
		{"one_vs_many", []int{5}, []int{1, 2, 3, 4, 5}, []int{6, 2, 3, 4, 5}},
		{"many_vs_one", []int{1, 2, 3, 4, 5}, []int{5}, []int{6, 2, 3, 4, 5}},

		// Result longer than both inputs
		{"result_longer", []int{9, 9}, []int{1}, []int{0, 0, 1}},
		{"result_much_longer", []int{9, 9, 9}, []int{9, 9, 9}, []int{8, 9, 9, 1}},

		// Specific patterns
		{"alternating_pattern", []int{1, 9, 1, 9}, []int{9, 1, 9, 1}, []int{0, 1, 1, 1, 1}},
		{"ascending_pattern", []int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{6, 8, 0, 3, 1}},
		{"descending_pattern", []int{4, 3, 2, 1}, []int{8, 7, 6, 5}, []int{2, 1, 9, 6}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			l1 := utils.SliceToList(tst.l1)
			l2 := utils.SliceToList(tst.l2)
			got := addTwoNumbers(l1, l2)
			gotSlice := utils.ListToSlice(got)
			if !reflect.DeepEqual(gotSlice, tst.want) {
				t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", tst.l1, tst.l2, gotSlice, tst.want)
			}
		})
	}
}
