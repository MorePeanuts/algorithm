// Package leetcode0217 solves LeetCode 217. Contains Duplicate
package leetcode0217

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, num := range nums {
		_, exists := set[num]
		if exists {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
