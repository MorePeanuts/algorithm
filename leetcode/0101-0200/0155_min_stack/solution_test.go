package leetcode0155

import (
	"math"
	"testing"
)

func TestMinStackExamples(t *testing.T) {
	// Example 1:
	// ["MinStack","push","push","push","getMin","pop","top","getMin"]
	// [[],[-2],[0],[-3],[],[],[],[]]
	// Output: [null,null,null,null,-3,null,0,-2]
	t.Run("example_1", func(t *testing.T) {
		minStack := Constructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)

		if got := minStack.GetMin(); got != -3 {
			t.Errorf("GetMin() = %v, want -3", got)
		}

		minStack.Pop()

		if got := minStack.Top(); got != 0 {
			t.Errorf("Top() = %v, want 0", got)
		}

		if got := minStack.GetMin(); got != -2 {
			t.Errorf("GetMin() = %v, want -2", got)
		}
	})

	t.Run("example_2", func(t *testing.T) {
		minStack := Constructor2()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)

		if got := minStack.GetMin(); got != -3 {
			t.Errorf("GetMin() = %v, want -3", got)
		}

		minStack.Pop()

		if got := minStack.Top(); got != 0 {
			t.Errorf("Top() = %v, want 0", got)
		}

		if got := minStack.GetMin(); got != -2 {
			t.Errorf("GetMin() = %v, want -2", got)
		}
	})
}

