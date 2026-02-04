package leetcode0853

func carFleet2(target int, position []int, speed []int) int {
	bucket := make([]bool, target)
	p2s := make(map[int]float64)
	for i, pos := range position {
		bucket[pos] = true
		p2s[pos] = float64(speed[i])
	}
	sortedPos := make([]int, 0, len(position))
	for i := target - 1; i >= 0; i-- {
		if bucket[i] {
			sortedPos = append(sortedPos, i)
		}
	}
	arrive := float64(target-sortedPos[0]) / p2s[sortedPos[0]]
	res := 1
	for _, pos := range sortedPos[1:] {
		arr := float64(target-pos) / p2s[pos]
		if arr > arrive {
			arrive = arr
			res++
		}
	}
	return res
}
