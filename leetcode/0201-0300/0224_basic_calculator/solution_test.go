// Package leetcode0224 solves LeetCode 224. Basic Calculator
package leetcode0224

import "testing"

func TestCalculateExamples(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			"example_1",
			"1 + 1",
			2,
		},
		{
			"example_2",
			" 2-1 + 2 ",
			3,
		},
		{
			"example_3",
			"(1+(4+5+2)-3)+(6+8)",
			23,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := calculate(tst.s)
			if got != tst.want {
				t.Errorf("calculate(%q) = %v, want %v", tst.s, got, tst.want)
			}
			got = calculate2(tst.s)
			if got != tst.want {
				t.Errorf("calculate(%q) = %v, want %v", tst.s, got, tst.want)
			}
			got = calculate3(tst.s)
			if got != tst.want {
				t.Errorf("calculate(%q) = %v, want %v", tst.s, got, tst.want)
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
		// Single number cases
		{"single_zero", "0", 0},
		{"single_digit", "5", 5},
		{"single_negative", "-1", -1},
		{"single_negative_zero", "-0", 0},
		{"multiple_digits", "123", 123},
		{"negative_multiple_digits", "-456", -456},
		{"large_number", "2147483647", 2147483647},
		{"negative_large_number", "-2147483648", -2147483648},

		// Addition only
		{"simple_addition", "1+2", 3},
		{"addition_no_spaces", "10+20", 30},
		{"multiple_additions", "1+2+3+4+5", 15},
		{"addition_with_spaces", " 100 + 200 + 300 ", 600},

		// Subtraction only
		{"simple_subtraction", "5-3", 2},
		{"subtraction_no_spaces", "100-50", 50},
		{"multiple_subtractions", "10-1-2-3", 4},
		{"subtraction_with_spaces", " 500 - 200 - 100 ", 200},
		{"subtraction_negative_result", "3-5", -2},

		// Mixed addition and subtraction
		{"mixed_add_sub", "1+2-3", 0},
		{"mixed_sub_add", "5-3+4", 6},
		{"mixed_alternating", "1-2+3-4+5", 3},
		{"mixed_with_spaces", " 10 + 5 - 8 + 3 ", 10},

		// Parentheses cases
		{"simple_parentheses", "(1)", 1},
		{"negative_parentheses", "(-1)", -1},
		{"parentheses_addition", "(1+2)", 3},
		{"parentheses_subtraction", "(5-2)", 3},
		{"parentheses_before_add", "(1+2)+3", 6},
		{"parentheses_after_add", "1+(2+3)", 6},
		{"parentheses_before_sub", "(5-2)-1", 2},
		{"parentheses_after_sub", "5-(2-1)", 4},
		{"nested_parentheses", "((1))", 1},
		{"double_nested", "((1+2)-3)", 0},
		{"multiple_parentheses", "(1+2)+(3+4)", 10},
		{"multiple_parentheses_mixed", "(10-5)+(3-1)", 7},
		{"complex_nested", "(1+(2+(3+4)))", 10},
		{"parentheses_with_spaces", " ( 1 + 2 ) ", 3},

		// Unary minus cases
		{"unary_minus_start", "-1+2", 1},
		{"unary_minus_before_parentheses", "-(1+2)", -3},
		{"unary_minus_inside_parentheses", "(-1+2)", 1},
		{"unary_minus_with_add", "-10+5", -5},
		{"unary_minus_with_sub", "-10-5", -15},
		{"unary_minus_after_parentheses", "(1)-(-2)", 3},

		// Multiple spaces
		{"multiple_spaces_between", "1   +   2", 3},
		{"spaces_around_parentheses", "  (  1  +  2  )  ", 3},

		// Complex combinations
		{"complex_1", "1-(2-(3-(4-5)))", 3},
		{"complex_2", "-(1+(2-(3+4)))", 4},
		{"complex_3", "((-1-2)-(3+4))-(-5)", -5},
		{"complex_4", "100 - (200 + (300 - 400))", 0},
		{"complex_5", "-(-(-(-1+2)-3)-4)-5", -5},

		// Edge cases from constraints
		{"max_length_single_number", "0", 0},
		{"all_operators_and_parentheses", "(1-1)+(1-1)", 0},
		{"long_addition_chain", "1+1+1+1+1+1+1+1+1+1", 10},
		{"alternating_signs", "1-1+1-1+1-1+1-1+1-1", 0},

		// Zero handling
		{"zero_add", "0+0", 0},
		{"zero_sub", "0-0", 0},
		{"zero_mixed", "0+5-0", 5},
		{"zero_in_parentheses", "(0+5)-(0-3)", 8},
		{"zero_with_unary", "-0+5", 5},

		// Leetcode
		{"Left parenthesis plus consecutive spaces", "1-(     -2)", 3},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := calculate(tst.s)
			if got != tst.want {
				t.Errorf("calculate(%q) = %v, want %v", tst.s, got, tst.want)
			}
			got = calculate2(tst.s)
			if got != tst.want {
				t.Errorf("calculate(%q) = %v, want %v", tst.s, got, tst.want)
			}
			got = calculate3(tst.s)
			if got != tst.want {
				t.Errorf("calculate(%q) = %v, want %v", tst.s, got, tst.want)
			}
		})
	}
}
