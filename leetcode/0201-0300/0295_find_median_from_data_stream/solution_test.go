package leetcode0295

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

// op represents a single operation on MedianFinder.
type op struct {
	name string   // "MedianFinder", "addNum", or "findMedian"
	args []int    // args for addNum; empty for findMedian/Constructor
	want float64  // expected result for findMedian; 0 for others
}

// runOps executes a sequence of operations and verifies findMedian results.
func runOps(t *testing.T, ops []op) {
	t.Helper()
	var mf MedianFinder
	for _, o := range ops {
		switch o.name {
		case "MedianFinder":
			mf = Constructor()
		case "addNum":
			mf.AddNum(o.args[0])
		case "findMedian":
			got := mf.FindMedian()
			if !utils.Float64AlmostEqual(got, o.want) {
				t.Errorf("FindMedian() = %v, want %v", got, o.want)
			}
		}
	}
}

func TestMedianFinderExamples(t *testing.T) {
	tests := []struct {
		name string
		ops  []op
	}{
		{
			"example_1",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{2}, 0},
				{"findMedian", nil, 1.5},
				{"addNum", []int{3}, 0},
				{"findMedian", nil, 2.0},
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			runOps(t, tst.ops)
		})
	}
}

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name string
		ops  []op
	}{
		{
			"single_element",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{5}, 0},
				{"findMedian", nil, 5.0},
			},
		},
		{
			"two_elements_even_median",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{2}, 0},
				{"findMedian", nil, 1.5},
			},
		},
		{
			"three_elements_odd_median",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{2}, 0},
				{"addNum", []int{3}, 0},
				{"findMedian", nil, 2.0},
			},
		},
		{
			"descending_order",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{5}, 0},
				{"addNum", []int{4}, 0},
				{"addNum", []int{3}, 0},
				{"addNum", []int{2}, 0},
				{"addNum", []int{1}, 0},
				{"findMedian", nil, 3.0},
			},
		},
		{
			"ascending_order",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{2}, 0},
				{"addNum", []int{3}, 0},
				{"addNum", []int{4}, 0},
				{"addNum", []int{5}, 0},
				{"findMedian", nil, 3.0},
			},
		},
		{
			"all_same_elements",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{7}, 0},
				{"addNum", []int{7}, 0},
				{"addNum", []int{7}, 0},
				{"findMedian", nil, 7.0},
			},
		},
		{
			"all_same_elements_even_count",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{3}, 0},
				{"addNum", []int{3}, 0},
				{"findMedian", nil, 3.0},
			},
		},
		{
			"negative_numbers",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{-1}, 0},
				{"addNum", []int{-2}, 0},
				{"addNum", []int{-3}, 0},
				{"findMedian", nil, -2.0},
			},
		},
		{
			"mixed_negative_and_positive",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{-5}, 0},
				{"addNum", []int{5}, 0},
				{"findMedian", nil, 0.0},
			},
		},
		{
			"min_constraint_value",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{-100000}, 0},
				{"findMedian", nil, -100000.0},
			},
		},
		{
			"max_constraint_value",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{100000}, 0},
				{"findMedian", nil, 100000.0},
			},
		},
		{
			"min_and_max_values",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{-100000}, 0},
				{"addNum", []int{100000}, 0},
				{"findMedian", nil, 0.0},
			},
		},
		{
			"zero_value",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{0}, 0},
				{"addNum", []int{0}, 0},
				{"findMedian", nil, 0.0},
			},
		},
		{
			"duplicate_median_pair",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{2}, 0},
				{"addNum", []int{2}, 0},
				{"addNum", []int{3}, 0},
				{"addNum", []int{3}, 0},
				{"findMedian", nil, 2.5},
			},
		},
		{
			"interleaved_find_after_each_add",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{6}, 0},
				{"findMedian", nil, 6.0},
				{"addNum", []int{10}, 0},
				{"findMedian", nil, 8.0},
				{"addNum", []int{2}, 0},
				{"findMedian", nil, 6.0},
				{"addNum", []int{6}, 0},
				{"findMedian", nil, 6.0},
				{"addNum", []int{5}, 0},
				{"findMedian", nil, 6.0},
				{"addNum", []int{0}, 0},
				{"findMedian", nil, 5.5},
			},
		},
		{
			"large_odd_count",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{3}, 0},
				{"addNum", []int{5}, 0},
				{"addNum", []int{7}, 0},
				{"addNum", []int{9}, 0},
				{"addNum", []int{11}, 0},
				{"addNum", []int{13}, 0},
				{"findMedian", nil, 7.0},
			},
		},
		{
			"large_even_count",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{3}, 0},
				{"addNum", []int{5}, 0},
				{"addNum", []int{7}, 0},
				{"addNum", []int{9}, 0},
				{"addNum", []int{11}, 0},
				{"findMedian", nil, 6.0},
			},
		},
		{
			"two_elements_one_negative",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{-3}, 0},
				{"addNum", []int{4}, 0},
				{"findMedian", nil, 0.5},
			},
		},
		{
			"median_after_removing_balance",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{2}, 0},
				{"addNum", []int{3}, 0},
				{"addNum", []int{4}, 0},
				{"addNum", []int{5}, 0},
				{"addNum", []int{6}, 0},
				{"addNum", []int{7}, 0},
				{"addNum", []int{8}, 0},
				{"findMedian", nil, 4.5},
			},
		},
		{
			"repeated_small_large_alternating",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{100}, 0},
				{"addNum", []int{2}, 0},
				{"addNum", []int{99}, 0},
				{"findMedian", nil, 50.5},
			},
		},
		{
			"single_negative",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{-1}, 0},
				{"findMedian", nil, -1.0},
			},
		},
		{
			"two_same_negatives",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{-5}, 0},
				{"addNum", []int{-5}, 0},
				{"findMedian", nil, -5.0},
			},
		},
		{
			"median_shifts_with_new_elements",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{10}, 0},
				{"findMedian", nil, 10.0},
				{"addNum", []int{20}, 0},
				{"findMedian", nil, 15.0},
				{"addNum", []int{30}, 0},
				{"findMedian", nil, 20.0},
				{"addNum", []int{40}, 0},
				{"findMedian", nil, 25.0},
			},
		},
		{
			"many_duplicates",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{1}, 0},
				{"findMedian", nil, 1.0},
			},
		},
		{
			"boundary_values_interleaved",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{-100000}, 0},
				{"addNum", []int{100000}, 0},
				{"addNum", []int{-100000}, 0},
				{"addNum", []int{100000}, 0},
				{"findMedian", nil, 0.0},
			},
		},
		{
			"insertion_into_middle",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{5}, 0},
				{"addNum", []int{3}, 0},
				{"findMedian", nil, 3.0},
			},
		},
		{
			"four_elements_even_median_fractional",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{1}, 0},
				{"addNum", []int{2}, 0},
				{"addNum", []int{3}, 0},
				{"addNum", []int{4}, 0},
				{"findMedian", nil, 2.5},
			},
		},
		{
			"reverse_sorted_even",
			[]op{
				{"MedianFinder", nil, 0},
				{"addNum", []int{4}, 0},
				{"addNum", []int{3}, 0},
				{"addNum", []int{2}, 0},
				{"addNum", []int{1}, 0},
				{"findMedian", nil, 2.5},
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			runOps(t, tst.ops)
		})
	}
}
