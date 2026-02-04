// Package leetcode0150 solves LeetCode 150. Evaluate Reverse Polish Notation
package leetcode0150

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		if i, error := strconv.Atoi(token); error == nil {
			stack = append(stack, i)
		} else {
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			num := 0
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				num = num1 + num2
			case "-":
				num = num1 - num2
			case "*":
				num = num1 * num2
			case "/":
				num = num1 / num2
			}
			stack = append(stack, num)
		}
	}
	return stack[0]
}
