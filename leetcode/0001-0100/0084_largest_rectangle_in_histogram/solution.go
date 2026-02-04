// Package leetcode0084 solves LeetCode 84. Largest Rectangle in Histogram
package leetcode0084

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	lStack := make([]int, 0)
	rStack := make([]int, 0)
	for i := range n {
		for len(lStack) > 0 {
			top := lStack[len(lStack)-1]
			if heights[top] <= heights[i] {
				break
			}
			right[top] = i
			lStack = lStack[:len(lStack)-1]
		}
		lStack = append(lStack, i)
		for len(rStack) > 0 {
			top := rStack[len(rStack)-1]
			if heights[top] <= heights[n-1-i] {
				break
			}
			left[top] = n - 1 - i
			rStack = rStack[:len(rStack)-1]
		}
		rStack = append(rStack, n-1-i)
	}
	for _, i := range lStack {
		right[i] = n
	}
	for _, i := range rStack {
		left[i] = -1
	}
	maxArea := heights[0]
	for i, height := range heights {
		area := height * (right[i] - left[i] - 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