func TestMinStack(t *testing.T) {
	tests := []struct {
		name       string
		operations []string
		values     []int
		expected   []int // -1 means no return value (for push/pop)
	}{
		{
			name:       "single_push_top",
			operations: []string{"push", "top"},
			values:     []int{5, 0},
			expected:   []int{-1, 5},
		},
		{
			name:       "single_push_getMin",
			operations: []string{"push", "getMin"},
			values:     []int{10, 0},
			expected:   []int{-1, 10},
		},
		{
			name:       "push_pop_push_top",
			operations: []string{"push", "pop", "push", "top"},
			values:     []int{1, 0, 2, 0},
			expected:   []int{-1, -1, -1, 2},
		},
		{
			name:       "multiple_same_values",
			operations: []string{"push", "push", "push", "getMin", "pop", "getMin"},
			values:     []int{5, 5, 5, 0, 0, 0},
			expected:   []int{-1, -1, -1, 5, -1, 5},
		},
		{
			name:       "decreasing_sequence",
			operations: []string{"push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"},
			values:     []int{3, 2, 1, 0, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, 1, -1, 2, -1, 3},
		},
		{
			name:       "increasing_sequence",
			operations: []string{"push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"},
			values:     []int{1, 2, 3, 0, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, 1, -1, 1, -1, 1},
		},
		{
			name:       "negative_values",
			operations: []string{"push", "push", "push", "getMin", "top"},
			values:     []int{-1, -2, -3, 0, 0},
			expected:   []int{-1, -1, -1, -3, -3},
		},
		{
			name:       "mixed_positive_negative",
			operations: []string{"push", "push", "push", "push", "getMin", "top"},
			values:     []int{-5, 10, -3, 7, 0, 0},
			expected:   []int{-1, -1, -1, -1, -5, 7},
		},
		{
			name:       "min_changes_after_pop",
			operations: []string{"push", "push", "getMin", "pop", "getMin"},
			values:     []int{10, 5, 0, 0, 0},
			expected:   []int{-1, -1, 5, -1, 10},
		},
		{
			name:       "min_stays_same_after_pop",
			operations: []string{"push", "push", "getMin", "pop", "getMin"},
			values:     []int{5, 10, 0, 0, 0},
			expected:   []int{-1, -1, 5, -1, 5},
		},
		{
			name:       "zero_value",
			operations: []string{"push", "push", "push", "getMin", "top"},
			values:     []int{0, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, 0, 0},
		},
		{
			name:       "single_element_operations",
			operations: []string{"push", "top", "getMin", "pop"},
			values:     []int{42, 0, 0, 0},
			expected:   []int{-1, 42, 42, -1},
		},
		{
			name:       "push_same_min_twice",
			operations: []string{"push", "push", "getMin", "pop", "getMin"},
			values:     []int{1, 1, 0, 0, 0},
			expected:   []int{-1, -1, 1, -1, 1},
		},
		{
			name:       "large_positive_value",
			operations: []string{"push", "top", "getMin"},
			values:     []int{math.MaxInt32, 0, 0},
			expected:   []int{-1, math.MaxInt32, math.MaxInt32},
		},
		{
			name:       "large_negative_value",
			operations: []string{"push", "top", "getMin"},
			values:     []int{math.MinInt32, 0, 0},
			expected:   []int{-1, math.MinInt32, math.MinInt32},
		},
		{
			name:       "min_max_values",
			operations: []string{"push", "push", "getMin", "top"},
			values:     []int{math.MaxInt32, math.MinInt32, 0, 0},
			expected:   []int{-1, -1, math.MinInt32, math.MinInt32},
		},
		{
			name:       "alternating_push_pop",
			operations: []string{"push", "push", "pop", "push", "pop", "push", "getMin", "top"},
			values:     []int{3, 1, 0, 2, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, -1, -1, -1, 0, 0},
		},
		{
			name:       "pop_until_one",
			operations: []string{"push", "push", "push", "push", "pop", "pop", "pop", "top", "getMin"},
			values:     []int{4, 3, 2, 1, 0, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, -1, -1, -1, -1, 4, 4},
		},
		{
			name:       "repeated_getMin",
			operations: []string{"push", "push", "getMin", "getMin", "getMin"},
			values:     []int{5, 3, 0, 0, 0},
			expected:   []int{-1, -1, 3, 3, 3},
		},
		{
			name:       "repeated_top",
			operations: []string{"push", "push", "top", "top", "top"},
			values:     []int{1, 2, 0, 0, 0},
			expected:   []int{-1, -1, 2, 2, 2},
		},
		{
			name:       "new_min_in_middle",
			operations: []string{"push", "push", "push", "push", "push", "getMin"},
			values:     []int{5, 3, 1, 2, 4, 0},
			expected:   []int{-1, -1, -1, -1, -1, 1},
		},
		{
			name:       "v_shaped_values",
			operations: []string{"push", "push", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin"},
			values:     []int{5, 3, 1, 2, 4, 0, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, -1, -1, 1, -1, 1, -1, 1},
		},
		{
			name:       "inverted_v_shaped_values",
			operations: []string{"push", "push", "push", "push", "push", "getMin", "pop", "getMin"},
			values:     []int{1, 3, 5, 3, 1, 0, 0, 0},
			expected:   []int{-1, -1, -1, -1, -1, 1, -1, 1},
		},
		{
			name:       "zigzag_values",
			operations: []string{"push", "push", "push", "push", "getMin"},
			values:     []int{1, 10, 2, 9, 0},
			expected:   []int{-1, -1, -1, -1, 1},
		},
		{
			name:       "all_operations_sequence",
			operations: []string{"push", "push", "top", "getMin", "pop", "top", "getMin"},
			values:     []int{10, 20, 0, 0, 0, 0, 0},
			expected:   []int{-1, -1, 20, 10, -1, 10, 10},
		},
		{
			name:       "push_smaller_then_larger",
			operations: []string{"push", "push", "getMin", "push", "getMin"},
			values:     []int{10, 5, 0, 15, 0},
			expected:   []int{-1, -1, 5, -1, 5},
		},
		{
			name:       "consecutive_minimums",
			operations: []string{"push", "push", "push", "pop", "getMin", "pop", "getMin"},
			values:     []int{3, 2, 1, 0, 0, 0, 0},
			expected:   []int{-1, -1, -1, -1, 2, -1, 3},
		},
		{
			name:       "many_operations_stress",
			operations: []string{"push", "push", "push", "push", "push", "push", "push", "push", "push", "push", "getMin", "top"},
			values:     []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 0},
			expected:   []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, 1},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			minStack := Constructor()
			minStack2 := Constructor2()
			for i, op := range tst.operations {
				switch op {
				case "push":
					minStack.Push(tst.values[i])
					minStack2.Push(tst.values[i])
				case "pop":
					minStack.Pop()
					minStack2.Pop()
				case "top":
					if got := minStack.Top(); got != tst.expected[i] {
						t.Errorf("solution1: operation %d: Top() = %v, want %v", i, got, tst.expected[i])
					}
					if got := minStack2.Top(); got != tst.expected[i] {
						t.Errorf("solution2: operation %d: Top() = %v, want %v", i, got, tst.expected[i])
					}
				case "getMin":
					if got := minStack.GetMin(); got != tst.expected[i] {
						t.Errorf("solution1: operation %d: GetMin() = %v, want %v", i, got, tst.expected[i])
					}
					if got := minStack2.GetMin(); got != tst.expected[i] {
						t.Errorf("solution2: operation %d: GetMin() = %v, want %v", i, got, tst.expected[i])
					}
				}
			}
		})
	}
}
