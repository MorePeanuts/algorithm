// Package leetcode0875 solves LeetCode 875. Koko Eating Bananas
package leetcode0875

func minEatingSpeed(piles []int, h int) int {
	var l, r int
	l, r = 1, piles[0]
	for _, pile := range piles[1:] {
		if pile > r {
			r = pile
		}
	}

	for l < r {
		m := l + (r-l)/2
		cost := 0
		for _, pile := range piles {
			if pile%m == 0 {
				cost += pile / m
			} else {
				cost += pile/m + 1
			}
		}
		if cost <= h {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
