package leetcode0224

import "strconv"

func calculate2(s string) int {
	infix := expr2Infix(s)
	return evalSimplePostfix(infix)
}

func evalSimplePostfix(postfix []string) int {
	nums := make([]int, 0)
	ops := make([]string, 0)

	for _, token := range postfix {
		if num, err := strconv.Atoi(token); err == nil {
			nums = append(nums, num)
		} else {
			switch token {
			case "(":
				ops = append(ops, token)
			case ")":
				for len(ops) > 0 {
					topOp := ops[len(ops)-1]
					ops = ops[:len(ops)-1]
					if topOp == "(" {
						break
					}
					nums = evalTwoNums(nums, topOp)
				}
			default:
				for len(ops) > 0 {
					topOp := ops[len(ops)-1]
					if topOp == "(" {
						break
					}
					ops = ops[:len(ops)-1]
					nums = evalTwoNums(nums, topOp)
				}
				ops = append(ops, token)
			}
		}
	}
	for len(ops) > 0 {
		topOp := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		nums = evalTwoNums(nums, topOp)
	}
	return nums[0]
}

func evalTwoNums(nums []int, op string) []int {
	l, r := nums[len(nums)-2], nums[len(nums)-1]
	nums = nums[:len(nums)-2]
	if op == "+" {
		nums = append(nums, l+r)
	} else {
		nums = append(nums, l-r)
	}
	return nums
}
