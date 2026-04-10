package leetcode0143

import (
	"testing"

	"github.com/MorePeanuts/algorithm/leetcode/utils"
)

func equalIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestReorderListExamples(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{
			"example_1",
			[]int{1, 2, 3, 4},
			[]int{1, 4, 2, 3},
		},
		{
			"example_2",
			[]int{1, 2, 3, 4, 5},
			[]int{1, 5, 2, 4, 3},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head1 := utils.SliceToList(tst.head)
			reorderList(head1)
			got1 := utils.ListToSlice(head1)
			if !equalIntSlice(got1, tst.want) {
				t.Errorf("reorderList(%v) = %v, want %v", tst.head, got1, tst.want)
			}

			head2 := utils.SliceToList(tst.head)
			reorderList2(head2)
			got2 := utils.ListToSlice(head2)
			if !equalIntSlice(got2, tst.want) {
				t.Errorf("reorderList2(%v) = %v, want %v", tst.head, got2, tst.want)
			}

			head3 := utils.SliceToList(tst.head)
			reorderList3(head3)
			got3 := utils.ListToSlice(head3)
			if !equalIntSlice(got3, tst.want) {
				t.Errorf("reorderList3(%v) = %v, want %v", tst.head, got3, tst.want)
			}
		})
	}
}

func TestReorderList(t *testing.T) {
	tests := []struct {
		name string
		head []int
		want []int
	}{
		{"single_node", []int{1}, []int{1}},
		{"two_nodes", []int{1, 2}, []int{1, 2}},
		{"three_nodes", []int{1, 2, 3}, []int{1, 3, 2}},
		{"four_nodes_duplicate_values", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"five_nodes_duplicate_values", []int{2, 2, 3, 3, 3}, []int{2, 3, 2, 3, 3}},
		{"six_nodes", []int{1, 2, 3, 4, 5, 6}, []int{1, 6, 2, 5, 3, 4}},
		{"seven_nodes", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 7, 2, 6, 3, 5, 4}},
		{"eight_nodes", []int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1, 8, 2, 7, 3, 6, 4, 5}},
		{"descending_order", []int{5, 4, 3, 2, 1}, []int{5, 1, 4, 2, 3}},
		{"ascending_order", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}},
		{"all_same_value", []int{5, 5, 5, 5, 5, 5}, []int{5, 5, 5, 5, 5, 5}},
		{"alternating_values", []int{1, 2, 1, 2, 1, 2, 1}, []int{1, 1, 2, 2, 1, 1, 2}},
		{"large_values", []int{1000, 999, 998, 997, 996}, []int{1000, 996, 999, 997, 998}},
		{"small_values", []int{1, 2, 1, 1, 1, 2}, []int{1, 2, 2, 1, 1, 1}},
		{"nine_nodes", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 9, 2, 8, 3, 7, 4, 6, 5}},
		{"ten_nodes_duplicates", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}, []int{3, 3, 1, 5, 4, 6, 1, 2, 5, 9}},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			head1 := utils.SliceToList(tst.head)
			reorderList(head1)
			got1 := utils.ListToSlice(head1)
			if !equalIntSlice(got1, tst.want) {
				t.Errorf("reorderList(%v) = %v, want %v", tst.head, got1, tst.want)
			}

			head2 := utils.SliceToList(tst.head)
			reorderList2(head2)
			got2 := utils.ListToSlice(head2)
			if !equalIntSlice(got2, tst.want) {
				t.Errorf("reorderList2(%v) = %v, want %v", tst.head, got2, tst.want)
			}

			head3 := utils.SliceToList(tst.head)
			reorderList3(head3)
			got3 := utils.ListToSlice(head3)
			if !equalIntSlice(got3, tst.want) {
				t.Errorf("reorderList3(%v) = %v, want %v", tst.head, got3, tst.want)
			}
		})
	}
}
