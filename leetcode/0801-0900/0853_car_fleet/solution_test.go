package leetcode0853

import (
	"slices"
	"testing"
)

func TestCarFleetExamples(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		want     int
	}{
		{
			"example_1",
			12,
			[]int{10, 8, 0, 5, 3},
			[]int{2, 4, 1, 1, 3},
			3,
		},
		{
			"example_2",
			10,
			[]int{3},
			[]int{3},
			1,
		},
		{
			"example_3",
			100,
			[]int{0, 2, 4},
			[]int{4, 2, 1},
			1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := carFleet(tst.target, slices.Clone(tst.position), tst.speed)
			if got != tst.want {
				t.Errorf("carFleet(%d, %v, %v) = %d, want %d",
					tst.target, tst.position, tst.speed, got, tst.want)
			}

			got = carFleet2(tst.target, slices.Clone(tst.position), tst.speed)
			if got != tst.want {
				t.Errorf("carFleet(%d, %v, %v) = %d, want %d",
					tst.target, tst.position, tst.speed, got, tst.want)
			}
		})
	}
}

func TestCarFleet(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		want     int
	}{
		// Single car cases
		{"single_car_at_origin", 10, []int{0}, []int{1}, 1},
		{"single_car_near_target", 10, []int{9}, []int{1}, 1},
		{"single_car_fast", 100, []int{50}, []int{100}, 1},

		// Two cars cases
		{"two_cars_same_speed", 10, []int{2, 5}, []int{3, 3}, 2},
		{"two_cars_merge", 10, []int{0, 5}, []int{10, 1}, 1},
		{"two_cars_no_merge", 10, []int{0, 5}, []int{1, 10}, 2},
		{"two_cars_merge_at_target", 12, []int{10, 8}, []int{2, 4}, 1},

		// All cars become one fleet
		{"all_merge_chain", 20, []int{0, 5, 10, 15}, []int{4, 3, 2, 1}, 1},
		{"all_merge_fast_behind", 50, []int{0, 10, 20, 30}, []int{50, 40, 30, 20}, 1},

		// All cars remain separate
		{"all_separate_increasing_speed", 100, []int{10, 20, 30, 40}, []int{1, 2, 3, 4}, 4},
		{"all_separate_same_speed", 50, []int{10, 20, 30, 40}, []int{5, 5, 5, 5}, 4},

		// Mixed scenarios
		{"pairs_merge", 20, []int{0, 2, 10, 12}, []int{5, 4, 5, 4}, 2},
		{"one_catches_all", 100, []int{0, 10, 20, 30}, []int{100, 1, 1, 1}, 3},
		{"last_catches_none", 100, []int{0, 10, 20, 30}, []int{1, 2, 3, 100}, 4},

		// Edge positions
		{"car_at_zero", 10, []int{0}, []int{10}, 1},
		{"car_very_close_to_target", 1000000, []int{999999}, []int{1000000}, 1},
		{"cars_at_extreme_positions", 1000000, []int{0, 999999}, []int{1, 1000000}, 2},

		// Speed variations
		{"very_slow_car", 100, []int{50}, []int{1}, 1},
		{"very_fast_car", 100, []int{0}, []int{1000000}, 1},
		{"mixed_speeds", 20, []int{0, 5, 10}, []int{1, 5, 10}, 3},

		// Specific merge scenarios
		{"merge_exactly_at_target", 10, []int{0, 5}, []int{2, 1}, 1},
		{"merge_before_target", 100, []int{0, 50}, []int{10, 5}, 1},
		{"no_merge_parallel", 50, []int{0, 10, 20}, []int{10, 10, 10}, 3},

		// Three cars scenarios
		{"three_cars_all_merge", 30, []int{0, 10, 20}, []int{6, 4, 2}, 1},
		{"three_cars_two_merge", 30, []int{0, 10, 20}, []int{1, 4, 2}, 2},
		{"three_cars_none_merge", 30, []int{0, 10, 20}, []int{1, 2, 3}, 3},

		// Larger test cases
		{"five_cars_complex", 100, []int{0, 20, 40, 60, 80}, []int{10, 8, 6, 4, 2}, 1},
		{"five_cars_alternating", 50, []int{5, 10, 15, 20, 25}, []int{10, 1, 10, 1, 10}, 3},

		// Time-based merge scenarios
		{"slow_leader_blocks_all", 100, []int{0, 10, 20, 30, 40, 50}, []int{100, 100, 100, 100, 100, 1}, 1},
		{"fast_cars_catch_slow", 100, []int{0, 25, 50, 75}, []int{4, 3, 2, 1}, 1},

		// Boundary values
		{"minimum_target", 1, []int{0}, []int{1}, 1},
		{"large_target_small_positions", 1000000, []int{0, 1, 2}, []int{1, 1, 1}, 3},

		// Random-like distributions
		{"scattered_positions", 50, []int{5, 15, 25, 35, 45}, []int{9, 7, 5, 3, 1}, 1},
		{"reverse_order_speeds", 100, []int{10, 30, 50, 70, 90}, []int{1, 2, 3, 4, 5}, 5},

		// Additional edge cases
		{"two_cars_exact_same_arrival", 10, []int{0, 5}, []int{1, 2}, 2},
		{"many_cars_same_arrival_time", 100, []int{0, 50, 75, 90}, []int{1, 2, 4, 10}, 4},
		{"cascade_merge", 100, []int{0, 20, 40, 60, 80}, []int{5, 4, 3, 2, 1}, 1},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := carFleet(tst.target, slices.Clone(tst.position), tst.speed)
			if got != tst.want {
				t.Errorf("carFleet(%d, %v, %v) = %d, want %d",
					tst.target, tst.position, tst.speed, got, tst.want)
			}

			got = carFleet2(tst.target, slices.Clone(tst.position), tst.speed)
			if got != tst.want {
				t.Errorf("carFleet(%d, %v, %v) = %d, want %d",
					tst.target, tst.position, tst.speed, got, tst.want)
			}
		})
	}
}
