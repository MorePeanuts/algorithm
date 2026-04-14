// Package leetcode0153 solves LeetCode 153. Find Minimum in Rotated Sorted Array
package leetcode0153

func findMin(nums []int) int {
	rightNum := nums[len(nums)-1]
	if nums[0] < rightNum {
		return nums[0]
	}
	l, r := 0, len(nums)
	m := 0
	for l < r {
		m = l + (r-l)/2
		if nums[m] > rightNum {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
