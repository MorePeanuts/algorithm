package leetcode0042

func trap2(height []int) int {
	n := len(height)
	left := 0
	right := n - 1
	prefix := height[left]
	suffix := height[right]

	res := 0
	for right-left > 1 {
		if height[left] <= height[right] {
			res += max(prefix-height[left+1], 0)
			left++
			prefix = max(prefix, height[left])
		} else {
			res += max(suffix-height[right-1], 0)
			right--
			suffix = max(suffix, height[right])
		}
	}
	return res
}
