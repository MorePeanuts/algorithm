package leetcode0153

func findMin2(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] >= nums[0] {
			l = m + 1
		} else {
			r = m
		}
	}
	if l == len(nums) {
		return nums[l-1]
	}
	return nums[l]
}
