// Package leetcode0567 solves LeetCode 567. Permutation in String
package leetcode0567

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	n := len(s1)
	var count1, count2 [26]int
	for _, r := range s1 {
		count1[r-'a'] += 1
	}
	l := 0
	for _, r := range s2[l : l+n] {
		count2[r-'a'] += 1
	}
	if count1 == count2 {
		return true
	}
	l += 1
	for l+n <= len(s2) {
		count2[s2[l-1]-'a'] -= 1
		count2[s2[l+n-1]-'a'] += 1
		if count1 == count2 {
			return true
		}
		l += 1
	}
	return false
}
