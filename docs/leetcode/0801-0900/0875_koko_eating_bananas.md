---
link: https://leetcode.com/problems/koko-eating-bananas/
tags:
  - Medium
  - Array
  - Binary_Search
---
## Description
Koko loves to eat bananas. There are `n` piles of bananas, the `ith` pile has `piles[i]` bananas. The guards have gone and will come back in `h` hours.

Koko can decide her bananas-per-hour eating speed of `k`. Each hour, she chooses some pile of bananas and eats `k` bananas from that pile. If the pile has less than `k` bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return *the minimum integer* `k` *such that she can eat all the bananas within* `h` *hours*.

**Example 1:**

```
Input: piles = [3,6,7,11], h = 8
Output: 4
```

**Example 2:**

```
Input: piles = [30,11,23,4,20], h = 5
Output: 30
```

**Example 3:**

```
Input: piles = [30,11,23,4,20], h = 6
Output: 23
```

**Constraints:**

- `1 <= piles.length <= 104`
- `piles.length <= h <= 109`
- `1 <= piles[i] <= 109`

## Solution
### Approach 1
**Method Name**: Binary Search

**Principle:**
Perform binary search on the possible speed range to find the minimum speed that meets the time requirement. The minimum speed is 1, and the maximum speed is the size of the largest pile.

**Steps:**
1. Initialize left boundary `l = 1`, right boundary `r` as the maximum value in the piles
2. While `l < r`, calculate middle speed `m = l + (r-l)/2`
3. Calculate the time `cost` required to eat all bananas at speed `m`
4. If `cost <= h`, we can try slower speed, set `r = m`
5. If `cost > h`, we need faster speed, set `l = m + 1`
6. Return `l` when `l == r`

**Example:**
- Input: `piles = [3,6,7,11], h = 8`
- Initial range: `l = 1, r = 11`
- First middle value `m = 6`, time is `1+1+2+2=6 <= 8`, adjust `r = 6`
- Second middle value `m = 3`, time is `1+2+3+4=10 > 8`, adjust `l = 4`
- Third middle value `m = 5`, time is `1+2+2+3=8 <= 8`, adjust `r = 5`
- Fourth `l = 4, r = 5`, middle value `m = 4`, time is `1+2+2+3=8 <= 8`, adjust `r = 4`
- Now `l == r == 4`, return 4

```embed-go
PATH: "vault://leetcode/0801-0900/0875_koko_eating_bananas/solution.go"
TITLE: "leetcode 875. Koko Eating Bananas"
```
