package leetcode0076

import "testing"

func TestMinWindowExamples(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			"example_1",
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"example_2",
			"a",
			"a",
			"a",
		},
		{
			"example_3",
			"a",
			"aa",
			"",
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := minWindow(tst.s, tst.t)
			if got != tst.want {
				t.Errorf("minWindow(%q, %q) = %q, want %q", tst.s, tst.t, got, tst.want)
			}
		})
	}
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		// Single character cases
		{"single_char_match", "a", "a", "a"},
		{"single_char_no_match", "a", "b", ""},
		{"single_char_in_longer_string", "abc", "b", "b"},
		{"single_char_at_start", "abc", "a", "a"},
		{"single_char_at_end", "abc", "c", "c"},

		// t longer than s
		{"t_longer_than_s", "ab", "abc", ""},
		{"t_much_longer", "a", "aaaa", ""},

		// Exact match
		{"exact_match", "abc", "abc", "abc"},
		{"exact_match_single", "x", "x", "x"},

		// Duplicate characters in t
		{"duplicate_in_t", "aa", "aa", "aa"},
		{"duplicate_in_t_insufficient", "ab", "aa", ""},
		{"triple_char_needed", "aaa", "aaa", "aaa"},
		{"duplicate_scattered", "abba", "aa", "abba"},
		{"duplicate_adjacent", "aab", "aa", "aa"},

		// Case sensitivity
		{"case_sensitive_match", "aA", "aA", "aA"},
		{"case_sensitive_no_match", "aa", "aA", ""},
		{"case_sensitive_upper", "ABC", "ABC", "ABC"},
		{"case_sensitive_mixed", "AaBbCc", "abc", "aBbCc"},

		// Window at different positions
		{"window_at_start", "abcdef", "ab", "ab"},
		{"window_at_end", "defabc", "abc", "abc"},
		{"window_in_middle", "xyzabcdef", "abc", "abc"},

		// Multiple valid windows - should return minimum
		{"multiple_windows_choose_min", "ADOBECODEBANC", "ABC", "BANC"},
		{"overlapping_windows", "aabaa", "aa", "aa"},

		// Complex cases
		{"all_same_char", "aaaa", "aa", "aa"},
		{"all_same_char_exact", "aaa", "aaa", "aaa"},
		{"interleaved_chars", "abababab", "aaa", "ababa"},
		{"sparse_match", "axbycz", "abc", "axbyc"},

		// Edge cases with spaces (only letters per constraints, but good to test)
		{"consecutive_target", "abcabc", "abc", "abc"},
		{"reverse_order_target", "cba", "abc", "cba"},

		// Longer strings
		{"longer_string_start_window", "ABCDEFGHIJKLMNOPabc", "abc", "abc"},
		{"longer_string_middle_window", "xxxxxabcxxxxx", "abc", "abc"},
		{"longer_scattered", "aXXXbXXXc", "abc", "aXXXbXXXc"},

		// No match scenarios
		{"no_match_different_chars", "xyz", "abc", ""},
		{"no_match_missing_one", "ab", "abc", ""},
		{"no_match_case_mismatch", "ABC", "abc", ""},

		// Window contains extra characters
		{"window_with_extras", "aXYZbXYZc", "abc", "aXYZbXYZc"},
		{"tight_window", "bba", "ab", "ba"},

		// Repeated patterns
		{"repeated_pattern", "abcabcabc", "abc", "abc"},
		{"find_tighter_window", "bdab", "ab", "ab"},

		// Two character target
		{"two_char_adjacent", "ab", "ab", "ab"},
		{"two_char_separated", "aXb", "ab", "aXb"},
		{"two_char_reversed", "ba", "ab", "ba"},

		// Large count of same char needed
		{"five_a_needed", "aaaaa", "aaaaa", "aaaaa"},
		{"five_a_in_longer", "baaaaaXXX", "aaaaa", "aaaaa"},

		// Leetcode test
		{"leetcode test", "cabwefgewcwaefgcf", "cae", "cwae"},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := minWindow(tst.s, tst.t)
			if got != tst.want {
				t.Errorf("minWindow(%q, %q) = %q, want %q", tst.s, tst.t, got, tst.want)
			}
		})
	}
}
