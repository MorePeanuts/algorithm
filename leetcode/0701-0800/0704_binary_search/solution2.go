package leetcode0704

// Loop
func search2(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
