package leetcode0224

import "strconv"

func calculate3(s string) int {
	nums := make([]int, 0)
	ops := make([]string, 0)
	var l, r int
	for r < len(s) {
		currByte := s[r]
		if currByte >= '0' && currByte <= '9' {
			r++
		} else {
			if l != r {
				num, _ := strconv.Atoi(s[l:r])
				nums = append(nums, num)
			}
			if currByte == '-' && leftNonSpace(s, r) == '(' {
				nums = append(nums, 0)
			}
			if currByte != ' ' {
				op := s[r : r+1]
				switch op {
				case "(":
					ops = append(ops, "(")
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
					ops = append(ops, op)
				}
			}
			l, r = r+1, r+1
		}
	}
	if l != len(s) {
		num, _ := strconv.Atoi(s[l:r])
		nums = append(nums, num)
	}
	for len(ops) > 0 {
		topOp := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		nums = evalTwoNums(nums, topOp)
	}
	return nums[0]
}
