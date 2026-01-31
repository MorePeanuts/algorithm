// Package leetcode0053 solves LeetCode 53. Maximum Subarray
package leetcode0053

// O(N), dynamic programming
func maxSubArray(nums []int) int {
	lastSum := 0
	thisSum := 0
	maxSum := nums[0]
	for _, num := range nums {
		thisSum += num
		if lastSum < 0 {
			thisSum -= lastSum
		}
		if thisSum > maxSum {
			maxSum = thisSum
		}
		lastSum = thisSum
	}
	return maxSum
}
