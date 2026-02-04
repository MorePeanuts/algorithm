// Package leetcode0853 solves LeetCode 853. Car Fleet
package leetcode0853

import "sort"

func carFleet(target int, position []int, speed []int) int {
	p2s := make(map[int]float64)
	for i, pos := range position {
		p2s[pos] = float64(speed[i])
	}
	sort.Slice(position, func(i, j int) bool {
		return position[i] > position[j]
	})
	arrive := float64(target-position[0]) / p2s[position[0]]
	res := 1
	for _, pos := range position[1:] {
		arr := float64(target-pos) / p2s[pos]
		if arr > arrive {
			arrive = arr
			res++
		}
	}
	return res
}
