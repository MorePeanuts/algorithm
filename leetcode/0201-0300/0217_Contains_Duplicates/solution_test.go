package solution

import "testing"

func TestContainsDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"true case 1", []int{1, 2, 3, 1}, true},
		{"false case 1", []int{1, 2, 3, 4}, false},
		{"true case 2", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			actual := containsDuplicate(tst.nums)
			if actual != tst.expected {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tst.nums, actual, tst.expected)
			}
		})
	}
}
