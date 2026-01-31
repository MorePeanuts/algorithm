package leetcode0704

// Two-way comparison
func search3(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right-1 {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	if left < right && nums[left] == target {
		return left
	} else {
		return -1
	}
}
