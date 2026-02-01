// Package leetcode0238 solves LeetCode 238. Product of Array Except Self
package leetcode0238

func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftAccProd := make([]int, n)
	rightAccProd := make([]int, n)
	leftAccProd[0] = nums[0]
	rightAccProd[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		leftAccProd[i] = nums[i] * leftAccProd[i-1]
		rightAccProd[n-i-1] = nums[n-i-1] * rightAccProd[n-i]
	}
	res := make([]int, n)
	res[0] = rightAccProd[1]
	res[n-1] = leftAccProd[n-2]
	for i := 1; i < n-1; i++ {
		switch i {
		case 0:
			res[i] = rightAccProd[i+1]
		case n - 1:
			res[i] = leftAccProd[i-1]
		default:
			res[i] = rightAccProd[i+1] * leftAccProd[i-1]
		}
	}
	return res
}
