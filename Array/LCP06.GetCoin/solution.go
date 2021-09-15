package LCP06_GetCoin

func minCount(coins []int) int {
	res := 0
	for i := 0;i < len(coins);i++{
		// 位运算 向右移动一位相当于除2,和1按位与相当于除2取余
		res += (coins[i] + 1) >> 1
	}
	return res
}

