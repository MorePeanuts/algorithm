// Package leetcode0242 solves LeetCode 242. Valid Anagram
package leetcode0242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	for _, c := range t {
		count[c]--
		if count[c] < 0 {
			return false
		}
	}
	return true
}
