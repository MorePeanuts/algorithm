package leetcode0053

// O(N^2)
func maxSubArray3(nums []int) int {
	thisSum, maxSum := 0, nums[0]
	n := len(nums)
	for i := range n {
		thisSum = 0
		for j := i; j < n; j++ {
			thisSum += nums[j]
			if thisSum > maxSum {
				maxSum = thisSum
			}
		}
	}
	return maxSum
}
