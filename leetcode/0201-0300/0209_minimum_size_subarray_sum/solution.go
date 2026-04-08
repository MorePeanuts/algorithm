// Package leetcode0209 solves LeetCode 209. Minimum Size Subarray Sum
package leetcode0209

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left := 0
	right := 0
	sum := nums[0]
	res := n + 1

	for right < n && left <= right {
		if sum < target {
			right++
			if right < n {
				sum += nums[right]
			}
		} else {
			res = min(res, right-left+1)
			left++
			if left <= right {
				sum -= nums[left-1]
			}
		}
	}
	if res == n+1 {
		return 0
	}
	return res
}
