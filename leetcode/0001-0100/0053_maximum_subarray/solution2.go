package leetcode0053

// O(N log(N)), divide and conquer
func maxSubArray2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	center := n / 2
	maxLeftSum := maxSubArray2(nums[:center])
	maxRightSum := maxSubArray2(nums[center:])
	maxLeftBorderSum := nums[center-1]
	maxRightBorderSum := nums[center]
	leftBorderSum := 0
	for i := center - 1; i >= 0; i-- {
		leftBorderSum += nums[i]
		if leftBorderSum > maxLeftBorderSum {
			maxLeftBorderSum = leftBorderSum
		}
	}
	rightBorderSum := 0
	for i := center; i < n; i++ {
		rightBorderSum += nums[i]
		if rightBorderSum > maxRightBorderSum {
			maxRightBorderSum = rightBorderSum
		}
	}
	return max(maxLeftSum, maxRightSum, maxLeftBorderSum+maxRightBorderSum)
}
