package leetcode0003

import "testing"

func TestLengthOfLongestSubstringExamples(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"example_1", "abcabcbb", 3},
		{"example_2", "bbbbb", 1},
		{"example_3", "pwwkew", 3},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tst.s)
			if got != tst.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tst.s, got, tst.want)
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		// Empty and single character
		{"empty_string", "", 0},
		{"single_char", "a", 1},
		{"single_space", " ", 1},
		{"single_digit", "5", 1},
		{"single_symbol", "@", 1},

		// All same characters
		{"all_same_a", "aaaa", 1},
		{"all_same_space", "    ", 1},
		{"all_same_digit", "1111", 1},

		// All unique characters
		{"all_unique_abc", "abc", 3},
		{"all_unique_abcdef", "abcdef", 6},
		{"all_unique_mixed", "a1b2c3", 6},
		{"all_unique_symbols", "!@#$%", 5},

		// Repeating at different positions
		{"repeat_at_start", "aabcd", 4},
		{"repeat_at_end", "abcdd", 4},
		{"repeat_in_middle", "abcbc", 3},
		{"repeat_wrap_around", "abcabc", 3},

		// Two characters
		{"two_same", "aa", 1},
		{"two_different", "ab", 2},

		// Longer patterns
		{"alternating_ab", "ababab", 2},
		{"alternating_abc", "abcabcabc", 3},
		{"increasing_then_repeat", "abcdefga", 7},

		// With spaces
		{"space_in_middle", "ab cd", 5},
		{"spaces_and_letters", "a b c", 3},
		{"multiple_spaces", "a  b", 2},

		// With digits
		{"digits_only", "0123456789", 10},
		{"digits_repeat", "01234567890", 10},
		{"mixed_letters_digits", "a1b2c3d4", 8},

		// With symbols
		{"symbols_mixed", "a!b@c#", 6},
		{"symbols_repeat", "!@#!@#", 3},

		// Complex cases
		{"dvdf", "dvdf", 3},
		{"anviaj", "anviaj", 5},
		{"tmmzuxt", "tmmzuxt", 5},
		{"ohvhjdml", "ohvhjdml", 6},

		// Substring at end
		{"longest_at_end", "aaaaabcdef", 6},
		{"longest_at_start", "abcdefaaaa", 6},

		// Edge cases with consecutive repeats
		{"double_repeat", "aabb", 2},
		{"triple_repeat", "aaabbb", 2},
		{"pattern_aabcb", "aabcb", 3},

		// Unicode-like ASCII edge cases
		{"printable_ascii_range", " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", 95},

		// Long unique sequence
		{"alphabet_lowercase", "abcdefghijklmnopqrstuvwxyz", 26},
		{"alphabet_uppercase", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 26},
		{"alphanumeric", "abcdefghijklmnopqrstuvwxyz0123456789", 36},

		// Boundary behaviors
		{"repeat_after_long_unique", "abcdefghijklmnopa", 16},
		{"repeat_every_other", "abacadae", 3},

		// Special patterns
		{"palindrome", "abcba", 3},
		{"reverse_sorted", "zyxwvutsrqponmlkjihgfedcba", 26},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tst.s)
			if got != tst.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tst.s, got, tst.want)
			}
		})
	}
}
