package solution

import (
	"reflect"
	"sort"
	"testing"
)

// TODO: move this utils function to somewhere else and use generics.
func matchTwo2dStringSlice(got, want [][]string) bool {
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

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			"normal",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			"empty string",
			[]string{""},
			[][]string{{""}},
		},
		{
			"single string",
			[]string{"a"},
			[][]string{{"a"}},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := groupAnagrams(tst.strs)
			if !matchTwo2dStringSlice(got, tst.expected) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.expected)
			}

			got = groupAnagrams2(tst.strs)
			if !matchTwo2dStringSlice(got, tst.expected) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.expected)
			}

			got = groupAnagrams3(tst.strs)
			if !matchTwo2dStringSlice(got, tst.expected) {
				t.Errorf("groupAnagrams(%v) = %v, want %v", tst.strs, got, tst.expected)
			}
		})
	}
}
