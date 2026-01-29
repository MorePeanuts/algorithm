package leetcode0242

import (
	"strings"
	"testing"
)

func TestIsAnagramExamples(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			"example_1",
			"anagram",
			"nagaram",
			true,
		},
		{
			"example_2",
			"rat",
			"car",
			false,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := isAnagram(tst.s, tst.t)
			if got != tst.want {
				t.Errorf("isAnagram(%v, %v) = %v, want %v", tst.s, tst.t, got, tst.want)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		// 单字符测试
		{"single_char_same", "a", "a", true},
		{"single_char_diff", "a", "b", false},

		// 长度不等
		{"diff_length_s_longer", "abc", "ab", false},
		{"diff_length_t_longer", "ab", "abc", false},
		{"empty_vs_single", "", "a", false},
		{"single_vs_empty", "a", "", false},

		// 相同字符串
		{"identical_short", "abc", "abc", true},
		{"identical_medium", "algorithm", "algorithm", true},

		// 简单异位词
		{"simple_anagram_1", "ab", "ba", true},
		{"simple_anagram_2", "abc", "cba", true},
		{"simple_anagram_3", "abc", "bca", true},
		{"simple_anagram_4", "listen", "silent", true},
		{"simple_anagram_5", "evil", "vile", true},

		// 非异位词（相同长度但字符不同）
		{"not_anagram_1", "abc", "abd", false},
		{"not_anagram_2", "hello", "world", false},
		{"not_anagram_3", "abc", "xyz", false},

		// 重复字符
		{"repeated_chars_same", "aabb", "abab", true},
		{"repeated_chars_same_2", "aabb", "bbaa", true},
		{"repeated_chars_diff", "aabb", "aaab", false},
		{"all_same_char", "aaaa", "aaaa", true},
		{"all_same_char_diff_count", "aaa", "aaaa", false},

		// 字符频率不匹配
		{"freq_mismatch_1", "aab", "aba", true},
		{"freq_mismatch_2", "aab", "bbb", false},
		{"freq_mismatch_3", "aabc", "abcc", false},
		{"freq_mismatch_4", "abcc", "aabc", false},

		// 边界字符（a 和 z）
		{"boundary_chars_a", "aaa", "aaa", true},
		{"boundary_chars_z", "zzz", "zzz", true},
		{"boundary_chars_az", "az", "za", true},
		{"all_alphabet", "abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", true},

		// 包含所有字母的变体
		{"alphabet_repeated", "aabbccdd", "dcbadcba", true},
		{"alphabet_repeated_diff", "aabbccdd", "aabbccde", false},

		// 较长字符串
		{"long_same", strings.Repeat("ab", 100), strings.Repeat("ba", 100), true},
		{"long_diff", strings.Repeat("ab", 100), strings.Repeat("ac", 100), false},
		{"long_anagram", strings.Repeat("abc", 50), strings.Repeat("cab", 50), true},

		// 特殊模式
		{"palindrome_vs_self", "abba", "abba", true},
		{"palindrome_vs_anagram", "abba", "baab", true},
		{"alternating", "ababab", "bababa", true},
		{"alternating_diff", "ababab", "babbab", false},

		// 单字符重复
		{"single_repeated_10", strings.Repeat("x", 10), strings.Repeat("x", 10), true},
		{"single_repeated_diff_len", strings.Repeat("x", 10), strings.Repeat("x", 9), false},

		// 两字符组合
		{"two_chars_balanced", "aabb", "baba", true},
		{"two_chars_unbalanced", "aaab", "abbb", false},

		// 接近但不是异位词
		{"close_but_not_1", "ab", "ac", false},
		{"close_but_not_2", "abc", "adc", false},
		{"close_but_not_3", "abcd", "abce", false},

		// 实际单词异位词
		{"word_anagram_1", "cinema", "iceman", true},
		{"word_anagram_2", "binary", "brainy", true},
		{"word_anagram_3", "adobe", "abode", true},
		{"word_not_anagram", "apple", "papel", true},
		{"word_not_anagram_2", "apple", "apply", false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := isAnagram(tst.s, tst.t)
			if got != tst.want {
				t.Errorf("isAnagram(%q, %q) = %v, want %v", tst.s, tst.t, got, tst.want)
			}
		})
	}
}
