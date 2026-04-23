// Package leetcode0215 solves LeetCode 215. Kth Largest Element in an Array
package leetcode0215

import "container/heap"

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := IntMaxHeap(nums)
	heap.Init(&h)
	var res int
	for range k {
		res = heap.Pop(&h).(int)
	}
	return res
}
