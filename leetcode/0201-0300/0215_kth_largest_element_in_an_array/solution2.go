package leetcode0215

import "container/heap"

type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest2(nums []int, k int) int {
	h := IntMinHeap(nums[:k])
	heap.Init(&h)
	for _, num := range nums[k:] {
		if num > h[0] {
			heap.Pop(&h)
			heap.Push(&h, num)
		}
	}
	return h[0]
}
