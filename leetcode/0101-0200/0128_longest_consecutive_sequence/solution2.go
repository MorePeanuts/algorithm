package leetcode0128

func longestConsecutive2(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	maxLen := 0
	for num := range set {
		// Skip if it is not the starting point of a continuous sequence.
		if _, exist := set[num-1]; exist {
			continue
		}
		length := 1
		for {
			if _, exist := set[num+length]; !exist {
				break
			}
			length++
		}
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}
