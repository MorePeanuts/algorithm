// Package leetcode0621 solves LeetCode 621. Task Scheduler
package leetcode0621

import "container/heap"

type Task struct {
	char  byte
	count int
}

type TaskHeap []*Task

func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TaskHeap) Push(x any)        { *h = append(*h, x.(*Task)) }

func (h *TaskHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	total := len(tasks)
	done := 0
	res := 0
	count := map[byte]int{}
	for _, task := range tasks {
		count[task]++
	}
	h := make(TaskHeap, 0, len(count))
	for task, c := range count {
		h = append(h, &Task{task, c})
	}
	heap.Init(&h)
	queue := make([]*Task, n+1)
	for done < total {
		ready := queue[0]
		queue = queue[1:]
		if ready != nil {
			heap.Push(&h, ready)
		}
		if h.Len() > 0 {
			currTask := heap.Pop(&h).(*Task)
			currTask.count--
			done++
			if currTask.count > 0 {
				queue = append(queue, currTask)
			} else {
				queue = append(queue, nil)
			}
		} else {
			queue = append(queue, nil)
		}
		res++
	}
	return res
}
