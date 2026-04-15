// Package leetcode0981 solves LeetCode 981. Time Based Key-Value Store
package leetcode0981

type TimeMap struct {
	timestamps map[string][]int
	values     map[string][]string
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]int), make(map[string][]string)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timestamps[key] = append(this.timestamps[key], timestamp)
	this.values[key] = append(this.values[key], value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	pos := searchPos(this.timestamps[key], timestamp)
	if pos == -1 {
		return ""
	}
	return this.values[key][pos]
}

func searchPos(timestamps []int, target int) int {
	l, r := 0, len(timestamps)
	for l < r {
		m := l + (r-l)/2
		if target < timestamps[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
