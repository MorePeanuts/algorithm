// Package leetcode0227 solves LeetCode 227. Basic Calculator II
package leetcode0227

import "strconv"

func calculate(s string) int {
	infix := expr2Infix(s)
	postfix := infix2Postfix(infix)
	return evalPostfix(postfix)
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
			for len(stack) > 0 {
				op := stack[len(stack)-1]
				if (token == "*" || token == "/") && (op == "+" || op == "-") {
					break
				}
				stack = stack[:len(stack)-1]
				postfix = append(postfix, op)
			}
			stack = append(stack, token)
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
			stack = evalTwoNums(stack, token)
		}
	}
	return stack[0]
}

func evalTwoNums(nums []int, op string) []int {
	l, r := nums[len(nums)-2], nums[len(nums)-1]
	nums = nums[:len(nums)-2]
	switch op {
	case "+":
		nums = append(nums, l+r)
	case "-":
		nums = append(nums, l-r)
	case "*":
		nums = append(nums, l*r)
	case "/":
		nums = append(nums, l/r)
	}
	return nums
}
