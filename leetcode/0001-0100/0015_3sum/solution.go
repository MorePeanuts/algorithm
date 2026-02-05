// Package leetcode0015 solves LeetCode 15. 3Sum
package leetcode0015

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i, num := range nums {
		if num > 0 {
			break
		}
		if i > 0 && nums[i-1] == num {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] > -num {
				r--
			} else if nums[l]+nums[r] < -num {
				l++
			} else {
				res = append(res, []int{num, nums[l], nums[r]})
				for l < r {
					r--
					l++
					if nums[l] != nums[l-1] || nums[r] != nums[r+1] {
						break
					}
				}
			}
		}
	}
	return res
}
