// Package solution 0001
package solution

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)
	var res [2]int
	for i, num := range nums {
		diff := target - num
		idx, ok := hashTable[diff]
		if ok {
			res[0], res[1] = idx, i
			break
		}
		hashTable[num] = i
	}
	return res[:]
}
