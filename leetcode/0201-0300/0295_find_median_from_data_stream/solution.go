// Package leetcode0295 solves LeetCode 295. Find Median from Data Stream
package leetcode0295

import "container/heap"

type (
	MaxHeap []int
	MinHeap []int
)

func (h MaxHeap) Len() int           { return len(h) }
func (h MinHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }

func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h *MinHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.minHeap) == len(this.maxHeap) {
		if len(this.maxHeap) > 0 && num < this.maxHeap[0] {
			heap.Push(&this.minHeap, this.maxHeap[0])
			heap.Pop(&this.maxHeap)
			heap.Push(&this.maxHeap, num)
		} else {
			heap.Push(&this.minHeap, num)
		}
	} else {
		if len(this.minHeap) > 0 && num <= this.minHeap[0] {
			heap.Push(&this.maxHeap, num)
		} else {
			heap.Push(&this.maxHeap, this.minHeap[0])
			heap.Pop(&this.minHeap)
			heap.Push(&this.minHeap, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	right := this.minHeap[0]
	if len(this.minHeap) == len(this.maxHeap) {
		left := this.maxHeap[0]
		return float64(right+left) / 2
	} else {
		return float64(right)
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
