package utils

import (
	"math"
	"slices"
	"sort"
)

// Float64AlmostEqual checks if two float64 values are almost equal within a small epsilon.
func Float64AlmostEqual(a, b float64) bool {
	const epsilon = 1e-9
	return math.Abs(a-b) < epsilon
}

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

	return slices.EqualFunc(sortedGot, sortedWant, slices.Equal)
}

// MatchIntSlice compares two integer slices ignoring order.
// Returns true if both slices contain the same elements (with same multiplicities).
func MatchIntSlice(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	gotCopy := make([]int, len(got))
	wantCopy := make([]int, len(want))
	copy(gotCopy, got)
	copy(wantCopy, want)
	sort.Ints(gotCopy)
	sort.Ints(wantCopy)
	return slices.Equal(gotCopy, wantCopy)
}

// MatchStringSlice compares two string slices ignoring order.
// Returns true if both slices contain the same elements (with same multiplicities).
func MatchStringSlice(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	gotCopy := make([]string, len(got))
	wantCopy := make([]string, len(want))
	copy(gotCopy, got)
	copy(wantCopy, want)
	sort.Strings(gotCopy)
	sort.Strings(wantCopy)
	return slices.Equal(gotCopy, wantCopy)
}

// MatchTwo2dIntSlice is used to judge whether two two-dimensional int slices are identical,
// ignoring the order of ints within the inner slices and the order of inner slices within the
// outer slice.
func MatchTwo2dIntSlice(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}

	// 1. Copy data (to avoid modifying the original data)
	// 2. Sort the inner slice
	// 3. Sort the outer slice
	normalize := func(input [][]int) [][]int {
		output := make([][]int, len(input))
		for i, group := range input {
			groupCopy := make([]int, len(group))
			copy(groupCopy, group)
			sort.Ints(groupCopy)
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

	return slices.EqualFunc(sortedGot, sortedWant, slices.Equal)
}

// MatchTwo2dIntSlicePreserveInner is used to judge whether two two-dimensional int slices are identical,
// preserving the order of ints within the inner slices but ignoring the order of inner slices within the
// outer slice. Useful for comparing slices of points where inner slice order matters.
func MatchTwo2dIntSlicePreserveInner(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}

	// 1. Copy data (to avoid modifying the original data)
	// 2. Sort the outer slice without modifying inner slices
	normalize := func(input [][]int) [][]int {
		output := make([][]int, len(input))
		for i, group := range input {
			groupCopy := make([]int, len(group))
			copy(groupCopy, group)
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

	return slices.EqualFunc(sortedGot, sortedWant, slices.Equal)
}

// EqualTrees checks if two binary trees are equal in structure and values.
func EqualTrees(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return EqualTrees(t1.Left, t2.Left) && EqualTrees(t1.Right, t2.Right)
}
