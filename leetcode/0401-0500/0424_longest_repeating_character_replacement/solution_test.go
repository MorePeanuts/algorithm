package leetcode0424

import "testing"

func TestCharacterReplacementExamples(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"example_1", "ABAB", 2, 4},
		{"example_2", "AABABBA", 1, 4},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := characterReplacement(tst.s, tst.k)
			if got != tst.want {
				t.Errorf("characterReplacement(%q, %d) = %d, want %d", tst.s, tst.k, got, tst.want)
			}
		})
	}
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		// Single character
		{"single_char", "A", 0, 1},
		{"single_char_with_k", "A", 1, 1},

		// All same characters
		{"all_same_2", "AA", 0, 2},
		{"all_same_3", "AAA", 0, 3},
		{"all_same_5", "AAAAA", 0, 5},
		{"all_same_with_k", "AAAA", 2, 4},

		// Two different characters
		{"two_chars_alternating", "ABAB", 0, 1},
		{"two_chars_alternating_k1", "ABAB", 1, 3},
		{"two_chars_grouped", "AABB", 0, 2},
		{"two_chars_grouped_k1", "AABB", 1, 3},
		{"two_chars_grouped_k2", "AABB", 2, 4},

		// k equals length (can change all)
		{"k_equals_length", "ABCD", 4, 4},
		{"k_equals_length_2", "ABCDEF", 6, 6},

		// k is zero
		{"k_zero_mixed", "ABCABC", 0, 1},
		{"k_zero_with_repeats", "AABBC", 0, 2},
		{"k_zero_longest_at_end", "ABCCC", 0, 3},

		// k is very large (larger than needed)
		{"k_larger_than_needed", "ABC", 10, 3},
		{"k_larger_than_length", "ABAB", 100, 4},

		// Various patterns
		{"pattern_aba", "ABA", 1, 3},
		{"pattern_abba", "ABBA", 1, 3},
		{"pattern_abba_k2", "ABBA", 2, 4},
		{"pattern_aabaa", "AABAA", 1, 5},
		{"pattern_abcba", "ABCBA", 2, 4},

		// Longer strings with patterns
		{"long_alternating", "ABABABABAB", 2, 5},
		{"long_alternating_k4", "ABABABABAB", 4, 9},
		{"long_alternating_k5", "ABABABABAB", 5, 10},

		// Multiple character types
		{"three_chars", "ABCABC", 2, 4},
		{"three_chars_k3", "ABCABC", 3, 5},
		{"four_chars", "ABCDABCD", 2, 3},
		{"four_chars_k4", "ABCDABCD", 4, 6},

		// Edge cases with position
		{"best_at_start", "AAABBB", 1, 4},
		{"best_at_end", "BBBAAAA", 1, 5},
		{"best_in_middle", "BAAAAB", 1, 5},

		// Complex patterns
		{"complex_1", "AABABBBAA", 2, 6},
		{"complex_2", "ABCDEFGH", 3, 4},
		{"complex_3", "AABAABAA", 1, 5},
		{"complex_4", "ABBBACCC", 2, 5},

		// All 26 letters
		{"all_letters", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 25, 26},
		{"all_letters_k0", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 0, 1},
		{"all_letters_k1", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, 2},

		// Repeated sections
		{"repeated_sections", "AAABBBCCC", 2, 5},
		{"repeated_sections_k3", "AAABBBCCC", 3, 6},
		{"repeated_sections_k4", "AAABBBCCC", 4, 7},

		// Long same character with interruptions
		{"long_with_interrupt", "AAAABAAA", 1, 8},
		{"long_with_interrupt_k0", "AAAABAAA", 0, 4},
		{"double_interrupt", "AAABAABA", 2, 8},

		// Boundary values
		{"two_chars", "AB", 0, 1},
		{"two_chars_k1", "AB", 1, 2},
		{"three_same", "BBB", 0, 3},

		// Window sliding scenarios
		{"slide_needed", "AABCDDD", 2, 5},
		{"slide_needed_2", "ABCDDDD", 2, 6},
		{"slide_needed_3", "DDDDABC", 2, 6},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := characterReplacement(tst.s, tst.k)
			if got != tst.want {
				t.Errorf("characterReplacement(%q, %d) = %d, want %d", tst.s, tst.k, got, tst.want)
			}
		})
	}
}
