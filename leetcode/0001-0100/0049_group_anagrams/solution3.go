package leetcode0049

import (
	"sort"
	"strings"
)

func groupAnagrams3(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		sortedStr := strings.Join(chars, "")
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}
	res := make([][]string, 0, len(anagrams))
	for _, lst := range anagrams {
		res = append(res, lst)
	}
	return res
}
