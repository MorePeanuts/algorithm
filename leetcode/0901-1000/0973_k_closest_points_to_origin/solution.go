// Package leetcode0973 solves LeetCode 973. K Closest Points to Origin
package leetcode0973

import "container/heap"

type PointHeap [][]int

func (h PointHeap) Len() int      { return len(h) }
func (h PointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h PointHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] < h[j][0]*h[j][0]+h[j][1]*h[j][1]
}

func (h *PointHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *PointHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := PointHeap(points)
	heap.Init(&h)
	res := make([][]int, 0)
	for range k {
		p := heap.Pop(&h).([]int)
		res = append(res, p)
	}
	return res
}
