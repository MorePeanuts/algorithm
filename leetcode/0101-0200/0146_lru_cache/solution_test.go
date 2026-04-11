package leetcode0146

import "testing"

type operation struct {
	name string
	key  int
	val  int
	want int
}

func TestLRUCacheExamples(t *testing.T) {
	tests := []struct {
		name       string
		capacity   int
		operations []operation
	}{
		{
			"example_1",
			2,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"get", 1, 0, 1},
				{"put", 3, 3, 0},
				{"get", 2, 0, -1},
				{"put", 4, 4, 0},
				{"get", 1, 0, -1},
				{"get", 3, 0, 3},
				{"get", 4, 0, 4},
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			cache := Constructor(tst.capacity)
			for _, op := range tst.operations {
				switch op.name {
				case "put":
					cache.Put(op.key, op.val)
				case "get":
					got := cache.Get(op.key)
					if got != op.want {
						t.Errorf("Get(%d) = %d, want %d", op.key, got, op.want)
					}
				}
			}
		})
	}
}

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name       string
		capacity   int
		operations []operation
	}{
		{
			"capacity_1",
			1,
			[]operation{
				{"put", 1, 1, 0},
				{"get", 1, 0, 1},
				{"put", 2, 2, 0},
				{"get", 1, 0, -1},
				{"get", 2, 0, 2},
			},
		},
		{
			"get_non_existent",
			2,
			[]operation{
				{"get", 1, 0, -1},
				{"put", 1, 1, 0},
				{"get", 1, 0, 1},
				{"get", 2, 0, -1},
			},
		},
		{
			"update_existing_key",
			2,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"put", 1, 10, 0},
				{"get", 1, 0, 10},
				{"get", 2, 0, 2},
			},
		},
		{
			"update_makes_key_recent",
			2,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"put", 1, 10, 0}, // 1 becomes recent
				{"put", 3, 3, 0}, // evict 2
				{"get", 1, 0, 10},
				{"get", 2, 0, -1},
				{"get", 3, 0, 3},
			},
		},
		{
			"evict_order_1",
			3,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"put", 3, 3, 0},
				{"put", 4, 4, 0}, // evict 1
				{"get", 1, 0, -1},
				{"get", 2, 0, 2},
				{"get", 3, 0, 3},
				{"get", 4, 0, 4},
			},
		},
		{
			"evict_order_with_get",
			3,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"put", 3, 3, 0},
				{"get", 1, 0, 1},  // 1 becomes recent
				{"put", 4, 4, 0}, // evict 2
				{"get", 1, 0, 1},
				{"get", 2, 0, -1},
				{"get", 3, 0, 3},
				{"get", 4, 0, 4},
			},
		},
		{
			"same_key_put_multiple_times",
			2,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 1, 2, 0},
				{"put", 1, 3, 0},
				{"get", 1, 0, 3},
				{"put", 2, 2, 0},
				{"put", 3, 3, 0}, // evict 1
				{"get", 1, 0, -1},
			},
		},
		{
			"capacity_3000_max",
			3000,
			[]operation{
				{"put", 0, 100, 0},
				{"put", 10000, 99999, 0},
				{"get", 0, 0, 100},
				{"get", 10000, 0, 99999},
			},
		},
		{
			"key_zero_allowed",
			2,
			[]operation{
				{"put", 0, 100, 0},
				{"get", 0, 0, 100},
				{"put", 1, 200, 0},
				{"put", 2, 300, 0}, // evict 0
				{"get", 0, 0, -1},
				{"get", 1, 0, 200},
				{"get", 2, 0, 300},
			},
		},
		{
			"value_zero_allowed",
			2,
			[]operation{
				{"put", 1, 0, 0},
				{"get", 1, 0, 0},
				{"put", 2, 1, 0},
				{"put", 3, 2, 0}, // evict 1
				{"get", 1, 0, -1},
			},
		},
		{
			"get_middle_node",
			4,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"put", 3, 3, 0},
				{"put", 4, 4, 0},
				{"get", 2, 0, 2}, // move 2 to tail
				{"put", 5, 5, 0}, // evict 1
				{"get", 1, 0, -1},
				{"get", 2, 0, 2},
				{"get", 3, 0, 3},
				{"get", 4, 0, 4},
				{"get", 5, 0, 5},
			},
		},
		{
			"alternate_get_put",
			3,
			[]operation{
				{"put", 1, 1, 0},
				{"get", 1, 0, 1},
				{"put", 2, 2, 0},
				{"get", 2, 0, 2},
				{"put", 3, 3, 0},
				{"get", 3, 0, 3},
				{"put", 4, 4, 0}, // evict 1
				{"get", 1, 0, -1},
				{"get", 2, 0, 2},
				{"get", 3, 0, 3},
				{"get", 4, 0, 4},
			},
		},
		{
			"put_same_key_after_eviction",
			2,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"put", 3, 3, 0}, // evict 1
				{"get", 1, 0, -1},
				{"put", 1, 100, 0}, // evict 2
				{"get", 1, 0, 100},
				{"get", 2, 0, -1},
				{"get", 3, 0, 3},
			},
		},
		{
			"lru_chain_access",
			4,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"put", 3, 3, 0},
				{"put", 4, 4, 0},
				{"get", 1, 0, 1}, // order: 2,3,4,1
				{"get", 2, 0, 2}, // order: 3,4,1,2
				{"put", 5, 5, 0}, // evict 3
				{"get", 3, 0, -1},
				{"get", 4, 0, 4},
				{"get", 1, 0, 1},
				{"get", 2, 0, 2},
				{"get", 5, 0, 5},
			},
		},
		{
			"continuous_puts",
			3,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"put", 3, 3, 0},
				{"put", 4, 4, 0},
				{"put", 5, 5, 0},
				{"put", 6, 6, 0},
				{"get", 1, 0, -1},
				{"get", 2, 0, -1},
				{"get", 3, 0, -1},
				{"get", 4, 0, 4},
				{"get", 5, 0, 5},
				{"get", 6, 0, 6},
			},
		},
		{
			"get_head_then_put",
			2,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"get", 1, 0, 1}, // head was 1, now 2 becomes head, 1 moves to tail
				{"put", 3, 3, 0}, // evict 2
				{"get", 1, 0, 1},
				{"get", 2, 0, -1},
				{"get", 3, 0, 3},
			},
		},
		{
			"get_tail_no_change",
			2,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"get", 2, 0, 2}, // already tail, no change
				{"put", 3, 3, 0}, // evict 1
				{"get", 1, 0, -1},
				{"get", 2, 0, 2},
				{"get", 3, 0, 3},
			},
		},
		{
			"complex_scenario_1",
			3,
			[]operation{
				{"put", 1, 10, 0},
				{"put", 2, 20, 0},
				{"get", 1, 0, 10},
				{"put", 3, 30, 0},
				{"get", 2, 0, 20},
				{"put", 4, 40, 0}, // evict 1
				{"get", 1, 0, -1},
				{"get", 3, 0, 30},
				{"get", 4, 0, 40},
				{"put", 2, 200, 0}, // update 2
				{"put", 5, 50, 0},  // evict 3
				{"get", 3, 0, -1},
				{"get", 2, 0, 200},
				{"get", 4, 0, 40},
				{"get", 5, 0, 50},
			},
		},
		{
			"complex_scenario_2",
			4,
			[]operation{
				{"put", 10, 100, 0},
				{"put", 20, 200, 0},
				{"put", 30, 300, 0},
				{"put", 40, 400, 0},
				{"get", 10, 0, 100},
				{"get", 30, 0, 300},
				{"put", 50, 500, 0}, // evict 20
				{"put", 60, 600, 0}, // evict 40
				{"get", 20, 0, -1},
				{"get", 40, 0, -1},
				{"get", 10, 0, 100},
				{"get", 30, 0, 300},
				{"get", 50, 0, 500},
				{"get", 60, 0, 600},
			},
		},
		{
			"single_operation",
			1,
			[]operation{
				{"put", 5, 55, 0},
			},
		},
		{
			"put_after_capacity_reached_multiple_times",
			2,
			[]operation{
				{"put", 1, 1, 0},
				{"put", 2, 2, 0},
				{"put", 3, 3, 0}, // evict 1
				{"put", 4, 4, 0}, // evict 2
				{"put", 5, 5, 0}, // evict 3
				{"put", 6, 6, 0}, // evict 4
				{"get", 1, 0, -1},
				{"get", 2, 0, -1},
				{"get", 3, 0, -1},
				{"get", 4, 0, -1},
				{"get", 5, 0, 5},
				{"get", 6, 0, 6},
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			cache := Constructor(tst.capacity)
			for _, op := range tst.operations {
				switch op.name {
				case "put":
					cache.Put(op.key, op.val)
				case "get":
					got := cache.Get(op.key)
					if got != op.want {
						t.Errorf("Get(%d) = %d, want %d", op.key, got, op.want)
					}
				}
			}
		})
	}
}
