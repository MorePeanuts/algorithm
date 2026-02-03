// Package leetcode0125 solves LeetCode 125. Valid Palindrome
package leetcode0125

import "unicode"

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		for l < len(s) && !unicode.IsDigit(rune(s[l])) && !unicode.IsLetter(rune(s[l])) {
			l++
		}
		for r >= 0 && !unicode.IsDigit(rune(s[r])) && !unicode.IsLetter(rune(s[r])) {
			r--
		}
		if l >= r {
			break
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
	}
	return true
}
