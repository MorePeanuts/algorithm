// Package leetcode0003 solves LeetCode 3. Longest Substring Without Repeating Characters
package leetcode0003

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	res := 1
	pos := map[byte]int{s[0]: 0}
	for l, r := 0, 1; r < len(s); r++ {
		i, ok := pos[s[r]]
		if ok {
			for j := l; j < i+1; j++ {
				delete(pos, s[j])
			}
			l = i + 1
		}
		pos[s[r]] = r
		res = max(res, len(pos))
	}
	return res
}
