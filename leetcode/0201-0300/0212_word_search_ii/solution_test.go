// Package leetcode0212 tests LeetCode 212. Word Search II
package leetcode0212

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestFindWordsExamples(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		{
			"example_1",
			[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
			[]string{"oath", "pea", "eat", "rain"},
			[]string{"eat", "oath"},
		},
		{
			"example_2",
			[][]byte{{'a', 'b'}, {'c', 'd'}},
			[]string{"abcb"},
			[]string{},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := findWords(tst.board, tst.words)
			if !utils.MatchStringSlice(got, tst.want) {
				t.Errorf("findWords(%v, %v) = %v, want %v", tst.board, tst.words, got, tst.want)
			}
			got = findWords2(tst.board, tst.words)
			if !utils.MatchStringSlice(got, tst.want) {
				t.Errorf("findWords2(%v, %v) = %v, want %v", tst.board, tst.words, got, tst.want)
			}
			got = findWords3(tst.board, tst.words)
			if !utils.MatchStringSlice(got, tst.want) {
				t.Errorf("findWords3(%v, %v) = %v, want %v", tst.board, tst.words, got, tst.want)
			}
		})
	}
}

func TestFindWords(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		{
			"empty_board",
			[][]byte{},
			[]string{"a"},
			[]string{},
		},
		{
			"empty_words",
			[][]byte{{'a'}},
			[]string{},
			[]string{},
		},
		{
			"single_cell_match",
			[][]byte{{'a'}},
			[]string{"a"},
			[]string{"a"},
		},
		{
			"single_cell_no_match",
			[][]byte{{'a'}},
			[]string{"b"},
			[]string{},
		},
		{
			"all_same_chars",
			[][]byte{{'a', 'a'}, {'a', 'a'}},
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa"},
			[]string{"a", "aa", "aaa", "aaaa"},
		},
		{
			"word_longer_than_board",
			[][]byte{{'a', 'b'}, {'c', 'd'}},
			[]string{"abcde"},
			[]string{},
		},
		{
			"same_word_multiple_paths",
			[][]byte{{'a', 'b'}, {'b', 'a'}},
			[]string{"aba"},
			[]string{"aba"},
		},
		{
			"visit_same_cell_twice",
			[][]byte{{'a', 'b'}, {'c', 'd'}},
			[]string{"aba", "abcba"},
			[]string{},
		},
		{
			"word_at_boundary",
			[][]byte{{'x', 'x', 'x'}, {'x', 'a', 'x'}, {'x', 'x', 'x'}},
			[]string{"a", "xa", "ax", "xax"},
			[]string{"a", "xa", "ax", "xax"},
		},
		{
			"crossing_paths",
			[][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}},
			[]string{"abc", "cfi", "beh", "adg", "aei"},
			[]string{"abc", "cfi", "beh", "adg"},
		},
		{
			"snake_path",
			[][]byte{{'a', 'b', 'c'}, {'f', 'e', 'd'}, {'g', 'h', 'i'}},
			[]string{"abcdefghi", "abcdefgh", "ihgfedcba"},
			[]string{"abcdefghi", "abcdefgh", "ihgfedcba"},
		},
		{
			"single_row_repeating",
			[][]byte{{'a', 'b', 'a', 'b', 'a', 'b'}},
			[]string{"abab", "ababa", "ababab", "bababa", "aaaa", "bbbb"},
			[]string{"abab", "ababa", "ababab", "bababa"},
		},
		{
			"large_board_small_words",
			[][]byte{{'a', 'b', 'c', 'd', 'e'}, {'f', 'g', 'h', 'i', 'j'}, {'k', 'l', 'm', 'n', 'o'}, {'p', 'q', 'r', 's', 't'}, {'u', 'v', 'w', 'x', 'y'}},
			[]string{"a", "m", "y", "abc", "mno", "vwx", "fghij", "klmno", "abcde", "fgh", "rst"},
			[]string{"a", "m", "y", "abc", "mno", "vwx", "fghij", "klmno", "abcde", "fgh", "rst"},
		},
		{
			"zigzag_path",
			[][]byte{{'a', 'x', 'x', 'x'}, {'b', 'c', 'x', 'x'}, {'x', 'd', 'e', 'x'}, {'x', 'x', 'f', 'g'}},
			[]string{"abcdefg", "gfedcba", "abcde", "cdefg", "abcfg"},
			[]string{"abcdefg", "gfedcba", "abcde", "cdefg"},
		},
		{
			"prefix_overlap_words",
			[][]byte{{'o', 'a', 'b', 'n'}, {'o', 't', 'a', 'e'}, {'a', 'h', 'k', 'r'}, {'a', 'f', 'l', 'v'}},
			[]string{"oa", "oaa"},
			[]string{"oa", "oaa"},
		},
		{
			"word_longer_than_total_cells",
			[][]byte{{'a', 'a'}},
			[]string{"aaa"},
			[]string{},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := findWords(tst.board, tst.words)
			if !utils.MatchStringSlice(got, tst.want) {
				t.Errorf("findWords(%v, %v) = %v, want %v", tst.board, tst.words, got, tst.want)
			}
			got = findWords2(tst.board, tst.words)
			if !utils.MatchStringSlice(got, tst.want) {
				t.Errorf("findWords2(%v, %v) = %v, want %v", tst.board, tst.words, got, tst.want)
			}
			got = findWords3(tst.board, tst.words)
			if !utils.MatchStringSlice(got, tst.want) {
				t.Errorf("findWords3(%v, %v) = %v, want %v", tst.board, tst.words, got, tst.want)
			}
		})
	}
}
