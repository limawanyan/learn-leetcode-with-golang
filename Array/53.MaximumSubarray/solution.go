package _3_MaximumSubarray

import "math"

// 动态规划
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 1;i < len(nums);i++{
		// 如果当前元素+当前元素之前的最连续和大于当前元素
		// 就相加,否则保持现有元素不变
		// 比如之前的最大连续和为-1，当前元素为1，则最大元素和为1
		// 如果之前的最大连续和为1，当前元素为-1，修改当前元素和为0
		if nums[i] < nums[i] + nums[i-1]{
			nums[i] = nums[i] + nums[i-1]
		}
		// 保存全局最大连续和
		if nums[i] > maxSum{
			maxSum = nums[i]
		}
	}
	return maxSum
}

// 贪心
func maxSubArray2(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	// 全局最优连续和
	res := math.MinInt32
	// 当前连续和
	count := 0
	for i := 0;i < len(nums);i++{
		count += nums[i]
		// 当前连续和大于全局则更新全局的最优连续和
		if count > res{
			res = count
		}
		// 如果当前连续和小于0,舍弃之前的重新开始计算
		if count < 0{
			count = 0
		}
	}
	return res
}

// 暴力破解
func maxSubArray3(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	result := math.MinInt32
	count := 0
	for i := 0;i < len(nums);i++{
		count = 0
		for j := i;j < len(nums);j++{
			count += nums[j]
			if count > result{
				result = count
			}
		}
	}
	return result
}

// 返回子序列
func maxSubArray4(nums []int) ([]int,int) {
	if len(nums) == 1 {
		return nums,0
	}
	res := math.MinInt32
	count := 0
	left,right := 0,0
	for i:=0;i < len(nums);i++{
		count += nums[i]
		if count > res {
			res = count
			right = i
		}
		if count < 0 {
			count = 0
			left = i+1
		}
	}
	return nums[left:right+1],res
}