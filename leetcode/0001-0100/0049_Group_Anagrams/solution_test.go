package solution

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestGroupAnagramsExamples(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
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
			if !utils.MatchTwo2dStringSlice(got, tst.expected) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.expected)
			}

			got = groupAnagrams2(tst.strs)
			if !utils.MatchTwo2dStringSlice(got, tst.expected) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.expected)
			}

			got = groupAnagrams3(tst.strs)
			if !utils.MatchTwo2dStringSlice(got, tst.expected) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.expected)
			}
		})
	}
}
