---
link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
tags:
  - Easy
  - Array
  - Dynamic_Programming
---
## Description
You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day.

You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the future** to sell that stock.

Return *the maximum profit you can achieve from this transaction*. If you cannot achieve any profit, return `0`.

**Example 1:**

```
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
```

**Example 2:**

```
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
```

**Constraints:**

- `1 <= prices.length <= 10^5`
- `0 <= prices[i] <= 10^4`

## Solution
### Approach 1
**Two Pointers Greedy**: Track the minimum buy price while iterating, calculate maximum profit at each sell point.

**Principle:**
Use two pointers where the left pointer tracks the minimum buy price seen so far, and the right pointer iterates through each day's price. If the current price is lower than the minimum, update the minimum; otherwise, calculate the profit and update the maximum.

**Steps:**
1. Initialize left pointer `l=0` (buy day), right pointer `r=1` (sell day), max profit `res=0`
2. Iterate through the array; if `prices[r] < prices[l]`, a lower buy price is found, move `l` to `r`
3. Otherwise, calculate profit `prices[r] - prices[l]` and update max profit
4. Return the maximum profit

```embed-go
PATH: "vault://leetcode/0101-0200/0121_best_time_to_buy_and_sell_stock/solution.go"
TITLE: "leetcode 121. Best Time to Buy and Sell Stock"
```
