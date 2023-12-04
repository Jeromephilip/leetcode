/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
    maxProfit := 0
	lpoint := 0
	rpoint := 0
	for rpoint < len(prices) {
		if prices[lpoint] > prices[rpoint] {
			lpoint = rpoint
		} else if (prices[rpoint] - prices[lpoint]) > maxProfit {
			maxProfit = prices[rpoint] - prices[lpoint]
		}
		rpoint++
	}

	return maxProfit
}
// @lc code=end

