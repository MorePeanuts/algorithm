package leetcode0227

import "strconv"

func calculate2(s string) int {
	infix := expr2IntInfix(s)
	ops1 := map[int]func(int, int) int{
		-3: func(l, r int) int { return l * r },
		-4: func(l, r int) int { return l / r },
	}
	infix = evalSimpleInfix(infix, ops1)
	ops2 := map[int]func(int, int) int{
		-1: func(l, r int) int { return l + r },
		-2: func(l, r int) int { return l - r },
	}
	infix = evalSimpleInfix(infix, ops2)
	return infix[0]
}

func expr2IntInfix(expr string) []int {
	var res []int
	var l, r int
	for r < len(expr) {
		currByte := expr[r]
		if currByte >= '0' && currByte <= '9' {
			r++
		} else {
			if l != r {
				num, _ := strconv.Atoi(expr[l:r])
				res = append(res, num)
			}
			switch currByte {
			case '+':
				res = append(res, -1)
			case '-':
				res = append(res, -2)
			case '*':
				res = append(res, -3)
			case '/':
				res = append(res, -4)
			}
			l, r = r+1, r+1
		}
	}
	if l != len(expr) {
		num, _ := strconv.Atoi(expr[l:r])
		res = append(res, num)
	}
	return res
}

func evalSimpleInfix(infix []int, ops map[int]func(int, int) int) []int {
	res := make([]int, 0)
	var i int
	for i < len(infix) {
		if op, ok := ops[infix[i]]; !ok {
			res = append(res, infix[i])
		} else {
			res[len(res)-1] = op(res[len(res)-1], infix[i+1])
			i++
		}
		i++
	}
	return res
}
