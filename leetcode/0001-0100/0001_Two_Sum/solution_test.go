package solution

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{"case 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"case 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"case 3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			actual := twoSum(tst.nums, tst.target)
			if actual[0] != tst.expected[0] || actual[1] != tst.expected[1] {
				t.Errorf("twoSum(%v, %v) = %v, want %v", tst.nums, tst.target, actual, tst.expected)
			}
		})
	}
}
