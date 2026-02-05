package leetcode0042

import "testing"

func TestTrapExamples(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			"example_1",
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			"example_2",
			[]int{4, 2, 0, 3, 2, 5},
			9,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := trap(tst.height)
			if got != tst.want {
				t.Errorf("trap(%v) = %v, want %v", tst.height, got, tst.want)
			}
		})
	}
}

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		// Edge cases: minimum length
		{"single_element", []int{5}, 0},
		{"two_elements", []int{1, 2}, 0},
		{"three_elements_no_trap", []int{1, 2, 3}, 0},
		{"three_elements_with_trap", []int{2, 0, 2}, 2},

		// All same height
		{"all_zeros", []int{0, 0, 0, 0}, 0},
		{"all_same_height", []int{3, 3, 3, 3}, 0},
		{"all_ones", []int{1, 1, 1, 1, 1}, 0},

		// Monotonic sequences (no water)
		{"increasing", []int{1, 2, 3, 4, 5}, 0},
		{"decreasing", []int{5, 4, 3, 2, 1}, 0},
		{"strictly_increasing", []int{0, 1, 2, 3, 4, 5, 6}, 0},
		{"strictly_decreasing", []int{6, 5, 4, 3, 2, 1, 0}, 0},

		// Simple valley patterns
		{"simple_valley", []int{3, 0, 3}, 3},
		{"deep_valley", []int{5, 0, 5}, 5},
		{"asymmetric_valley_left_higher", []int{5, 0, 3}, 3},
		{"asymmetric_valley_right_higher", []int{3, 0, 5}, 3},
		{"shallow_valley", []int{2, 1, 2}, 1},

		// Multiple valleys
		{"two_valleys", []int{3, 0, 3, 0, 3}, 6},
		{"three_valleys", []int{2, 0, 2, 0, 2, 0, 2}, 6},
		{"adjacent_valleys", []int{3, 1, 2, 1, 3}, 5},

		// Step patterns
		{"step_up_then_down", []int{1, 2, 3, 2, 1}, 0},
		{"step_down_then_up", []int{3, 2, 1, 2, 3}, 4},
		{"stairs_up", []int{1, 1, 2, 2, 3, 3}, 0},
		{"stairs_down", []int{3, 3, 2, 2, 1, 1}, 0},

		// Peak and valley combinations
		{"single_peak", []int{0, 1, 2, 3, 2, 1, 0}, 0},
		{"two_peaks_one_valley", []int{3, 1, 3, 1, 3}, 4},
		{"mountain_range", []int{1, 3, 1, 3, 1, 3, 1}, 4},

		// Complex patterns
		{"zigzag", []int{1, 3, 1, 3, 1, 3, 1, 3, 1}, 6},
		{"wave_pattern", []int{2, 1, 2, 1, 2, 1, 2}, 3},
		{"complex_terrain", []int{4, 2, 3, 1, 2, 1, 3, 2, 4}, 14},

		// Boundary conditions with zeros
		{"zeros_at_edges", []int{0, 3, 0, 3, 0}, 3},
		{"zeros_in_middle", []int{3, 0, 0, 0, 3}, 9},
		{"alternating_with_zero", []int{2, 0, 2, 0, 2, 0, 2}, 6},

		// Single wall scenarios
		{"single_high_wall_left", []int{10, 1, 1, 1, 2}, 3},
		{"single_high_wall_right", []int{2, 1, 1, 1, 10}, 3},
		{"single_high_wall_middle", []int{1, 1, 10, 1, 1}, 0},

		// Wide valleys
		{"wide_shallow_valley", []int{2, 1, 1, 1, 1, 1, 2}, 5},
		{"wide_deep_valley", []int{5, 0, 0, 0, 0, 0, 5}, 25},

		// Narrow valleys
		{"many_narrow_valleys", []int{3, 2, 3, 2, 3, 2, 3, 2, 3}, 4},

		// Edge cases with max constraint values
		{"large_height_difference", []int{100000, 0, 100000}, 100000},
		{"very_tall_walls", []int{99999, 1, 99999}, 99998},

		// Irregular patterns
		{"irregular_1", []int{5, 2, 1, 2, 1, 5}, 14},
		{"irregular_2", []int{1, 0, 2, 1, 0, 1, 3}, 5},
		{"irregular_3", []int{2, 6, 3, 8, 2, 7, 2, 5, 1}, 11},

		// Plateau patterns
		{"plateau_middle", []int{1, 3, 3, 3, 1}, 0},
		{"plateau_with_valley", []int{3, 3, 1, 3, 3}, 2},
		{"double_plateau", []int{2, 2, 1, 1, 2, 2}, 2},

		// Descending then ascending
		{"v_shape", []int{4, 3, 2, 1, 2, 3, 4}, 9},
		{"u_shape", []int{4, 1, 1, 1, 4}, 9},
		{"w_shape", []int{3, 0, 2, 0, 3}, 7},

		// Longer sequences
		{"long_flat", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{"long_valley", []int{5, 1, 1, 1, 1, 1, 1, 1, 1, 5}, 32},
		{"long_zigzag", []int{3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3}, 10},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := trap(tst.height)
			if got != tst.want {
				t.Errorf("trap(%v) = %v, want %v", tst.height, got, tst.want)
			}
		})
	}
}
