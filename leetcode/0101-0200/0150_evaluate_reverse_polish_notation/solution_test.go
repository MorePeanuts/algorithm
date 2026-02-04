package leetcode0150

import "testing"

func TestEvalRPNExamples(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			"example_1",
			[]string{"2", "1", "+", "3", "*"},
			9,
		},
		{
			"example_2",
			[]string{"4", "13", "5", "/", "+"},
			6,
		},
		{
			"example_3",
			[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			22,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := evalRPN(tst.tokens)
			if got != tst.want {
				t.Errorf("evalRPN(%v) = %v, want %v", tst.tokens, got, tst.want)
			}
		})
	}
}

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		// Single number
		{"single_positive", []string{"42"}, 42},
		{"single_negative", []string{"-42"}, -42},
		{"single_zero", []string{"0"}, 0},

		// Simple operations
		{"simple_add", []string{"1", "2", "+"}, 3},
		{"simple_sub", []string{"5", "3", "-"}, 2},
		{"simple_mul", []string{"3", "4", "*"}, 12},
		{"simple_div", []string{"8", "2", "/"}, 4},

		// Operations with negative numbers
		{"add_negative", []string{"-1", "-2", "+"}, -3},
		{"sub_negative", []string{"-5", "-3", "-"}, -2},
		{"mul_negative", []string{"-3", "4", "*"}, -12},
		{"mul_two_negative", []string{"-3", "-4", "*"}, 12},
		{"div_negative", []string{"-8", "2", "/"}, -4},
		{"div_two_negative", []string{"-8", "-2", "/"}, 4},

		// Division truncation toward zero
		{"div_truncate_positive", []string{"7", "3", "/"}, 2},
		{"div_truncate_negative", []string{"-7", "3", "/"}, -2},
		{"div_truncate_neg_divisor", []string{"7", "-3", "/"}, -2},
		{"div_truncate_both_neg", []string{"-7", "-3", "/"}, 2},
		{"div_result_zero", []string{"1", "3", "/"}, 0},
		{"div_neg_result_zero", []string{"-1", "3", "/"}, 0},

		// Chained same operations
		{"chain_add", []string{"1", "2", "+", "3", "+", "4", "+"}, 10},
		{"chain_mul", []string{"2", "3", "*", "4", "*"}, 24},
		{"chain_sub", []string{"10", "3", "-", "2", "-"}, 5},

		// Mixed operations
		{"mixed_add_mul", []string{"2", "3", "+", "4", "*"}, 20},
		{"mixed_mul_add", []string{"2", "3", "*", "4", "+"}, 10},
		{"mixed_sub_div", []string{"10", "2", "-", "4", "/"}, 2},

		// Complex expressions
		{"complex_1", []string{"5", "1", "2", "+", "4", "*", "+", "3", "-"}, 14},
		{"complex_2", []string{"3", "4", "+", "2", "*", "7", "/"}, 2},
		{"complex_3", []string{"4", "2", "/", "3", "+", "5", "*"}, 25},

		// Edge values from constraints
		{"max_operand", []string{"200", "1", "+"}, 201},
		{"min_operand", []string{"-200", "1", "-"}, -201},
		{"max_times_max", []string{"200", "200", "*"}, 40000},
		{"min_times_max", []string{"-200", "200", "*"}, -40000},

		// Zero operations
		{"add_zero", []string{"5", "0", "+"}, 5},
		{"sub_zero", []string{"5", "0", "-"}, 5},
		{"mul_zero", []string{"5", "0", "*"}, 0},
		{"zero_div", []string{"0", "5", "/"}, 0},

		// Nested operations
		{"nested_1", []string{"2", "1", "+", "3", "*"}, 9},
		{"nested_2", []string{"4", "13", "5", "/", "+"}, 6},
		{"deeply_nested", []string{"1", "2", "3", "4", "+", "*", "+"}, 15},

		// Expression with many operators
		{"many_ops", []string{"1", "2", "+", "3", "+", "4", "+", "5", "+"}, 15},
		{"alternating_ops", []string{"10", "2", "*", "5", "/", "3", "+", "1", "-"}, 6},

		// Large intermediate results
		{"large_intermediate", []string{"100", "100", "*", "50", "/"}, 200},

		// Subtraction order matters
		{"sub_order_1", []string{"3", "5", "-"}, -2},
		{"sub_order_2", []string{"5", "3", "-"}, 2},

		// Division order matters
		{"div_order_1", []string{"2", "8", "/"}, 0},
		{"div_order_2", []string{"8", "2", "/"}, 4},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := evalRPN(tst.tokens)
			if got != tst.want {
				t.Errorf("evalRPN(%v) = %v, want %v", tst.tokens, got, tst.want)
			}
		})
	}
}
