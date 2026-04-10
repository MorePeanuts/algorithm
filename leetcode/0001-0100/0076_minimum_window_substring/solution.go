// Package leetcode0076 solves LeetCode 76. Minimum Window Substring
package leetcode0076

func minWindow(s string, t string) string {
	count := make(map[byte]int)
	miss := make(map[byte]bool)

	for i := range t {
		count[t[i]]++
		miss[t[i]] = true
	}
	l, r := 0, 0
	res := ""
	minLen := len(s) + 1
	for r = range s {
		if _, ok := count[s[r]]; ok {
			count[s[r]]--
			if count[s[r]] <= 0 {
				delete(miss, s[r])
			}
		}
		for len(miss) == 0 {
			if r+1-l < minLen {
				res = s[l : r+1]
				minLen = r + 1 - l
			}
			if _, ok := count[s[l]]; ok {
				count[s[l]]++
				if count[s[l]] > 0 {
					miss[s[l]] = true
				}
			}
			l++
		}
	}
	return res
}
