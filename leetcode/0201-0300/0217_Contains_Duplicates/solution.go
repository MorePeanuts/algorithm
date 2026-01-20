// Package solution 0217
package solution

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, num := range nums {
		_, exists := set[num]
		if exists {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
