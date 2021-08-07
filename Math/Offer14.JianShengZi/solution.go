package Offer14_JianShengZi

import "math"

// 数学
func cuttingRope(n int) int {
	if n <= 3 {
		return n-1
	}
	quotient := n/3
	remainder := n%3
	if remainder == 0{
		return int(math.Pow(3,float64(quotient)))
	}else if remainder == 1 {
		return int(math.Pow(3,float64(quotient-1))) * 4
	}
	return  int(math.Pow(3,float64(quotient))) * 2
}

// 贪心 时间复杂度O(n) 空间复杂度O(1)
func cuttingRope2(n int) int {
	if n == 2{
		return 1
	}
	if n == 3{
		return 2
	}
	if n == 4{
		return 4
	}
	result := 1
	for n > 4{
		result *= 3
		n -= 3
	}
	result *= n
	return result
}

// 动态规划 时间复杂度：O(n^2) 空间复杂度：O(n)
func cuttingRope3(n int) int {
	dp := make([]int,n+1)
	dp[2] = 1
	max := func(a,b int) int {
		if a > b{
			return a
		}
		return b
	}
	// 外层循环为绳子长度，内层循环获取比较最大乘积
	// 假如现在为长度10的绳子
	// 指剪一刀划分为两根绳子：(i-j)*j: 9*1,8*2,7*3,6*4...
	// 也可以根据前面绳子已经得到的最大乘积dp[i-j]*j: dp[9]*1,dp[8]*2,dp[7]*3...
	// 然后比较这两者，获取最大乘积方案
	for i := 3;i <= n;i++{
		for j := 1;j < i;j++{
			dp[i] = max(dp[i],max(j*(i-j),dp[i-j]*j))
		}
	}
	return dp[n]
}