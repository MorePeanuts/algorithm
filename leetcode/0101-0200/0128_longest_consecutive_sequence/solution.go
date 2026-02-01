// Package leetcode0128 solves LeetCode 128. Longest Consecutive Sequence
package leetcode0128

func longestConsecutive(nums []int) int {
	left2right := make(map[int]int)
	right2left := make(map[int]int)
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, exist := set[num]; exist {
			continue
		}
		set[num] = struct{}{}
		leftValue, leftExist := left2right[num]
		rightValue, rightExist := right2left[num]
		// if num in left2right, not in right2left, expand left2right
		if leftExist && (!rightExist) {
			delete(left2right, num)
			left2right[num-1] = leftValue
			delete(right2left, leftValue)
			right2left[leftValue] = num - 1
		}
		// if num in right2left, not in left2right, expand right2left
		if rightExist && (!leftExist) {
			delete(right2left, num)
			right2left[num+1] = rightValue
			delete(left2right, rightValue)
			left2right[rightValue] = num + 1
		}
		// if num both in left2right and right2left, connect them
		if leftExist && rightExist {
			delete(left2right, num)
			delete(right2left, num)
			left2right[rightValue] = leftValue
			right2left[leftValue] = rightValue
		}
		// if num both not in left2right and right2left, append them
		if (!leftExist) && (!rightExist) {
			left2right[num-1] = num + 1
			right2left[num+1] = num - 1
		}
	}
	maxLen := 0
	for k, v := range left2right {
		if v-k-1 > maxLen {
			maxLen = v - k - 1
		}
	}
	return maxLen
}
