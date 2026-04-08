package leetcode0209

func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	res := n + 1
	left := 0
	sum := 0

	for right, num := range nums {
		sum += num
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res > n {
		return 0
	}
	return res
}
