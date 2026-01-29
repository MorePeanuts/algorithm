package utils

import (
	"reflect"
	"sort"
)

// MatchTwo2dStringSlice is used to judge whether two two-dimensional string slices are identical,
// ignoring the order of strings within the inner slices and the order of inner slices within the
// outer slice.
func MatchTwo2dStringSlice(got, want [][]string) bool {
	if len(got) != len(want) {
		return false
	}

	// 1. Copy data (to avoid modifying the original data)
	// 2. Sort the inner slice
	// 3. Sort the outer slice
	normalize := func(input [][]string) [][]string {
		output := make([][]string, len(input))
		for i, group := range input {
			groupCopy := make([]string, len(group))
			copy(groupCopy, group)
			sort.Strings(groupCopy)
			output[i] = groupCopy
		}
		sort.Slice(output, func(i, j int) bool {
			a, b := output[i], output[j]
			lenA, lenB := len(a), len(b)
			minLen := min(lenA, lenB)
			for k := range minLen {
				if a[k] != b[k] {
					return a[k] < b[k]
				}
			}
			return lenA < lenB
		})
		return output
	}

	sortedGot := normalize(got)
	sortedWant := normalize(want)

	return reflect.DeepEqual(sortedGot, sortedWant)
}
