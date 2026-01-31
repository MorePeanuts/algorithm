package leetcode0053

// O(N^3)
func maxSubArray4(nums []int) int {
	thisSum, maxSum := 0, nums[0]
	n := len(nums)
	for i := range n {
		for j := i; j < n; j++ {
			thisSum = 0
			for k := i; k <= j; k++ {
				thisSum += nums[k]
			}
			if thisSum > maxSum {
				maxSum = thisSum
			}
		}
	}
	return maxSum
}
