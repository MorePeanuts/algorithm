// Package leetcode0704 solves LeetCode 704. Binary Search
package leetcode0704

// Recursion
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	center := n / 2
	if nums[center] == target {
		return center
	} else if target < nums[center] {
		return search(nums[:center], target)
	} else {
		right := search(nums[center+1:], target)
		if right == -1 {
			return right
		} else {
			return center + right + 1
		}
	}
}
