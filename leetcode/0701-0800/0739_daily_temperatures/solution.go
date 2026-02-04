// Package leetcode0739 solves LeetCode 739. Daily Temperatures
package leetcode0739

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, t := range temperatures {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if temperatures[top] >= t {
				break
			}
			answer[top] = i - top
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return answer
}
