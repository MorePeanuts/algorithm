package leetcode0074

import "testing"

func TestSearchMatrixExamples(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			"example_1",
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			3,
			true,
		},
		{
			"example_2",
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			13,
			false,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := searchMatrix(tst.matrix, tst.target)
			if got != tst.want {
				t.Errorf("searchMatrix(%v, %v) = %v, want %v", tst.matrix, tst.target, got, tst.want)
			}
		})
	}
}

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		// Single element matrix
		{"single_element_found", [][]int{{5}}, 5, true},
		{"single_element_not_found", [][]int{{5}}, 3, false},

		// Single row
		{"single_row_found_first", [][]int{{1, 3, 5, 7}}, 1, true},
		{"single_row_found_last", [][]int{{1, 3, 5, 7}}, 7, true},
		{"single_row_found_middle", [][]int{{1, 3, 5, 7}}, 3, true},
		{"single_row_not_found_less", [][]int{{1, 3, 5, 7}}, 0, false},
		{"single_row_not_found_greater", [][]int{{1, 3, 5, 7}}, 8, false},
		{"single_row_not_found_between", [][]int{{1, 3, 5, 7}}, 4, false},

		// Single column
		{"single_column_found_first", [][]int{{1}, {3}, {5}, {7}}, 1, true},
		{"single_column_found_last", [][]int{{1}, {3}, {5}, {7}}, 7, true},
		{"single_column_found_middle", [][]int{{1}, {3}, {5}, {7}}, 3, true},
		{"single_column_not_found_less", [][]int{{1}, {3}, {5}, {7}}, 0, false},
		{"single_column_not_found_greater", [][]int{{1}, {3}, {5}, {7}}, 8, false},
		{"single_column_not_found_between", [][]int{{1}, {3}, {5}, {7}}, 4, false},

		// Target at boundaries
		{"target_first_row_first_col", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1, true},
		{"target_first_row_last_col", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 3, true},
		{"target_last_row_first_col", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 7, true},
		{"target_last_row_last_col", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9, true},

		// Target between rows
		{"target_between_rows_less_all", [][]int{{1, 3, 5}, {7, 9, 11}, {13, 15, 17}}, 0, false},
		{"target_between_rows_greater_all", [][]int{{1, 3, 5}, {7, 9, 11}, {13, 15, 17}}, 18, false},
		{"target_between_rows_1_and_2", [][]int{{1, 3, 5}, {7, 9, 11}, {13, 15, 17}}, 6, false},
		{"target_between_rows_2_and_3", [][]int{{1, 3, 5}, {7, 9, 11}, {13, 15, 17}}, 12, false},

		// Target within row but not present
		{"target_in_row_range_not_present", [][]int{{1, 4, 7}, {10, 13, 16}, {19, 22, 25}}, 5, false},

		// Negative numbers
		{"negative_target_found", [][]int{{-10, -5, 0}, {5, 10, 15}}, -5, true},
		{"negative_target_not_found", [][]int{{-10, -5, 0}, {5, 10, 15}}, -3, false},
		{"all_negative_found", [][]int{{-20, -15, -10}, {-5, -3, -1}}, -15, true},
		{"all_negative_not_found", [][]int{{-20, -15, -10}, {-5, -3, -1}}, -7, false},

		// Maximum size matrix (100x100) - we'll do 10x10 for practicality
		{
			"large_matrix_found",
			func() [][]int {
				m := make([][]int, 10)
				for i := 0; i < 10; i++ {
					m[i] = make([]int, 10)
					for j := 0; j < 10; j++ {
						m[i][j] = i*10 + j
					}
				}
				return m
			}(),
			55,
			true,
		},
		{
			"large_matrix_not_found",
			func() [][]int {
				m := make([][]int, 10)
				for i := 0; i < 10; i++ {
					m[i] = make([]int, 10)
					for j := 0; j < 10; j++ {
						m[i][j] = i*10 + j
					}
				}
				return m
			}(),
			100,
			false,
		},

		// Edge values from constraints
		{"min_value_found", [][]int{{-10000, 0}, {10000, 20000}}, -10000, true},
		{"max_value_found", [][]int{{-10000, 0}, {10000, 20000}}, 20000, true},
		{"value_out_of_range_negative", [][]int{{-10000, 0}, {10000, 20000}}, -10001, false},
		{"value_out_of_range_positive", [][]int{{-10000, 0}, {10000, 20000}}, 20001, false},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := searchMatrix(tst.matrix, tst.target)
			if got != tst.want {
				t.Errorf("searchMatrix(%v, %v) = %v, want %v", tst.matrix, tst.target, got, tst.want)
			}
		})
	}
}
