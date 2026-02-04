package leetcode0739

func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for j < n && temperatures[j] <= temperatures[i] {
			if answer[j] == 0 {
				j = n
			} else {
				j += answer[j]
			}
		}
		if j < n {
			answer[i] = j - i
		}
	}
	return answer
}
