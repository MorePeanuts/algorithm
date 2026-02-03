package leetcode0125

import "testing"

func TestIsPalindromeExamples(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example_1",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "example_2",
			s:    "race a car",
			want: false,
		},
		{
			name: "example_3",
			s:    " ",
			want: true,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := isPalindrome(tst.s)
			if got != tst.want {
				t.Errorf("isPalindrome(%q) = %v, want %v", tst.s, got, tst.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		// Single character cases
		{"single_letter_a", "a", true},
		{"single_letter_Z", "Z", true},
		{"single_digit", "5", true},
		{"single_space", " ", true},
		{"single_punctuation", ".", true},

		// Empty after filtering
		{"only_spaces", "   ", true},
		{"only_punctuation", ",.!?;:", true},
		{"mixed_non_alphanumeric", "  ,.  !@#$ ", true},

		// Simple palindromes (letters only)
		{"simple_aa", "aa", true},
		{"simple_aba", "aba", true},
		{"simple_abba", "abba", true},
		{"simple_abcba", "abcba", true},
		{"simple_racecar", "racecar", true},
		{"simple_level", "level", true},
		{"simple_noon", "noon", true},
		{"simple_deed", "deed", true},

		// Simple non-palindromes (letters only)
		{"simple_ab", "ab", false},
		{"simple_abc", "abc", false},
		{"simple_hello", "hello", false},
		{"simple_world", "world", false},

		// Case insensitivity
		{"mixed_case_Aa", "Aa", true},
		{"mixed_case_ABA", "ABA", true},
		{"mixed_case_AbA", "AbA", true},
		{"mixed_case_RaceCar", "RaceCar", true},
		{"mixed_case_not_palindrome", "Ab", false},

		// With numbers
		{"number_121", "121", true},
		{"number_12321", "12321", true},
		{"number_123", "123", false},
		{"alphanumeric_a1a", "a1a", true},
		{"alphanumeric_1a1", "1a1", true},
		{"alphanumeric_a1b", "a1b", false},
		{"alphanumeric_a11a", "a11a", true},

		// With spaces and punctuation (should be ignored)
		{"with_spaces_a_a", "a a", true},
		{"with_spaces_ab_ba", "ab ba", true},
		{"with_comma_a_comma_a", "a,a", true},
		{"with_punctuation_ab_exclaim_ba", "ab!ba", true},
		{"phrase_was_it_a_car", "Was it a car or a cat I saw", true},
		{"phrase_never_odd_even", "Never odd or even", true},
		{"phrase_no_lemon", "No lemon, no melon", true},
		{"phrase_not_palindrome", "This is not a palindrome", false},

		// Edge cases with mixed characters
		{"colon_and_spaces", "a:b::b:a", true},
		{"dots_and_letters", "a.b.b.a", true},
		{"numbers_with_symbols", "1,2,3,2,1", true},
		{"alphanumeric_mixed", "A1B2B1A", true},
		{"alphanumeric_mixed_not", "A1B2C1A", false},

		// Two character cases
		{"two_same_letters", "aa", true},
		{"two_same_letters_mixed_case", "aA", true},
		{"two_different_letters", "ab", false},
		{"two_chars_with_space", "a a", true},
		{"two_chars_with_punctuation", "a!a", true},

		// Longer palindromes
		{"long_madam", "madam", true},
		{"long_refer", "refer", true},
		{"long_civic", "civic", true},
		{"long_kayak", "kayak", true},
		{"long_rotator", "rotator", true},
		{"long_repaper", "repaper", true},

		// Complex phrases
		{"phrase_plan", "A man, a plan, a canal: Panama", true},
		{"phrase_race_car", "race a car", false},
		{"phrase_step_on_pets", "Step on no pets", true},
		{"phrase_live_evil", "Live not on evil", true},
		{"phrase_was_it_rat", "Was it a rat I saw", true},

		// All same character
		{"all_a", "aaaa", true},
		{"all_a_with_spaces", "a a a a", true},
		{"all_1", "1111", true},

		// Boundary: long strings
		{"medium_length_palindrome", "abcdefghijklmnopqrstuvwxyzzyxwvutsrqponmlkjihgfedcba", true},
		{"medium_length_not_palindrome", "abcdefghijklmnopqrstuvwxyz", false},

		// Special ASCII characters
		{"with_tabs", "a\ta", true},
		{"with_newline", "a\na", true},
		{"with_special_chars", "a@#$%^&*()a", true},
		{"unicode_like_ascii", "a~`a", true},

		// Numbers only
		{"all_zeros", "000", true},
		{"sequential_digits", "0123456789", false},
		{"repeated_pattern", "123321", true},
		{"single_zero", "0", true},

		// Alternating patterns
		{"alternating_ab", "ababababa", true},
		{"alternating_ab_not", "ababab", false},
		{"alternating_12", "121212121", true},

		// Almost palindromes
		{"almost_abca", "abca", false},
		{"almost_abcda", "abcda", false},
		{"almost_racecat", "racecat", false},
		{"off_by_one_middle", "abcda", false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := isPalindrome(tst.s)
			if got != tst.want {
				t.Errorf("isPalindrome(%q) = %v, want %v", tst.s, got, tst.want)
			}
		})
	}
}
