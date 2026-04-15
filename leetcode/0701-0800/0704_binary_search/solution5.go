package leetcode0704

func search5(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if target <= nums[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	if target == nums[l] {
		return l
	}
	return -1
}
