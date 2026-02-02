// Package leetcode0049 solves LeetCode 49. Group Anagrams
package leetcode0049

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]byte][]string)
	for _, str := range strs {
		key := [26]byte{}
		for _, c := range str {
			key[c-'a'] += byte(c)
		}
		anagrams[key] = append(anagrams[key], str)
	}
	res := make([][]string, 0, len(anagrams))
	for _, lst := range anagrams {
		res = append(res, lst)
	}
	return res
}
