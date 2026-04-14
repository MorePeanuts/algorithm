// Package leetcode0074 solves LeetCode 74. Search a 2D Matrix
package leetcode0074

func searchMatrix(matrix [][]int, target int) bool {
	firstNums := make([]int, 0, len(matrix))
	for _, line := range matrix {
		firstNums = append(firstNums, line[0])
	}
	pos, ok := binarySearchPos(firstNums, target)
	if !ok {
		if pos == -1 {
			return false
		}
		_, ok = binarySearchPos(matrix[pos], target)
	}
	return ok
}

func binarySearchPos(nums []int, target int) (pos int, ok bool) {
	l, r := 0, len(nums)
	m := 0
	for l < r {
		m = l + (r-l)/2
		if target < nums[m] {
			r = m
		} else if target > nums[m] {
			l = m + 1
		} else {
			return m, true
		}
	}
	// nums[pos] < target < nums[pos+1]
	return l - 1, false
}
