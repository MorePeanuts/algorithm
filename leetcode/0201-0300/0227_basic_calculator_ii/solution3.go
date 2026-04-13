package leetcode0227

import (
	"strconv"
)

func calculate3(s string) int {
	stack := make([]int, 0)
	var lastOp byte = '+'
	var l, r int
	for r < len(s) {
		currByte := s[r]
		if currByte >= '0' && currByte <= '9' {
			r++
		} else {
			if l != r {
				num, _ := strconv.Atoi(s[l:r])
				stack = append(stack, num)
			}
			if currByte != ' ' {
				switch lastOp {
				case '-':
					stack[len(stack)-1] *= -1
				case '*':
					stack[len(stack)-2] *= stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				case '/':
					stack[len(stack)-2] /= stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
				lastOp = currByte
			}
			l, r = r+1, r+1
		}
	}
	if l != len(s) {
		num, _ := strconv.Atoi(s[l:r])
		stack = append(stack, num)
	}
	switch lastOp {
	case '-':
		stack[len(stack)-1] *= -1
	case '*':
		stack[len(stack)-2] *= stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	case '/':
		stack[len(stack)-2] /= stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	res := 0
	for _, num := range stack {
		res += num
	}
	return res
}
