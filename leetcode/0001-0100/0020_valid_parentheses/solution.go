// Package leetcode0020 solves LeetCode 20. Valid Parentheses
package leetcode0020

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]rune, 0, len(s))
	// (: 40, ): 41, {: 123, }: 125, [: 91, ]: 93
	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
			if len(stack) != 0 && (r-stack[len(stack)-1] == 1 || r-stack[len(stack)-1] == 2) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
