package leetcode0238

func productExceptSelf3(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	rightAcc := nums[n-1]
	res[0] = nums[0]
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i]
	}
	res[n-1] = res[n-2]
	for i := n - 2; i > 0; i-- {
		res[i] = res[i-1] * rightAcc
		rightAcc *= nums[i]
	}
	res[0] = rightAcc
	return res
}
