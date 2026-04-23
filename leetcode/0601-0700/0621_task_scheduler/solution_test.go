// Package leetcode0621 solves LeetCode 621. Task Scheduler
package leetcode0621

import "testing"

func TestLeastIntervalExamples(t *testing.T) {
	tests := []struct {
		name  string
		tasks []byte
		n     int
		want  int
	}{
		{
			"example_1",
			[]byte{'A', 'A', 'A', 'B', 'B', 'B'},
			2,
			8,
		},
		{
			"example_2",
			[]byte{'A', 'C', 'A', 'B', 'D', 'B'},
			1,
			6,
		},
		{
			"example_3",
			[]byte{'A', 'A', 'A', 'B', 'B', 'B'},
			3,
			10,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := leastInterval(tst.tasks, tst.n)
			if got != tst.want {
				t.Errorf("leastInterval(%v, %d) = %d, want %d", tst.tasks, tst.n, got, tst.want)
			}
			got = leastInterval2(tst.tasks, tst.n)
			if got != tst.want {
				t.Errorf("leastInterval2(%v, %d) = %d, want %d", tst.tasks, tst.n, got, tst.want)
			}
		})
	}
}

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		name  string
		tasks []byte
		n     int
		want  int
	}{
		{"single_task", []byte{'A'}, 0, 1},
		{"single_task_with_cooldown", []byte{'A'}, 5, 1},
		{"all_same_tasks", []byte{'A', 'A', 'A', 'A', 'A'}, 0, 5},
		{"all_same_tasks_with_cooldown", []byte{'A', 'A', 'A', 'A', 'A'}, 2, 13},
		{"all_unique_tasks", []byte{'A', 'B', 'C', 'D', 'E'}, 100, 5},
		{"two_tasks_alternating", []byte{'A', 'B', 'A', 'B', 'A', 'B'}, 0, 6},
		{"two_tasks_with_gap", []byte{'A', 'A', 'B', 'B'}, 2, 5},
		{"three_most_frequent", []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C'}, 1, 9},
		{"three_most_frequent_with_larger_cooldown", []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C'}, 2, 9},
		{"n_zero_always_tasks_length", []byte{'X', 'Y', 'X', 'Z', 'W', 'X', 'Y', 'Z'}, 0, 8},
		{"one_dominant_task", []byte{'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E'}, 2, 10},
		{"dominant_task_with_multiple_others", []byte{'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2, 9},
		{"cooldown_larger_than_task_types", []byte{'A', 'A', 'B', 'B', 'C', 'C'}, 4, 8},
		{"exact_fill_no_idle", []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'D', 'D'}, 2, 10},
		{"all_26_letters_once", []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}, 10, 26},
		{"many_of_one_few_of_others", []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2, 16},
		{"max_length_edge_case_single_type", func() []byte {
			b := make([]byte, 10000)
			for i := range b {
				b[i] = 'X'
			}
			return b
		}(), 100, 10000 + (10000-1)*100},
		{"max_length_many_types", func() []byte {
			b := make([]byte, 10000)
			for i := range b {
				b[i] = 'A' + byte(i%26)
			}
			return b
		}(), 0, 10000},
		{"equal_frequency_all_tasks", []byte{'A', 'B', 'C', 'D', 'E', 'F', 'A', 'B', 'C', 'D', 'E', 'F', 'A', 'B', 'C', 'D', 'E', 'F'}, 2, 18},
		{"sparse_frequencies", []byte{'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N'}, 3, 16},
		{"large_cooldown_small_tasks", []byte{'A', 'B', 'A', 'B'}, 100, 103},
		{"single_task_multiple_copies", []byte{'Z', 'Z', 'Z', 'Z', 'Z', 'Z', 'Z', 'Z', 'Z', 'Z'}, 5, 55},
		{"interleaved_perfectly", []byte{'A', 'B', 'C', 'A', 'B', 'C', 'A', 'B', 'C'}, 2, 9},
		{"mostly_one_type_with_sprinkles", []byte{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'B', 'C', 'D', 'E'}, 2, 19},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			got := leastInterval(tst.tasks, tst.n)
			if got != tst.want {
				t.Errorf("leastInterval(tasks length=%d, n=%d) = %d, want %d", len(tst.tasks), tst.n, got, tst.want)
			}
			got = leastInterval2(tst.tasks, tst.n)
			if got != tst.want {
				t.Errorf("leastInterval2(tasks length=%d, n=%d) = %d, want %d", len(tst.tasks), tst.n, got, tst.want)
			}
		})
	}
}
