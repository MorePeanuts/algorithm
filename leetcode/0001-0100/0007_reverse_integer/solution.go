// Package leetcode0007 solves LeetCode 7. Reverse Integer
package leetcode0007

import "math"

func reverse(x int) int {
	var res int
	for x != 0 {
		last := x % 10
		if res > math.MaxInt32/10 || (res*10 > math.MaxInt32-last) {
			return 0
		}
		if res < math.MinInt32/10 || (res*10 < math.MinInt32-last) {
			return 0
		}
		res = res*10 + last
		x /= 10
	}

	return res
}
