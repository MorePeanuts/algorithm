package leetcode0020

import (
	"strings"
	"testing"
)

func TestIsValidExamples(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"example_1", "()", true},
		{"example_2", "()[]{}", true},
		{"example_3", "(]", false},
		{"example_4", "([])", true},
		{"example_5", "([)]", false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := isValid(tst.s)
			if got != tst.want {
				t.Errorf("isValid(%q) = %v, want %v", tst.s, got, tst.want)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		// Single character (always invalid since unpaired)
		{"single_open_paren", "(", false},
		{"single_close_paren", ")", false},
		{"single_open_bracket", "[", false},
		{"single_close_bracket", "]", false},
		{"single_open_brace", "{", false},
		{"single_close_brace", "}", false},

		// Simple valid pairs
		{"simple_paren", "()", true},
		{"simple_bracket", "[]", true},
		{"simple_brace", "{}", true},

		// Nested valid
		{"nested_paren_in_bracket", "[()]", true},
		{"nested_bracket_in_brace", "{[]}", true},
		{"nested_brace_in_paren", "({})", true},
		{"deeply_nested", "([{}])", true},
		{"triple_nested", "{[()]}", true},

		// Sequential valid
		{"sequential_same_type", "()()()", true},
		{"sequential_mixed", "()[]{}()[]", true},
		{"sequential_all_types", "(){}[]", true},

		// Mixed nested and sequential
		{"nested_and_sequential", "()([]){}[[]]", true},
		{"complex_valid", "{[()]}[]({})", true},

		// Invalid: wrong order
		{"wrong_order_paren_bracket", "(]", false},
		{"wrong_order_bracket_brace", "[}", false},
		{"wrong_order_brace_paren", "{)", false},
		{"interleaved_invalid", "([)]", false},
		{"interleaved_invalid_2", "{[}]", false},
		{"interleaved_invalid_3", "[(])", false},

		// Invalid: unmatched open
		{"unmatched_open_paren", "(", false},
		{"unmatched_open_bracket", "[", false},
		{"unmatched_open_brace", "{", false},
		{"unmatched_open_at_end", "()[", false},
		{"unmatched_open_nested", "(()", false},
		{"unmatched_multiple_open", "(((", false},

		// Invalid: unmatched close
		{"unmatched_close_paren", ")", false},
		{"unmatched_close_bracket", "]", false},
		{"unmatched_close_brace", "}", false},
		{"unmatched_close_at_start", ")[]", false},
		{"unmatched_close_nested", "())", false},
		{"unmatched_multiple_close", ")))", false},

		// Invalid: close before open
		{"close_before_open_paren", ")(", false},
		{"close_before_open_bracket", "][", false},
		{"close_before_open_brace", "}{", false},
		{"close_before_open_mixed", "}{[]", false},

		// Edge cases with length
		{"two_chars_valid", "()", true},
		{"two_chars_invalid", "(}", false},
		{"odd_length", "(()", false},
		{"odd_length_2", "())", false},
		{"odd_length_3", "([)", false},

		// Longer sequences
		{"long_valid_sequential", strings.Repeat("()", 50), true},
		{"long_valid_nested", strings.Repeat("(", 50) + strings.Repeat(")", 50), true},
		{"long_invalid_unmatched", strings.Repeat("(", 51), false},
		{"long_mixed_valid", strings.Repeat("()[]{}", 20), true},

		// Complex patterns
		{"alternating_types", "()[](){}[]{}()", true},
		{"all_parens_nested", "(((())))((()))", true},
		{"all_brackets_nested", "[[[[]]]][[[]]]", true},
		{"all_braces_nested", "{{{{}}}}{{{}}}}", false},
		{"all_braces_nested_valid", "{{{{}}}}{{{}}}", true},

		// Tricky cases
		{"empty_in_middle", "()()()", true},
		{"valid_then_invalid", "())(", false},
		{"invalid_then_valid", ")(()", false},
		{"almost_valid", "([]){", false},
		{"almost_valid_2", "{[()]", false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := isValid(tst.s)
			if got != tst.want {
				t.Errorf("isValid(%q) = %v, want %v", tst.s, got, tst.want)
			}
		})
	}
}
