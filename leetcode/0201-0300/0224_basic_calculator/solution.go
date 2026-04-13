// Package leetcode0224 solves LeetCode 224. Basic Calculator
package leetcode0224

import "strconv"

func calculate(s string) int {
	infix := expr2Infix(s)
	postfix := infix2Postfix(infix)
	return evalPostfix(postfix)
}

func leftNonSpace(s string, idx int) byte {
	for i := idx - 1; i >= 0; i-- {
		if s[i] != ' ' {
			return s[i]
		}
	}
	return '('
}

func expr2Infix(expr string) []string {
	var res []string
	var l, r int
	for r < len(expr) {
		currByte := expr[r]
		if currByte >= '0' && currByte <= '9' {
			r++
		} else {
			if l != r {
				res = append(res, expr[l:r])
			}
			if currByte == '-' && leftNonSpace(expr, r) == '(' {
				res = append(res, "0")
			}
			if currByte != ' ' {
				res = append(res, expr[r:r+1])
			}
			l, r = r+1, r+1
		}
	}
	if l != len(expr) {
		res = append(res, expr[l:r])
	}
	return res
}

func infix2Postfix(infix []string) []string {
	stack := make([]string, 0)
	postfix := make([]string, 0)
	for _, token := range infix {
		if _, err := strconv.Atoi(token); err == nil {
			postfix = append(postfix, token)
		} else {
			switch token {
			case ")":
				for len(stack) != 0 {
					op := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					if op == "(" {
						break
					}
					postfix = append(postfix, op)
				}
			case "(":
				stack = append(stack, token)
			default:
				for len(stack) > 0 {
					topOp := stack[len(stack)-1]
					if topOp == "(" {
						break
					}
					postfix = append(postfix, topOp)
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, token)
			}
		}
	}
	for len(stack) != 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return postfix
}

func evalPostfix(postfix []string) int {
	stack := make([]int, 0)
	for _, token := range postfix {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			l, r := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			res := 0
			switch token {
			case "+":
				res = l + r
			case "-":
				res = l - r
			}
			stack = append(stack, res)
		}
	}
	return stack[0]
}
