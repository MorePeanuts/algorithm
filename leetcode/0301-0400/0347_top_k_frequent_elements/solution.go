// Package leetcode0347 solves LeetCode 347. Top K Frequent Elements
package leetcode0347

import "container/heap"

type PairHeap [][2]int

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *PairHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num] += 1
	}
	pairs := make([][2]int, 0, len(freq))
	for k, v := range freq {
		pairs = append(pairs, [2]int{k, v})
	}
	hp := PairHeap(pairs)
	heap.Init(&hp)
	res := make([]int, k)
	for i := range k {
		res[i] = heap.Pop(&hp).([2]int)[0]
	}
	return res
}
