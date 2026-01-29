package leetcode0049

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestGroupAnagramsExamples(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			"example_1",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			"example_2",
			[]string{""},
			[][]string{{""}},
		},
		{
			"example_3",
			[]string{"a"},
			[][]string{{"a"}},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := groupAnagrams(tst.strs)
			if !utils.MatchTwo2dStringSlice(got, tst.want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.want)
			}

			got = groupAnagrams2(tst.strs)
			if !utils.MatchTwo2dStringSlice(got, tst.want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.want)
			}

			got = groupAnagrams3(tst.strs)
			if !utils.MatchTwo2dStringSlice(got, tst.want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.want)
			}
		})
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		// 单元素测试
		{"single_empty_string", []string{""}, [][]string{{""}}},
		{"single_char_a", []string{"a"}, [][]string{{"a"}}},
		{"single_char_z", []string{"z"}, [][]string{{"z"}}},
		{"single_long_string", []string{"abcdefghij"}, [][]string{{"abcdefghij"}}},

		// 全相同元素
		{"all_same_empty", []string{"", "", ""}, [][]string{{"", "", ""}}},
		{"all_same_single_char", []string{"a", "a", "a"}, [][]string{{"a", "a", "a"}}},
		{"all_same_word", []string{"abc", "abc", "abc"}, [][]string{{"abc", "abc", "abc"}}},

		// 全是异位词
		{"all_anagrams_two", []string{"ab", "ba"}, [][]string{{"ab", "ba"}}},
		{"all_anagrams_three", []string{"abc", "bca", "cab"}, [][]string{{"abc", "bca", "cab"}}},
		{"all_anagrams_six", []string{"abc", "acb", "bac", "bca", "cab", "cba"}, [][]string{{"abc", "acb", "bac", "bca", "cab", "cba"}}},

		// 无异位词（每个字符串独立成组）
		{"no_anagrams_two", []string{"a", "b"}, [][]string{{"a"}, {"b"}}},
		{"no_anagrams_three", []string{"a", "bb", "ccc"}, [][]string{{"a"}, {"bb"}, {"ccc"}}},
		{"no_anagrams_diff_lengths", []string{"a", "ab", "abc"}, [][]string{{"a"}, {"ab"}, {"abc"}}},

		// 混合情况
		{"mixed_with_empty", []string{"", "a", ""}, [][]string{{"", ""}, {"a"}}},
		{"mixed_anagrams_and_singles", []string{"ab", "ba", "c"}, [][]string{{"ab", "ba"}, {"c"}}},
		{"mixed_multiple_groups", []string{"ab", "cd", "ba", "dc"}, [][]string{{"ab", "ba"}, {"cd", "dc"}}},

		// 重复字符
		{"repeated_chars_same", []string{"aa", "aa"}, [][]string{{"aa", "aa"}}},
		{"repeated_chars_anagram", []string{"aab", "aba", "baa"}, [][]string{{"aab", "aba", "baa"}}},
		{"repeated_chars_diff", []string{"aab", "abb"}, [][]string{{"aab"}, {"abb"}}},

		// 长度差异
		{"diff_length_no_anagram", []string{"a", "aa", "aaa"}, [][]string{{"a"}, {"aa"}, {"aaa"}}},
		{"same_chars_diff_counts", []string{"ab", "aab", "aaab"}, [][]string{{"ab"}, {"aab"}, {"aaab"}}},

		// 边界字符
		{"all_a", []string{"a", "a", "a", "a", "a"}, [][]string{{"a", "a", "a", "a", "a"}}},
		{"all_z", []string{"z", "z", "z"}, [][]string{{"z", "z", "z"}}},
		{"a_and_z", []string{"az", "za"}, [][]string{{"az", "za"}}},
		{"full_alphabet", []string{"abcdefghijklmnopqrstuvwxyz"}, [][]string{{"abcdefghijklmnopqrstuvwxyz"}}},
		{"full_alphabet_anagram", []string{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"}, [][]string{{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"}}},

		// 空字符串混合
		{"empty_with_single_char", []string{"", "a", "b", ""}, [][]string{{"", ""}, {"a"}, {"b"}}},
		{"empty_with_anagrams", []string{"", "ab", "ba", ""}, [][]string{{"", ""}, {"ab", "ba"}}},

		// 多组混合
		{"three_groups", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{"four_groups", []string{"a", "b", "ab", "ba", "c", "abc"}, [][]string{{"a"}, {"b"}, {"ab", "ba"}, {"c"}, {"abc"}}},
		{"five_groups", []string{"a", "b", "c", "d", "e"}, [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}}},

		// 较长字符串
		{"long_string_single", []string{"aaaaabbbbbccccc"}, [][]string{{"aaaaabbbbbccccc"}}},
		{"long_string_anagram", []string{"aabbcc", "abcabc", "ccbbaa"}, [][]string{{"aabbcc", "abcabc", "ccbbaa"}}},
		{"long_repeated_char", []string{"aaaaaaaaaa", "aaaaaaaaaa"}, [][]string{{"aaaaaaaaaa", "aaaaaaaaaa"}}},

		// 特殊模式
		{"palindrome_anagrams", []string{"aba", "aab", "baa"}, [][]string{{"aba", "aab", "baa"}}},
		{"consecutive_chars", []string{"abc", "def", "ghi"}, [][]string{{"abc"}, {"def"}, {"ghi"}}},
		{"same_freq_diff_chars", []string{"aabb", "ccdd"}, [][]string{{"aabb"}, {"ccdd"}}},

		// 大量重复
		{"many_same_empty", []string{"", "", "", "", ""}, [][]string{{"", "", "", "", ""}}},
		{"many_same_char", []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}, [][]string{{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}}},

		// 复杂混合
		{"complex_mixed_1", []string{"", "a", "b", "ab", "ba", "abc", "bca", "cab"}, [][]string{{""}, {"a"}, {"b"}, {"ab", "ba"}, {"abc", "bca", "cab"}}},
		{"complex_mixed_2", []string{"listen", "silent", "enlist", "hello", "world"}, [][]string{{"listen", "silent", "enlist"}, {"hello"}, {"world"}}},
		{"complex_mixed_3", []string{"dusty", "study", "night", "thing"}, [][]string{{"dusty", "study"}, {"night", "thing"}}},

		// 两字符组合
		{"two_char_all_anagrams", []string{"ab", "ba", "ab", "ba"}, [][]string{{"ab", "ba", "ab", "ba"}}},
		{"two_char_mixed", []string{"ab", "cd", "ba", "dc", "ef"}, [][]string{{"ab", "ba"}, {"cd", "dc"}, {"ef"}}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := groupAnagrams(tst.strs)
			if !utils.MatchTwo2dStringSlice(got, tst.want) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.want)
			}

			got = groupAnagrams2(tst.strs)
			if !utils.MatchTwo2dStringSlice(got, tst.want) {
				t.Errorf("groupAnagrams2(%v) = %v, want %v", tst.strs, got, tst.want)
			}

			got = groupAnagrams3(tst.strs)
			if !utils.MatchTwo2dStringSlice(got, tst.want) {
				t.Errorf("groupAnagrams3(%v) = %v, want %v", tst.strs, got, tst.want)
			}
		})
	}
}
