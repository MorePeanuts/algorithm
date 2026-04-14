package leetcode0704

// Two-way comparison
func search3(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if target < nums[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	if l > 0 && target == nums[l-1] {
		return l - 1
	} else {
		return -1
	}
}
