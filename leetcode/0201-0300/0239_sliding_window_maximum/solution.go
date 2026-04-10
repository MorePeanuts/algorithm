// Package leetcode0239 solves LeetCode 239. Sliding Window Maximum
package leetcode0239

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)

	for _, num := range nums[:k] {
		for len(queue) > 0 && queue[len(queue)-1] < num {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, num)
	}

	res := []int{queue[0]}
	for i := range len(nums) - k {
		// Dequeue: nums[i]
		if nums[i] == queue[0] {
			queue = queue[1:]
		}
		// Enqueue: nums[i+k]
		for len(queue) > 0 && nums[i+k] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i+k])
		// Append res
		res = append(res, queue[0])
	}

	return res
}
