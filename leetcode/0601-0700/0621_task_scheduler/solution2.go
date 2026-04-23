package leetcode0621

import "slices"

func leastInterval2(tasks []byte, n int) int {
	cnt := make([]int, 26)
	for _, task := range tasks {
		cnt[task-'A']++
	}
	slices.SortFunc(cnt, func(a, b int) int {
		return b - a
	})
	res := (n+1)*(cnt[0]-1) + 1
	for _, c := range cnt[1:] {
		if c == cnt[0] {
			res++
		}
	}
	return max(res, len(tasks))
}
