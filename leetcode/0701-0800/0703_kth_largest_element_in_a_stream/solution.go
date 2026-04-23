// Package leetcode0703 solves LeetCode 703. Kth Largest Element in a Stream
package leetcode0703

import "container/heap"

type elements []int

func (es elements) Len() int {
	return len(es)
}

func (es elements) Less(i, j int) bool {
	return es[i] < es[j]
}

func (es elements) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

func (es *elements) Push(x any) {
	*es = append(*es, x.(int))
}

func (es *elements) Pop() any {
	old := *es
	x := old[len(old)-1]
	*es = old[0 : len(old)-1]
	return x
}

type KthLargest struct {
	nums elements
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	numbers := elements(nums)
	heap.Init(&numbers)
	for range len(numbers) - k {
		heap.Pop(&numbers)
	}
	return KthLargest{numbers, k}
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) < this.k {
		heap.Push(&this.nums, val)
	} else if this.nums[0] < val {
		heap.Push(&this.nums, val)
		heap.Pop(&this.nums)
	}
	return this.nums[0]
}
