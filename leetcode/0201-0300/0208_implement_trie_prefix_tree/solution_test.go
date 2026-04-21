// Package leetcode0208 tests LeetCode 208. Implement Trie (Prefix Tree)
package leetcode0208

import "testing"

func TestTrieExamples(t *testing.T) {
	trie := Constructor()

	// Example 1 operations
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Error("Search(apple) = false, want true")
	}
	if trie.Search("app") {
		t.Error("Search(app) = true, want false")
	}
	if !trie.StartsWith("app") {
		t.Error("StartsWith(app) = false, want true")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Error("Search(app) after insert = false, want true")
	}
}

func TestTrie(t *testing.T) {
	tests := []struct {
		name string
		// Define test steps as a sequence of operations
		ops []struct {
			op   string // "insert", "search", "startsWith"
			arg  string
			want bool
		}
	}{
		{
			name: "empty_trie",
			ops: []struct {
				op   string
				arg  string
				want bool
			}{
				{"search", "a", false},
				{"startsWith", "a", false},
			},
		},
		{
			name: "single_char",
			ops: []struct {
				op   string
				arg  string
				want bool
			}{
				{"insert", "a", false},
				{"search", "a", true},
				{"search", "aa", false},
				{"startsWith", "a", true},
				{"startsWith", "aa", false},
			},
		},
		{
			name: "prefix_is_word",
			ops: []struct {
				op   string
				arg  string
				want bool
			}{
				{"insert", "app", false},
				{"insert", "apple", false},
				{"insert", "application", false},
				{"search", "app", true},
				{"search", "apple", true},
				{"search", "application", true},
				{"search", "appl", false},
				{"startsWith", "app", true},
				{"startsWith", "appl", true},
				{"startsWith", "apple", true},
			},
		},
		{
			name: "multiple_inserts_same_word",
			ops: []struct {
				op   string
				arg  string
				want bool
			}{
				{"insert", "test", false},
				{"insert", "test", false},
				{"search", "test", true},
			},
		},
		{
			name: "search_nonexistent",
			ops: []struct {
				op   string
				arg  string
				want bool
			}{
				{"insert", "hello", false},
				{"search", "world", false},
				{"search", "hell", false},
				{"search", "hellow", false},
			},
		},
		{
			name: "startsWith_nonexistent",
			ops: []struct {
				op   string
				arg  string
				want bool
			}{
				{"insert", "hello", false},
				{"startsWith", "he", true},
				{"startsWith", "hel", true},
				{"startsWith", "hell", true},
				{"startsWith", "hello", true},
				{"startsWith", "hellow", false},
				{"startsWith", "a", false},
			},
		},
		{
			name: "all_letters",
			ops: []struct {
				op   string
				arg  string
				want bool
			}{
				{"insert", "abcdefghijklmnopqrstuvwxyz", false},
				{"search", "abcdefghijklmnopqrstuvwxyz", true},
				{"startsWith", "abcdefghijklmnopqrstuvwxyz", true},
				{"startsWith", "abcdefghijklmnopqrstuvwxy", true},
				{"search", "abcdefghijklmnopqrstuvwxy", false},
			},
		},
		{
			name: "overlapping_words",
			ops: []struct {
				op   string
				arg  string
				want bool
			}{
				{"insert", "a", false},
				{"insert", "aa", false},
				{"insert", "aaa", false},
				{"insert", "aaaa", false},
				{"search", "a", true},
				{"search", "aa", true},
				{"search", "aaa", true},
				{"search", "aaaa", true},
				{"search", "aaaaa", false},
				{"startsWith", "a", true},
				{"startsWith", "aa", true},
				{"startsWith", "aaa", true},
				{"startsWith", "aaaa", true},
				{"startsWith", "aaaaa", false},
			},
		},
		{
			name: "different_branches",
			ops: []struct {
				op   string
				arg  string
				want bool
			}{
				{"insert", "cat", false},
				{"insert", "car", false},
				{"insert", "dog", false},
				{"insert", "door", false},
				{"search", "cat", true},
				{"search", "car", true},
				{"search", "dog", true},
				{"search", "door", true},
				{"search", "ca", false},
				{"search", "do", false},
				{"startsWith", "ca", true},
				{"startsWith", "do", true},
				{"startsWith", "c", true},
				{"startsWith", "d", true},
				{"startsWith", "e", false},
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			trie := Constructor()
			for _, op := range tst.ops {
				switch op.op {
				case "insert":
					trie.Insert(op.arg)
				case "search":
					if got := trie.Search(op.arg); got != op.want {
						t.Errorf("Search(%q) = %v, want %v", op.arg, got, op.want)
					}
				case "startsWith":
					if got := trie.StartsWith(op.arg); got != op.want {
						t.Errorf("StartsWith(%q) = %v, want %v", op.arg, got, op.want)
					}
				}
			}
		})
	}
}
