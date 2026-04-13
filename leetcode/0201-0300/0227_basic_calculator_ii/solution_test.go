package leetcode0227

import "testing"

func TestCalculateExamples(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			"example_1",
			"3+2*2",
			7,
		},
		{
			"example_2",
			" 3/2 ",
			1,
		},
		{
			"example_3",
			" 3+5 / 2 ",
			5,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := calculate(tst.s)
			if got != tst.want {
				t.Errorf("calculate(%q) = %v, want %v", tst.s, got, tst.want)
			}

			got2 := calculate2(tst.s)
			if got2 != tst.want {
				t.Errorf("calculate2(%q) = %v, want %v", tst.s, got2, tst.want)
			}

			got3 := calculate3(tst.s)
			if got3 != tst.want {
				t.Errorf("calculate3(%q) = %v, want %v", tst.s, got3, tst.want)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"single_number", "0", 0},
		{"single_large_number", "2147483647", 2147483647},
		{"only_addition", "1+2+3", 6},
		{"only_subtraction", "10-5-3", 2},
		{"only_multiplication", "2*3*4", 24},
		{"only_division", "100/2/5", 10},
		{"mixed_add_sub", "10+5-3+2", 14},
		{"mixed_mul_div", "100*2/5*3", 120},
		{"add_mul_order", "1+2*3", 7},
		{"mul_add_order", "2*3+1", 7},
		{"add_div_order", "10+8/2", 14},
		{"div_add_order", "8/2+10", 14},
		{"sub_mul_order", "10-2*3", 4},
		{"mul_sub_order", "2*3-10", -4},
		{"sub_div_order", "10-6/2", 7},
		{"div_sub_order", "6/2-10", -7},
		{"all_four_ops", "3+5*2-8/4", 11},
		{"multiple_mul_div", "2*3*4/6*5", 20},
		{"division_truncate_positive", "7/3", 2},
		{"division_truncate_negative_lhs", "10-7/3", 8},
		{"division_truncate_negative_result", "3-7/3", 1},
		{"spaces_everywhere", "   123   +   456   *   2   ", 1035},
		{"no_spaces", "10+20*3-40/2", 50},
		{"multi_digit_numbers", "123+456*789", 359907},
		{"chain_precedence", "1+2*3-4*5+6", -7},
		{"all_multiplication", "1*2*3*4*5*6*7*8*9", 362880},
		{"all_division_large", "1000000/2/5/2/5/2", 1000000 / 2 / 5 / 2 / 5 / 2},
		{"alternating_ops", "10-1*2+3*4-5", 15},
		{"subtraction_chain", "100-50-25-12-6-3-1", 3},
		{"zero_in_middle", "100+0*50", 100},
		{"zero_multiplication", "0*12345", 0},
		{"zero_division", "0/123", 0},
		{"large_intermediate_result", "100000*100000/100000", 100000},
		{"division_exact", "123456/789", 123456 / 789},
		{"division_inexact", "123457/789", 123457 / 789},
		{"mixed_operations_1", "1000+500*2-2000/4", 1500},
		{"mixed_operations_2", "99-9*9+9/9", 19},
		{"trailing_spaces", "123+456   ", 579},
		{"leading_spaces", "   123+456", 579},
		{"spaces_around_ops", "123   +   456   *   789", 359907},
		{"number_at_end_with_mul", "1+2*3*4*5", 121},
		{"number_at_end_with_div", "1000/2/5/2", 50},
		{"multiple_divisions", "100/2/5/2/5", 1},
		{"subtraction_after_division", "100-100/2", 50},
		{"addition_after_multiplication", "100+100*2", 300},
		{"complex_1", "1+2*3-4/5+6*7-8/9", 49},
		{"complex_2", "10*9-8*7+6*5-4*3+2*1", 54},
		{"many_ops", "1+1+1+1+1+1+1+1+1+1", 10},
		{"many_muls", "2*2*2*2*2*2*2*2*2*2", 1024},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := calculate(tst.s)
			if got != tst.want {
				t.Errorf("calculate(%q) = %v, want %v", tst.s, got, tst.want)
			}

			got2 := calculate2(tst.s)
			if got2 != tst.want {
				t.Errorf("calculate2(%q) = %v, want %v", tst.s, got2, tst.want)
			}

			got3 := calculate3(tst.s)
			if got3 != tst.want {
				t.Errorf("calculate3(%q) = %v, want %v", tst.s, got3, tst.want)
			}
		})
	}
}
