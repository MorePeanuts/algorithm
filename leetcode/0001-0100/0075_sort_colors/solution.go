// Package leetcode0075 solves LeetCode 75. Sort Colors
package leetcode0075

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	pivot := 1
	for i := left; i < right+1; {
		switch {
		case nums[i] < pivot:
			nums[left], nums[i] = nums[i], nums[left]
			left++
			// Between left and i, there can only be pivot.
			i++
		case nums[i] == pivot:
			i++
		case nums[i] > pivot:
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
	}
}
