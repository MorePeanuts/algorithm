package solution

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			"case 1",
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			"case 2",
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			"case 3",
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := twoSum(tst.nums, tst.target)
			if got[0] != tst.want[0] || got[1] != tst.want[1] {
				t.Errorf("twoSum(%v, %v) = %v, want %v", tst.nums, tst.target, got, tst.want)
			}
		})
	}
}
