package leetcode0460

import "testing"

func TestLFUCacheExamples(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		ops      []struct {
			op    string // "get" or "put"
			key   int
			value int  // for put operations
			want  int  // expected return value
		}
	}{
		{
			"example_1",
			2,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"get", 1, 0, 1},
				{"put", 3, 3, 0},
				{"get", 2, 0, -1},
				{"get", 3, 0, 3},
				{"put", 4, 4, 0},
				{"get", 1, 0, -1},
				{"get", 3, 0, 3},
				{"get", 4, 0, 4},
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			lfu := Constructor(tst.capacity)
			for _, op := range tst.ops {
				switch op.op {
				case "get":
					got := lfu.Get(op.key)
					if got != op.want {
						t.Errorf("Get(%d) = %d, want %d", op.key, got, op.want)
					}
				case "put":
					lfu.Put(op.key, op.value)
				}
			}
		})
	}
}

func TestLFUCache(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		ops      []struct {
			op    string // "get" or "put"
			key   int
			value int  // for put operations
			want  int  // expected return value
		}
	}{
		{
			"capacity_1_put_then_get",
			1,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"get", 1, 0, 10},
			},
		},
		{
			"capacity_1_get_non_existent",
			1,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"get", 5, 0, -1},
			},
		},
		{
			"capacity_1_eviction",
			1,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"get", 1, 0, -1},
				{"get", 2, 0, 20},
			},
		},
		{
			"update_existing_key",
			2,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 1, 100, 0},
				{"get", 1, 0, 100},
			},
		},
		{
			"frequency_increment_on_put_update",
			3,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0}, // freq 1
				{"put", 2, 20, 0}, // freq 1
				{"put", 3, 30, 0}, // freq 1
				{"put", 1, 100, 0}, // freq 2
				{"put", 2, 200, 0}, // freq 2
				{"put", 4, 40, 0}, // should evict 3
				{"get", 3, 0, -1},
				{"get", 1, 0, 100},
				{"get", 2, 0, 200},
			},
		},
		{
			"tie_breaker_lru",
			2,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"get", 1, 0, 10}, // now 2 is LRU among freq 1
				{"put", 3, 30, 0}, // should evict 2
				{"get", 2, 0, -1},
				{"get", 1, 0, 10},
				{"get", 3, 0, 30},
			},
		},
		{
			"same_freq_lru_order",
			3,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"put", 3, 30, 0},
				{"get", 1, 0, 10},
				{"get", 2, 0, 20},
				{"get", 3, 0, 30}, // all freq 2, LRU order: 1, 2, 3
				{"put", 4, 40, 0}, // should evict 1
				{"get", 1, 0, -1},
				{"get", 2, 0, 20},
				{"get", 3, 0, 30},
				{"get", 4, 0, 40},
			},
		},
		{
			"max_key_and_value",
			1,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 100000, 1000000000, 0},
				{"get", 100000, 0, 1000000000},
			},
		},
		{
			"zero_key_and_value",
			1,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 0, 0, 0},
				{"get", 0, 0, 0},
			},
		},
		{
			"multiple_evictions_same_freq",
			2,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"get", 1, 0, 10}, // freq 2 for 1
				{"put", 3, 30, 0}, // evict 2
				{"get", 2, 0, -1},
				{"put", 4, 40, 0}, // evict 3
				{"get", 3, 0, -1},
				{"get", 1, 0, 10},
				{"get", 4, 0, 40},
			},
		},
		{
			"get_then_put_same_key",
			2,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"get", 1, 0, 10},
				{"put", 1, 100, 0}, // freq becomes 3
				{"put", 3, 30, 0}, // should evict 2
				{"get", 2, 0, -1},
				{"get", 1, 0, 100},
			},
		},
		{
			"rapid_put_operations",
			3,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"put", 3, 3, 0},
				{"put", 4, 4, 0}, // evict 1
				{"put", 5, 5, 0}, // evict 2
				{"put", 6, 6, 0}, // evict 3
				{"get", 1, 0, -1},
				{"get", 2, 0, -1},
				{"get", 3, 0, -1},
				{"get", 4, 0, 4},
				{"get", 5, 0, 5},
				{"get", 6, 0, 6},
			},
		},
		{
			"frequently_accessed_not_evicted",
			3,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"put", 3, 30, 0},
				{"get", 1, 0, 10},
				{"get", 1, 0, 10},
				{"get", 1, 0, 10}, // freq 4
				{"get", 2, 0, 20}, // freq 2
				{"get", 3, 0, 30}, // freq 2
				{"put", 4, 40, 0}, // should evict 2 (LRU between 2 and 3)
				{"get", 2, 0, -1},
				{"get", 1, 0, 10},
				{"get", 3, 0, 30},
				{"get", 4, 0, 40},
			},
		},
		{
			"put_same_key_multiple_times",
			2,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 1, 20, 0},
				{"put", 1, 30, 0}, // freq 3
				{"put", 2, 100, 0},
				{"put", 3, 200, 0}, // should evict 2
				{"get", 2, 0, -1},
				{"get", 1, 0, 30},
			},
		},
		{
			"complex_mixed_operations",
			3,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},   // freq 1
				{"put", 2, 20, 0},   // freq 1
				{"get", 1, 0, 10},   // freq 2
				{"put", 3, 30, 0},   // freq 1
				{"get", 2, 0, 20},   // freq 2
				{"get", 3, 0, 30},   // freq 2
				{"put", 4, 40, 0},   // evict 1 (all freq 2, 1 is LRU)
				{"get", 1, 0, -1},
				{"get", 2, 0, 20},   // freq 3
				{"get", 3, 0, 30},   // freq 3
				{"get", 4, 0, 40},   // freq 2
				{"put", 5, 50, 0},   // evict 4
				{"get", 4, 0, -1},
				{"get", 5, 0, 50},
			},
		},
		{
			"lru_order_maintained_after_gets",
			4,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"put", 3, 30, 0},
				{"put", 4, 40, 0},
				{"get", 4, 0, 40},
				{"get", 3, 0, 30},
				{"get", 2, 0, 20},
				{"get", 1, 0, 10}, // LRU order now: 4, 3, 2, 1 (most recent last)
				{"put", 5, 50, 0}, // should evict 4
				{"get", 4, 0, -1},
				{"get", 1, 0, 10},
			},
		},
		{
			"min_capacity",
			1,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 1, 0},
				{"get", 1, 0, 1},
				{"put", 2, 2, 0},
				{"get", 1, 0, -1},
				{"get", 2, 0, 2},
				{"put", 3, 3, 0},
				{"get", 2, 0, -1},
				{"get", 3, 0, 3},
			},
		},
		{
			"put_after_eviction",
			2,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"put", 3, 30, 0}, // evict 1
				{"put", 1, 100, 0}, // evict 2
				{"get", 2, 0, -1},
				{"get", 1, 0, 100},
				{"get", 3, 0, 30},
			},
		},
		{
			"get_non_existent_middle",
			2,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"get", 999, 0, -1},
				{"get", 1, 0, 10},
				{"get", 2, 0, 20},
			},
		},
		{
			"all_keys_same_freq_evict_lru",
			3,
			[]struct {
				op    string
				key   int
				value int
				want  int
			}{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"put", 3, 30, 0},
				{"get", 1, 0, 10},
				{"get", 2, 0, 20},
				{"get", 3, 0, 30}, // all freq 2, LRU: 1, 2, 3
				{"get", 2, 0, 20},
				{"get", 1, 0, 10}, // now LRU: 3, 2, 1
				{"put", 4, 40, 0}, // evict 3
				{"get", 3, 0, -1},
				{"get", 1, 0, 10},
				{"get", 2, 0, 20},
				{"get", 4, 0, 40},
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			lfu := Constructor(tst.capacity)
			for _, op := range tst.ops {
				switch op.op {
				case "get":
					got := lfu.Get(op.key)
					if got != op.want {
						t.Errorf("Get(%d) = %d, want %d", op.key, got, op.want)
					}
				case "put":
					lfu.Put(op.key, op.value)
				}
			}
		})
	}
}
