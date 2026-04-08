package leetcode0042

func trap(height []int) int {
	n := len(height)
	prefix := make([]int, n)
	suffix := make([]int, n)

	prefix[0], suffix[n-1] = 0, 0
	for i := range n - 1 {
		prefix[i+1] = max(prefix[i], height[i])
		suffix[n-i-2] = max(suffix[n-i-1], height[n-i-1])
	}

	res := 0
	for i, h := range height {
		board := min(prefix[i], suffix[i])
		res += max(board-h, 0)
	}
	return res
}
