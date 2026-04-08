// Package leetcode0042 solves LeetCode 42. Trapping Rain Water
package leetcode0042

func trap3(height []int) int {
	n := len(height)
	stack := make([]int, 0, n)
	res := 0
	top := 0
	for i := range n {
		// Maintain a monotonic decreasing stack
		for len(stack) > 0 {
			top = stack[len(stack)-1]
			// Pop until the height corresponding to the top of the stack is greater than or equal to the current height.
			if height[top] >= height[i] {
				break
			}
			stack = stack[:len(stack)-1]
		}
		// If the stack is not empty, then the right side is lower
		h := height[i]
		// If the stack is empty, then the left side is lower
		if len(stack) == 0 {
			h = height[top]
		}

		// Fill the lower position to the lower side
		for j := top + 1; j < i; j++ {
			res += h - height[j]
			height[j] = h
		}
		stack = append(stack, i)
	}
	return res
}
