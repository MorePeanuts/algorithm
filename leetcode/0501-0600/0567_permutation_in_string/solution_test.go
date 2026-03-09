package leetcode0567

import "testing"

func TestCheckInclusionExamples(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{"example_1", "ab", "eidbaooo", true},
		{"example_2", "ab", "eidboaoo", false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := checkInclusion(tst.s1, tst.s2)
			if got != tst.want {
				t.Errorf("checkInclusion(%q, %q) = %v, want %v", tst.s1, tst.s2, got, tst.want)
			}
		})
	}
}

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		// Basic cases
		{"single_char_match", "a", "a", true},
		{"single_char_in_string", "a", "bca", true},
		{"single_char_not_found", "a", "bcd", false},
		{"identical_strings", "abc", "abc", true},
		{"reverse_string", "abc", "cba", true},

		// s1 longer than s2
		{"s1_longer_than_s2", "abcd", "abc", false},
		{"s1_much_longer", "abcdefg", "abc", false},

		// Position tests
		{"permutation_at_start", "abc", "cabdef", true},
		{"permutation_at_end", "abc", "defbca", true},
		{"permutation_in_middle", "abc", "xxbacyy", true},
		{"exact_length_match", "abc", "bca", true},
		{"exact_length_no_match", "abc", "abd", false},

		// Repeated characters in s1
		{"s1_repeated_chars_match", "aab", "cbdaaob", false},
		{"s1_repeated_chars_found", "aab", "cbdaabo", true},
		{"s1_all_same_chars", "aaa", "baaac", true},
		{"s1_all_same_not_enough", "aaa", "baac", false},
		{"s1_double_chars", "abab", "cbabade", true},
		{"s1_double_chars_no_match", "abab", "cababde", true},

		// Repeated characters in s2
		{"s2_has_repeats_match", "ab", "aaaaab", true},
		{"s2_has_repeats_no_match", "ab", "aaaaaa", false},
		{"s2_all_same_s1_different", "ab", "bbbbbb", false},
		{"s2_all_same_s1_same", "bb", "bbbbbb", true},

		// Edge boundary cases
		{"two_char_s1_match", "ab", "ba", true},
		{"two_char_s1_no_match", "ab", "aa", false},
		{"s1_equals_s2_length_no_match", "abc", "abd", false},

		// Various permutations
		{"permutation_scattered_no", "abc", "axxbxxc", false},
		{"permutation_adjacent", "abc", "xxacbxx", true},
		{"multiple_possible_windows", "ab", "abab", true},
		{"overlapping_potential", "aba", "abaaba", true},

		// Longer strings
		{"longer_pattern", "abcde", "xxxxxedcbaxxxxx", true},
		{"longer_pattern_no_match", "abcde", "xxxxxabcdxxxxxx", false},
		{"longer_s2_no_match", "xyz", "abcdefghijklmnop", false},
		{"longer_s2_match_end", "xyz", "abcdefghijklmnozyx", true},

		// Tricky cases
		{"almost_permutation", "abc", "abdc", false},
		{"one_char_off", "abc", "abd", false},
		{"extra_char_in_window", "ab", "acb", false},
		{"window_has_extra", "ab", "aab", true},

		// All lowercase letters
		{"full_alphabet_match", "abcdefghij", "jihgfedcba", true},
		{"partial_alphabet_no_match", "abcdefghij", "jihgfedcbx", false},

		// Stress test with repeated patterns
		{"repeated_pattern_s2", "ab", "ababababab", true},
		{"no_permutation_long", "abc", "aaaaaaaaaa", false},

		// Single character edge cases
		{"single_z", "z", "abcdefghijklmnopqrstuvwxyz", true},
		{"single_a_not_in", "a", "bcdefghijklmnopqrstuvwxyz", false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := checkInclusion(tst.s1, tst.s2)
			if got != tst.want {
				t.Errorf("checkInclusion(%q, %q) = %v, want %v", tst.s1, tst.s2, got, tst.want)
			}
		})
	}
}
