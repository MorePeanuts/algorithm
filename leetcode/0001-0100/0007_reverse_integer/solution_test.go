package leetcode0007

import (
	"math"
	"testing"
)

func TestReverseExamples(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"example_1", 123, 321},
		{"example_2", -123, -321},
		{"example_3", 120, 21},
		{"example_4", 0, 0},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := reverse(tst.x)
			if got != tst.want {
				t.Errorf("reverse(%d) = %d, want %d", tst.x, got, tst.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		// 基本正数测试
		{"single_digit_1", 1, 1},
		{"single_digit_9", 9, 9},
		{"two_digits", 12, 21},
		{"three_digits", 321, 123},
		{"palindrome", 121, 121},
		{"large_positive", 1234567, 7654321},

		// 基本负数测试
		{"negative_single_digit", -1, -1},
		{"negative_two_digits", -12, -21},
		{"negative_three_digits", -321, -123},
		{"negative_palindrome", -121, -121},
		{"negative_large", -1234567, -7654321},

		// 零相关测试
		{"zero", 0, 0},
		{"trailing_zeros", 100, 1},
		{"trailing_zeros_2", 1000, 1},
		{"trailing_zeros_3", 10200, 201},
		{"negative_trailing_zeros", -100, -1},
		{"negative_trailing_zeros_2", -1200, -21},

		// 边界值测试 (32-bit signed integer)
		{"max_int32", math.MaxInt32, 0},                  // 2147483647 -> 7463847412 (overflow)
		{"min_int32", math.MinInt32, 0},                  // -2147483648 -> overflow
		{"near_max_no_overflow", 1463847412, 2147483641}, // reverse is within range
		{"near_max_overflow", 1563847412, 0},             // reverse would overflow
		{"near_min_no_overflow", -1463847412, -2147483641},
		{"near_min_overflow", -1563847412, 0},

		// 特殊模式测试
		{"all_ones", 111111111, 111111111},
		{"all_nines", 999999999, 999999999},
		{"negative_all_ones", -111111111, -111111111},
		{"alternating_digits", 12121212, 21212121},

		// 溢出边界精确测试
		{"overflow_boundary_1", 2147483647, 0},  // max int32
		{"overflow_boundary_2", -2147483648, 0}, // min int32
		{"overflow_boundary_3", 1000000003, 0},  // may overflow
		{"overflow_check_1", 1534236469, 0},     // 9646324351 > max int32
		{"overflow_check_2", -1534236469, 0},

		// 更多边缘情况
		{"two_digits_trailing_zero", 10, 1},
		{"three_digits_trailing_zero", 200, 2},
		{"mixed_zeros", 102030, 30201},
		{"negative_mixed_zeros", -102030, -30201},
		{"large_with_zeros", 1020304050, 504030201},
		{"negative_large_with_zeros", -1020304050, -504030201},

		// 回文数测试
		{"palindrome_2", 1221, 1221},
		{"palindrome_3", 12321, 12321},
		{"palindrome_4", 123321, 123321},
		{"negative_palindrome_2", -1221, -1221},

		// 小数值测试
		{"small_2", 2, 2},
		{"small_3", 3, 3},
		{"small_10", 10, 1},
		{"small_11", 11, 11},
		{"negative_small", -10, -1},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := reverse(tst.x)
			if got != tst.want {
				t.Errorf("reverse(%d) = %d, want %d", tst.x, got, tst.want)
			}
		})
	}
}
