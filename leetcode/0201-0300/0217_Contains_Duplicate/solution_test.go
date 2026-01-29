package solution

import "testing"

func TestContainsDuplicatesExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			"example_1",
			[]int{1, 2, 3, 1},
			true,
		},
		{
			"example_2",
			[]int{1, 2, 3, 4},
			false,
		},
		{
			"example_3",
			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			true,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := containsDuplicate(tst.nums)
			if got != tst.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tst.nums, got, tst.want)
			}
		})
	}
}
