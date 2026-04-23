// Package leetcode0973 solves LeetCode 973. K Closest Points to Origin
package leetcode0973

import (
	"fmt"
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func TestKClosestExamples(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{
			"example_1",
			[][]int{{1, 3}, {-2, 2}},
			1,
			[][]int{{-2, 2}},
		},
		{
			"example_2",
			[][]int{{3, 3}, {5, -1}, {-2, 4}},
			2,
			[][]int{{3, 3}, {-2, 4}},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := kClosest(tst.points, tst.k)
			if !utils.MatchTwo2dIntSlicePreserveInner(got, tst.want) {
				t.Errorf("kClosest(%v, %d) = %v, want %v", tst.points, tst.k, got, tst.want)
			}
		})
	}
}

func TestKClosest(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		k      int
		// For tests with unique k closest points, want specifies the exact expected set.
		// For tests with multiple valid answers, want is nil and checkValid should be used.
		want       [][]int
		checkValid func(got [][]int) bool // Optional custom validation function
	}{
		{"single_point", [][]int{{0, 0}}, 1, [][]int{{0, 0}}, nil},
		{"k_equals_length", [][]int{{1, 2}, {3, 4}, {5, 6}}, 3, [][]int{{1, 2}, {3, 4}, {5, 6}}, nil},
		{
			"all_same_distance",
			[][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
			2,
			nil,
			func(got [][]int) bool {
				// All points are distance 1, so any 2 points are valid
				if len(got) != 2 {
					return false
				}
				validPoints := map[string]bool{
					"1,0":  true,
					"0,1":  true,
					"-1,0": true,
					"0,-1": true,
				}
				for _, p := range got {
					key := fmt.Sprintf("%d,%d", p[0], p[1])
					if !validPoints[key] {
						return false
					}
				}
				return true
			},
		},
		{"negative_coordinates", [][]int{{-1, -1}, {-2, -2}, {-3, -3}}, 2, [][]int{{-1, -1}, {-2, -2}}, nil},
		{"mixed_positive_negative", [][]int{{-5, 5}, {3, -3}, {-2, 2}, {1, -1}}, 3, [][]int{{1, -1}, {-2, 2}, {3, -3}}, nil},
		{"large_coordinates", [][]int{{10000, 10000}, {0, 0}, {1, 1}}, 2, [][]int{{0, 0}, {1, 1}}, nil},
		{
			"zero_in_middle",
			[][]int{{10, 10}, {0, 0}, {5, 5}, {-5, -5}},
			2,
			nil,
			func(got [][]int) bool {
				if len(got) != 2 {
					return false
				}
				// Must include [0,0], and either [5,5] or [-5,-5]
				hasZero := false
				hasOther := false
				for _, p := range got {
					if p[0] == 0 && p[1] == 0 {
						hasZero = true
					} else if (p[0] == 5 && p[1] == 5) || (p[0] == -5 && p[1] == -5) {
						hasOther = true
					}
				}
				return hasZero && hasOther
			},
		},
		{
			"same_distance_different_order",
			[][]int{{2, 2}, {1, 3}, {3, 1}, {2, -2}},
			2,
			[][]int{{2, 2}, {2, -2}},
			nil,
		},
		{"k_1_with_many_points", [][]int{{10, 20}, {1, 1}, {5, 5}, {3, 3}}, 1, [][]int{{1, 1}}, nil},
		{"all_zero_except_one", [][]int{{0, 0}, {0, 0}, {0, 0}, {1, 1}}, 3, [][]int{{0, 0}, {0, 0}, {0, 0}}, nil},
		{"descending_order", [][]int{{5, 5}, {4, 4}, {3, 3}, {2, 2}, {1, 1}}, 3, [][]int{{1, 1}, {2, 2}, {3, 3}}, nil},
		{"ascending_order", [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, 3, [][]int{{1, 1}, {2, 2}, {3, 3}}, nil},
		{"random_order", [][]int{{3, 3}, {1, 1}, {4, 4}, {2, 2}, {5, 5}}, 4, [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}, nil},
		{"points_on_x_axis", [][]int{{10, 0}, {1, 0}, {5, 0}, {3, 0}, {2, 0}}, 3, [][]int{{1, 0}, {2, 0}, {3, 0}}, nil},
		{"points_on_y_axis", [][]int{{0, 10}, {0, 1}, {0, 5}, {0, 3}, {0, 2}}, 3, [][]int{{0, 1}, {0, 2}, {0, 3}}, nil},
		{"alternating_signs", [][]int{{-1, 2}, {1, -2}, {3, -4}, {-3, 4}}, 2, [][]int{{-1, 2}, {1, -2}}, nil},
		{"max_min_coordinates", [][]int{{10000, -10000}, {-10000, 10000}, {0, 0}}, 1, [][]int{{0, 0}}, nil},
		{
			"symmetric_points",
			[][]int{{1, 2}, {-1, -2}, {2, 1}, {-2, -1}},
			2,
			nil,
			func(got [][]int) bool {
				if len(got) != 2 {
					return false
				}
				// All 4 points have the same distance (5), so any 2 are valid
				validPoints := map[string]bool{
					"1,2":   true,
					"-1,-2": true,
					"2,1":   true,
					"-2,-1": true,
				}
				for _, p := range got {
					key := fmt.Sprintf("%d,%d", p[0], p[1])
					if !validPoints[key] {
						return false
					}
				}
				return true
			},
		},
		{"same_x_different_y", [][]int{{2, 5}, {2, 1}, {2, 3}, {2, 4}, {2, 2}}, 3, [][]int{{2, 1}, {2, 2}, {2, 3}}, nil},
		{"same_y_different_x", [][]int{{5, 2}, {1, 2}, {3, 2}, {4, 2}, {2, 2}}, 3, [][]int{{1, 2}, {2, 2}, {3, 2}}, nil},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := kClosest(tst.points, tst.k)
			if tst.checkValid != nil {
				if !tst.checkValid(got) {
					t.Errorf("kClosest(%v, %d) = %v, not a valid answer", tst.points, tst.k, got)
				}
			} else {
				if !utils.MatchTwo2dIntSlicePreserveInner(got, tst.want) {
					t.Errorf("kClosest(%v, %d) = %v, want %v", tst.points, tst.k, got, tst.want)
				}
			}
		})
	}
}
