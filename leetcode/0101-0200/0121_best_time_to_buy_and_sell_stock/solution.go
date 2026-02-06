// Package leetcode0121 solves LeetCode 121. Best Time to Buy and Sell Stock
package leetcode0121

func maxProfit(prices []int) int {
	res := 0
	for l, r := 0, 1; r < len(prices); r++ {
		if prices[r] < prices[l] {
			l = r
		} else {
			res = max(res, prices[r]-prices[l])
		}
	}
	return res
}
