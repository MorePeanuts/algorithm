// Package leetcode0033 solves LeetCode 33. Search in Rotated Sorted Array
package leetcode0033

func search(nums []int, target int) int {
	if nums[0] < nums[len(nums)-1] {
		return binarySearch(nums, target)
	}
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[len(nums)-1] {
			l = m + 1
		} else {
			r = m
		}
	}
	if target < nums[0] {
		res := binarySearch(nums[l:], target)
		if res == -1 {
			return -1
		}
		return res + l
	} else if target == nums[0] {
		return 0
	} else {
		return binarySearch(nums[:l], target)
	}
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return -1
}
