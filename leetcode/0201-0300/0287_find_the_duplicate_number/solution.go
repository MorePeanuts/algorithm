// Package leetcode0287 solves LeetCode 287. Find the Duplicate Number
package leetcode0287

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	fast = 0
	for nums[slow] != nums[fast] {
		slow = nums[slow]
		fast = nums[fast]
	}
	return nums[slow]
}
