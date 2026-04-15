// Package leetcode0981 solves LeetCode 981. Time Based Key-Value Store
package leetcode0981

import "testing"

func TestTimeMapExamples(t *testing.T) {
	tm := Constructor()

	// Example 1 operations
	tm.Set("foo", "bar", 1)
	if got := tm.Get("foo", 1); got != "bar" {
		t.Errorf("Get(\"foo\", 1) = %q, want %q", got, "bar")
	}
	if got := tm.Get("foo", 3); got != "bar" {
		t.Errorf("Get(\"foo\", 3) = %q, want %q", got, "bar")
	}

	tm.Set("foo", "bar2", 4)
	if got := tm.Get("foo", 4); got != "bar2" {
		t.Errorf("Get(\"foo\", 4) = %q, want %q", got, "bar2")
	}
	if got := tm.Get("foo", 5); got != "bar2" {
		t.Errorf("Get(\"foo\", 5) = %q, want %q", got, "bar2")
	}
}

func TestTimeMap(t *testing.T) {
	tests := []struct {
		name string
		// Sequence of operations to perform
		ops []struct {
			op       string // "set" or "get"
			key      string
			value    string // only for "set"
			timestamp int
			want     string // only for "get"
		}
	}{
		{
			"empty_key",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"get", "nonexistent", "", 1, ""},
			},
		},
		{
			"timestamp_before_first",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "key1", "val1", 5, ""},
				{"get", "key1", "", 3, ""},
			},
		},
		{
			"exact_timestamp_match",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "key1", "val1", 10, ""},
				{"get", "key1", "", 10, "val1"},
			},
		},
		{
			"timestamp_between_two_values",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "key1", "val1", 1, ""},
				{"set", "key1", "val2", 10, ""},
				{"get", "key1", "", 5, "val1"},
			},
		},
		{
			"multiple_timestamps_same_key",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "k", "v1", 1, ""},
				{"set", "k", "v2", 2, ""},
				{"set", "k", "v3", 3, ""},
				{"get", "k", "", 0, ""},
				{"get", "k", "", 1, "v1"},
				{"get", "k", "", 2, "v2"},
				{"get", "k", "", 3, "v3"},
				{"get", "k", "", 100, "v3"},
			},
		},
		{
			"multiple_keys",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "keyA", "a1", 5, ""},
				{"set", "keyB", "b1", 10, ""},
				{"set", "keyA", "a2", 15, ""},
				{"get", "keyA", "", 7, "a1"},
				{"get", "keyB", "", 7, ""},
				{"get", "keyB", "", 12, "b1"},
				{"get", "keyA", "", 20, "a2"},
			},
		},
		{
			"numeric_key_and_value",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "123", "456", 1, ""},
				{"get", "123", "", 1, "456"},
			},
		},
		{
			"max_length_key_value",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghij", "value1234567890value1234567890value1234567890value1234567890value1234567890value1234567890value1234567890value1234567890value1234567890value1234567890", 100, ""},
				{"get", "abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghij", "", 100, "value1234567890value1234567890value1234567890value1234567890value1234567890value1234567890value1234567890value1234567890value1234567890value1234567890"},
			},
		},
		{
			"large_timestamp",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "big", "time", 10000000, ""},
				{"get", "big", "", 10000000, "time"},
				{"get", "big", "", 9999999, ""},
			},
		},
		{
			"consecutive_timestamps",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "seq", "v1", 1, ""},
				{"set", "seq", "v2", 2, ""},
				{"set", "seq", "v3", 3, ""},
				{"set", "seq", "v4", 4, ""},
				{"set", "seq", "v5", 5, ""},
				{"get", "seq", "", 3, "v3"},
				{"get", "seq", "", 4, "v4"},
			},
		},
		{
			"non_consecutive_timestamps",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "gap", "a", 1, ""},
				{"set", "gap", "b", 100, ""},
				{"set", "gap", "c", 1000, ""},
				{"get", "gap", "", 50, "a"},
				{"get", "gap", "", 500, "b"},
				{"get", "gap", "", 2000, "c"},
			},
		},
		{
			"same_value_different_timestamps",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "same", "val", 1, ""},
				{"set", "same", "val", 2, ""},
				{"set", "same", "val", 3, ""},
				{"get", "same", "", 2, "val"},
			},
		},
		{
			"interleaved_operations",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "x", "x1", 1, ""},
				{"get", "x", "", 1, "x1"},
				{"set", "y", "y1", 2, ""},
				{"get", "y", "", 2, "y1"},
				{"set", "x", "x2", 3, ""},
				{"get", "x", "", 3, "x2"},
				{"get", "y", "", 3, "y1"},
			},
		},
		{
			"get_at_timestamp_exactly_equal",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "test", "first", 5, ""},
				{"set", "test", "second", 10, ""},
				{"set", "test", "third", 15, ""},
				{"get", "test", "", 5, "first"},
				{"get", "test", "", 10, "second"},
				{"get", "test", "", 15, "third"},
			},
		},
		{
			"many_timestamps_binary_search_check",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "many", "v01", 10, ""},
				{"set", "many", "v02", 20, ""},
				{"set", "many", "v03", 30, ""},
				{"set", "many", "v04", 40, ""},
				{"set", "many", "v05", 50, ""},
				{"set", "many", "v06", 60, ""},
				{"set", "many", "v07", 70, ""},
				{"set", "many", "v08", 80, ""},
				{"set", "many", "v09", 90, ""},
				{"set", "many", "v10", 100, ""},
				{"get", "many", "", 5, ""},
				{"get", "many", "", 15, "v01"},
				{"get", "many", "", 25, "v02"},
				{"get", "many", "", 35, "v03"},
				{"get", "many", "", 45, "v04"},
				{"get", "many", "", 55, "v05"},
				{"get", "many", "", 65, "v06"},
				{"get", "many", "", 75, "v07"},
				{"get", "many", "", 85, "v08"},
				{"get", "many", "", 95, "v09"},
				{"get", "many", "", 105, "v10"},
			},
		},
		{
			"single_key_single_value",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "single", "only", 100, ""},
				{"get", "single", "", 99, ""},
				{"get", "single", "", 100, "only"},
				{"get", "single", "", 101, "only"},
			},
		},
		{
			"keys_with_numbers",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "key123", "val456", 1, ""},
				{"set", "key456", "val789", 2, ""},
				{"get", "key123", "", 1, "val456"},
				{"get", "key456", "", 2, "val789"},
			},
		},
		{
			"overwrite_same_timestamp_not_allowed_but_test_strictly_increasing",
			[]struct {
				op       string
				key      string
				value    string
				timestamp int
				want     string
			}{
				{"set", "strict", "v1", 1, ""},
				{"set", "strict", "v2", 2, ""},
				{"set", "strict", "v3", 3, ""},
				{"get", "strict", "", 2, "v2"},
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			tm := Constructor()
			for _, op := range tst.ops {
				switch op.op {
				case "set":
					tm.Set(op.key, op.value, op.timestamp)
				case "get":
					if got := tm.Get(op.key, op.timestamp); got != op.want {
						t.Errorf("Get(%q, %d) = %q, want %q", op.key, op.timestamp, got, op.want)
					}
				}
			}
		})
	}
}
