// Package leetcode0424 solves LeetCode 424. Longest Repeating Character Replacement
package leetcode0424

func characterReplacement(s string, k int) int {
	var count [26]int
	maxCount := 0
	l := 0
	for r := 0; r < len(s); r++ {
		count[int(s[r]-'A')]++
		maxCount = max(maxCount, count[int(s[r]-'A')])
		if r-l+1-maxCount > k {
			count[int(s[l]-'A')]--
			l++
		}
	}
	return len(s) - l
}
