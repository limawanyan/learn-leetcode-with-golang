package _0_PowxN

// 暴力解法 时间复杂度O（n） 空间复杂度O（1）
func myPow(x float64, n int) float64 {
	if n < 0{
		x = 1/x
		n = -n
	}
	var ans float64 = 1
	for i:=0;i < n;i++{
		ans *= x
	}
	return ans
}

// 分治 迭代
func myPow2(x float64,n int) float64 {
	// 如果n=0直接返回1
	if n == 0{
		return 1
	}
	// 如果n为负数进行处理
	if n < 0{
		x = 1/x
		n = -n
	}
	// 返回值
	var pow float64 = 1
	// 当n小于0时结束循环
	for n > 0 {
		// 如果n为奇数，返回值还要乘剩下的x
		if n & 1 == 1{
			pow *= x
		}
		x *= x
		// 相当于n/2
		n = n>>1
	}
	return pow
}

// 二分法 递归
func myPow3(x float64,n int) float64 {
	// 当n=0结束递归
	if n == 0 {
		return 1
	}
	newX := x
	newN := n
	if n < 0{
		newX = 1/newX
		newN = -n
	}
	temp := myPow3(newX,newN/2)
	res := temp * temp
	// 奇数情况
	if newN % 2 != 0 {
		res *= newX
	}
	return res
}