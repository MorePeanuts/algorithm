package leetcode0121

import "testing"

func TestMaxProfitExamples(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			"example_1",
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			"example_2",
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := maxProfit(tst.prices)
			if got != tst.want {
				t.Errorf("maxProfit(%v) = %v, want %v", tst.prices, got, tst.want)
			}
		})
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		// Single element - no transaction possible
		{"single_element", []int{5}, 0},

		// Two elements
		{"two_elements_profit", []int{1, 5}, 4},
		{"two_elements_no_profit", []int{5, 1}, 0},
		{"two_elements_equal", []int{3, 3}, 0},

		// All same prices
		{"all_same_small", []int{5, 5, 5}, 0},
		{"all_same_large", []int{100, 100, 100, 100, 100}, 0},

		// Strictly increasing
		{"strictly_increasing_small", []int{1, 2, 3, 4, 5}, 4},
		{"strictly_increasing_large", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9},

		// Strictly decreasing
		{"strictly_decreasing_small", []int{5, 4, 3, 2, 1}, 0},
		{"strictly_decreasing_large", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 0},

		// Min at beginning, max at end
		{"min_first_max_last", []int{1, 3, 2, 5, 4, 10}, 9},

		// Min at end (no profit)
		{"min_at_end", []int{5, 4, 3, 2, 1, 0}, 0},

		// Max at beginning (no profit)
		{"max_at_beginning", []int{10, 1, 2, 3, 4}, 3},

		// Valley then peak
		{"valley_peak", []int{5, 1, 6}, 5},
		{"multiple_valleys", []int{5, 2, 6, 1, 8}, 7},

		// Peak then valley
		{"peak_valley_peak", []int{1, 5, 2, 8}, 7},

		// Local maxima and minima
		{"local_maxima_minima", []int{3, 2, 6, 5, 0, 3}, 4},

		// Best buy is not at minimum index
		{"best_not_at_global_min", []int{2, 1, 4, 0, 1}, 3},

		// Multiple opportunities - first is best
		{"first_opportunity_best", []int{1, 10, 2, 3, 4}, 9},

		// Multiple opportunities - last is best
		{"last_opportunity_best", []int{5, 4, 1, 10}, 9},

		// Boundary values - price = 0
		{"zero_price_start", []int{0, 1, 2, 3}, 3},
		{"zero_price_end", []int{3, 2, 1, 0}, 0},
		{"all_zeros", []int{0, 0, 0, 0}, 0},

		// Boundary values - max price 10^4
		{"max_price", []int{10000, 0, 10000}, 10000},
		{"max_profit_possible", []int{0, 10000}, 10000},

		// Alternating prices
		{"alternating_up_down", []int{1, 5, 1, 5, 1, 5}, 4},
		{"alternating_down_up", []int{5, 1, 5, 1, 5, 1}, 4},

		// Single dip pattern
		{"single_dip", []int{5, 4, 3, 2, 3, 4, 5}, 3},

		// Single peak pattern
		{"single_peak", []int{1, 2, 3, 4, 3, 2, 1}, 3},

		// Large profit at end
		{"large_profit_end", []int{100, 50, 25, 10, 5, 1, 1000}, 999},

		// Small fluctuations
		{"small_fluctuations", []int{10, 11, 10, 11, 10, 11}, 1},

		// Medium size array
		{"medium_array_increasing", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 14},
		{"medium_array_decreasing", []int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 0},

		// Repeated values with profit
		{"repeated_with_profit", []int{1, 1, 1, 5, 5, 5}, 4},

		// Profit of 1
		{"profit_of_one", []int{5, 4, 3, 2, 1, 2}, 1},

		// W shape
		{"w_shape", []int{5, 1, 5, 1, 5}, 4},

		// M shape
		{"m_shape", []int{1, 5, 1, 5, 1}, 4},

		// Long plateau then drop then rise
		{"plateau_drop_rise", []int{5, 5, 5, 1, 1, 1, 10, 10, 10}, 9},

		// Drop at very end
		{"drop_at_end", []int{1, 2, 3, 4, 5, 1}, 4},

		// Rise at very end
		{"rise_at_end", []int{5, 4, 3, 2, 1, 10}, 9},

		// Best profit in middle
		{"best_in_middle", []int{5, 1, 10, 2, 3}, 9},

		// Multiple same profits
		{"multiple_same_profits", []int{1, 5, 1, 5, 1, 5}, 4},

		// Three elements
		{"three_elements_asc", []int{1, 2, 3}, 2},
		{"three_elements_desc", []int{3, 2, 1}, 0},
		{"three_elements_valley", []int{3, 1, 2}, 1},
		{"three_elements_peak", []int{1, 3, 2}, 2},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := maxProfit(tst.prices)
			if got != tst.want {
				t.Errorf("maxProfit(%v) = %v, want %v", tst.prices, got, tst.want)
			}
		})
	}
}
