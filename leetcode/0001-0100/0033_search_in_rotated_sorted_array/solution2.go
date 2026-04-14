package leetcode0033

func search2(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[m] >= nums[l] {
			if target >= nums[l] && target < nums[m] {
				r = m
			} else {
				l = m + 1
			}
		} else {
			if target < nums[m] || target > nums[r-1] {
				r = m
			} else {
				l = m + 1
			}
		}
	}
	return -1
}
