// Package leetcode0167 solves LeetCode 167. Two Sum II - Input Array Is Sorted
package leetcode0167

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	res := make([]int, 2)
	for left < right {
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			res[0], res[1] = left+1, right+1
			break
		}
	}
	return res
}
