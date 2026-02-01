package leetcode0238

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	accProd := make([]int, n)
	res := make([]int, n)
	accProd[0] = nums[0]
	for i := 1; i < n; i++ {
		accProd[i] = accProd[i-1] * nums[i]
	}
	res[n-1] = accProd[n-2]
	rightAcc := nums[n-1]
	for i := n - 2; i > 0; i-- {
		res[i] = accProd[i-1] * rightAcc
		rightAcc *= nums[i]
	}
	res[0] = rightAcc
	return res
}
