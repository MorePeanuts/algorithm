// Package leetcode1046 solves LeetCode 1046. Last Stone Weight
package leetcode1046

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := IntHeap(stones)
	heap.Init(&h)
	for len(h) > 1 {
		y := heap.Pop(&h).(int)
		x := heap.Pop(&h).(int)
		if y > x {
			heap.Push(&h, y-x)
		}
	}
	if len(h) == 0 {
		return 0
	}
	return h[0]
}
