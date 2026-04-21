// Package leetcode0211 solves LeetCode 211. Design Add and Search Words Data Structure
package leetcode0211

import "testing"

func TestWordDictionaryExamples(t *testing.T) {
	tests := []struct {
		name     string
		actions  []string
		words    []string
		expected []bool
	}{
		{
			"example_1",
			[]string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
			[]string{"", "bad", "dad", "mad", "pad", "bad", ".ad", "b.."},
			[]bool{false, false, false, false, false, true, true, true},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			var dict WordDictionary
			for i, action := range tst.actions {
				switch action {
				case "WordDictionary":
					dict = Constructor()
				case "addWord":
					dict.AddWord(tst.words[i])
				case "search":
					got := dict.Search(tst.words[i])
					if got != tst.expected[i] {
						t.Errorf("search(%q) = %v, want %v", tst.words[i], got, tst.expected[i])
					}
				}
			}
		})
	}
}

func TestWordDictionary(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *WordDictionary
		search   string
		expected bool
	}{
		{
			"empty_dict_search_single_char",
			func() *WordDictionary {
				d := Constructor()
				return &d
			},
			"a",
			false,
		},
		{
			"empty_dict_search_dot",
			func() *WordDictionary {
				d := Constructor()
				return &d
			},
			".",
			false,
		},
		{
			"add_single_word_search_exact",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("a")
				return &d
			},
			"a",
			true,
		},
		{
			"add_single_word_search_dot",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("a")
				return &d
			},
			".",
			true,
		},
		{
			"add_single_word_search_wrong_char",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("a")
				return &d
			},
			"b",
			false,
		},
		{
			"add_two_words_search_exact",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("apple")
				d.AddWord("app")
				return &d
			},
			"apple",
			true,
		},
		{
			"add_two_words_search_prefix",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("apple")
				d.AddWord("app")
				return &d
			},
			"app",
			true,
		},
		{
			"add_two_words_search_prefix_dot",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("apple")
				d.AddWord("app")
				return &d
			},
			"ap.",
			true,
		},
		{
			"search_with_dot_in_middle",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("hello")
				d.AddWord("world")
				return &d
			},
			"h.llo",
			true,
		},
		{
			"search_with_dot_at_start",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("hello")
				return &d
			},
			".ello",
			true,
		},
		{
			"search_with_dot_at_end",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("hello")
				return &d
			},
			"hell.",
			true,
		},
		{
			"search_with_two_dots",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("abcde")
				return &d
			},
			"a.c.e",
			true,
		},
		{
			"search_with_two_consecutive_dots",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("abcde")
				return &d
			},
			"ab..e",
			true,
		},
		{
			"search_with_two_dots_no_match",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("abcde")
				return &d
			},
			"a..x",
			false,
		},
		{
			"max_length_word",
			func() *WordDictionary {
				d := Constructor()
				longWord := "abcdefghijklmnopqrstuvwxy" // 25 chars
				d.AddWord(longWord)
				return &d
			},
			"abcdefghijklmnopqrstuvwxy",
			true,
		},
		{
			"max_length_with_dots",
			func() *WordDictionary {
				d := Constructor()
				longWord := "abcdefghijklmnopqrstuvwxy"
				d.AddWord(longWord)
				return &d
			},
			"abcdefghijklmnopqrstuvwx.",
			true,
		},
		{
			"all_same_chars",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("aaaaa")
				return &d
			},
			"a.a.a",
			true,
		},
		{
			"search_longer_than_added",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("test")
				return &d
			},
			"testx",
			false,
		},
		{
			"search_shorter_than_added",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("testing")
				return &d
			},
			"test",
			false,
		},
		{
			"multiple_words_same_prefix",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("cat")
				d.AddWord("cats")
				d.AddWord("catalog")
				return &d
			},
			"cat.",
			true,
		},
		{
			"dot_matches_any_letter_but_not_empty",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("a")
				return &d
			},
			"..",
			false,
		},
		{
			"search_only_dots_exact_length",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("hello")
				d.AddWord("world")
				return &d
			},
			".....",
			true,
		},
		{
			"search_only_dots_wrong_length",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("hello")
				return &d
			},
			"....",
			false,
		},
		{
			"add_same_word_multiple_times",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("test")
				d.AddWord("test")
				return &d
			},
			"test",
			true,
		},
		{
			"search_with_dot_between_two_matches",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("abc")
				d.AddWord("axc")
				return &d
			},
			"a.c",
			true,
		},
		{
			"mix_of_matches_and_non_matches",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("apple")
				d.AddWord("apply")
				d.AddWord("orange")
				return &d
			},
			"appl.",
			true,
		},
		{
			"all_26_letters",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("abcdefghijklmnopqrstuvwxyz")
				return &d
			},
			"abcdefghijklmnopqrstuvwxyz",
			true,
		},
		{
			"two_dots_at_beginning",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("xyzabc")
				return &d
			},
			"..zabc",
			true,
		},
		{
			"two_dots_at_end",
			func() *WordDictionary {
				d := Constructor()
				d.AddWord("abcxyz")
				return &d
			},
			"abcx..",
			true,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			dict := tst.setup()
			got := dict.Search(tst.search)
			if got != tst.expected {
				t.Errorf("search(%q) = %v, want %v", tst.search, got, tst.expected)
			}
		})
	}
}
