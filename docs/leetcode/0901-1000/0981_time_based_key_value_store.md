---
link: https://leetcode.com/problems/time-based-key-value-store/
tags:
  - Medium
  - Design
  - Hash_Table
  - String
  - Binary_Search
---
## Description
Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

Implement the `TimeMap` class:

- `TimeMap()` Initializes the object of the data structure.
- `void set(String key, String value, int timestamp)` Stores the key `key` with the value `value` at the given time `timestamp`.
- `String get(String key, int timestamp)` Returns a value such that `set` was called previously, with `timestamp_prev <= timestamp`. If there are multiple such values, it returns the value associated with the largest `timestamp_prev`. If there are no values, it returns `""`.

**Example 1:**

```
Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]

Explanation
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
timeMap.get("foo", 1);         // return "bar"
timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
timeMap.get("foo", 4);         // return "bar2"
timeMap.get("foo", 5);         // return "bar2"
```

**Constraints:**

- `1 <= key.length, value.length <= 100`
- `key` and `value` consist of lowercase English letters and digits.
- `1 <= timestamp <= 107`
- All the timestamps `timestamp` of `set` are strictly increasing.
- At most `2 * 105` calls will be made to `set` and `get`.

## Solution
### Approach 1
**Method Name**: Hash Table + Binary Search

**Principle:**
Use two hash tables to store the timestamp array and value array for each key respectively. Leveraging the strictly increasing property of timestamps in set operations, quickly locate the position closest to and not greater than the target timestamp through binary search.

**Steps:**
1. Initialize two hash tables, `timestamps` stores the mapping from key to timestamp array, `values` stores the mapping from key to value array
2. `set` operation: Directly append the timestamp and value to the end of the corresponding key's array (leveraging the strictly increasing property of timestamps)
3. `get` operation: Use binary search to find the position in the timestamp array that is closest to and not greater than the target timestamp, return the value at the corresponding position

**Example:**
- Input: set("foo", "bar", 1) → timestamps["foo"] = [1], values["foo"] = ["bar"]
- Input: get("foo", 3) → Binary search finds position 0 (timestamp 1 ≤ 3) → Output "bar"
- Input: set("foo", "bar2", 4) → timestamps["foo"] = [1, 4], values["foo"] = ["bar", "bar2"]
- Input: get("foo", 4) → Binary search finds position 1 (timestamp 4 ≤ 4) → Output "bar2"

```embed-go
PATH: "vault://leetcode/0901-1000/0981_time_based_key_value_store/solution.go"
TITLE: "leetcode 981. Time Based Key-Value Store"
```
