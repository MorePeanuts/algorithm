// Package leetcode0011 solves LeetCode 11. Container With Most Water
package leetcode0011

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	area := 0
	for l < r {
		if height[l] <= height[r] {
			area = max(area, (r-l)*height[l])
			i := l + 1
			for i < r && height[i] <= height[l] {
				i++
			}
			l = i
		} else {
			area = max(area, (r-l)*height[r])
			i := r - 1
			for i > l && height[i] <= height[r] {
				i--
			}
			r = i
		}
	}
	return area
}
