package Array

import "math"

// 动态规划，时间复杂度O（n），空间复杂度O（1）
func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0{return 0}
	// 初始化动态数组
	dp := make([][]int,length)
	for i := 0;i < length;i++{
		dp[i] = make([]int,2)
	}
	// dp[0][0]初始化，假设本金为0，第一天买入股票
	// dp[x][0]在后续的循环中保存当前最低的股票价格
	dp[0][0] = -prices[0]
	// dp[0][1]初始化，表示第一天不买股票最大利润
	// dp[x][1]在后续的循环中保存当前股票价格能获取的最大利润
	dp[0][1] = 0
	var max func(a,b int) int
	max = func(a, b int) int {
		if a > b{return a}
		return b
	}
	for i := 1;i < length;i++{
		// 第i天持有股票可以有两种状态
		// 1.保持现状，仍旧持有昨天的股票，即:dp[i-1][0]
		// 1.第i天买入股票，所得现金就是：-prices[i]
		// 最终我们要通过比较两者，获取买入价格最低的，也就是-prices[i]最大的值
		dp[i][0] = 	max(dp[i-1][0],-prices[i])
		// 第i天不持有股票也可以有两种状态
		// 1.第i-1天就不持有股票，保持现状，所得金额就是昨天不持有股票所得现金:dp[i-1][i]（可以理解为是前一天把股票卖出去所得的金额）
		// 2.第i天把股票卖出去，所得金额就是当前股票价格+dp[i-i][0] （这里为什么是加，因为我们的dp[i][0]是买入股票的价格负数，所以其实就相当于当前价格-买入价格）
		// 最终我们要通过比较两者，获取我们能得到的最大利润
		dp[i][1] = max(dp[i-1][1],dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

// 遍历一次（贪心），时间复杂度O（n），空间复杂度O（1）
func maxProfit2(prices []int) int {
	// 保存历史最低价格
	minValue := math.MaxInt64
	// 能够获取的最大利润
	maxValue := 0
	for i := 0;i < len(prices);i++{
		// 当前股票价格是否低于历史最低价格
		if prices[i] < minValue{
			minValue = prices[i]
			// 当前卖出股票利润是否最大
		}else if prices[i] - minValue > maxValue{
			maxValue = prices[i] - minValue
		}
	}
	return maxValue
}