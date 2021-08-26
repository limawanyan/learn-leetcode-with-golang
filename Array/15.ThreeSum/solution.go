package _5_ThreeSum

import (
	"sort"
)

// 排序 + 双指针
func threeSum(nums []int) [][]int {
	// 对元素进行排序
	sort.Ints(nums)
	// 返回结果
	result := [][]int{}
	// 循环长度-2次，因为最后左右两个指针元素可以不用循环
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		// 如果当前元素已经大于0，直接跳出循环 剪枝
		if n1 > 0 {
			break
		}
		// 如果与上一个元素相同跳过 剪枝
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		// 左右两个指针
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			// 如果三数之和等于0，保存结果并跳过重复元素
			if n1+n2+n3 == 0 {
				result = append(result, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
				// 小于0说明元素太小了，左指针向大的元素方向移动
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return result
}

// map加速查询
func threeSum2(nums []int) [][]int {
	// 对数组进行排序
	sort.Ints(nums)
	// 存储nums元素数量
	var dataCountMap = make(map[int]int)
	// 将nums数组中的元素添加到map中
	for _,v := range nums {
		// 已经存在数量置为1，否则++
		if _,ok:= dataCountMap[v];!ok{
			dataCountMap[v] = 1
		}else{
			dataCountMap[v]++
		}
	}
	// 用于存储结果，达到去重作用
	var solutionMap = make(map[[2]int]struct{})
	for i := 0;i < len(nums) - 2;i++{
		for j := i+1;j < len(nums)-1;j++{
			// a+b+c = 0 c = -a-b
			expected := -nums[i] - nums[j]
			// 判断map中是否存在符合的元素，且元素要大于等于j索引上的元素，因为开始已经将数组排序了
			if num,ok := dataCountMap[expected];ok && expected >= nums[j]{
				// 如果和i或j索引元素一样将num局部变量--，作用是判断是否有足够数量的元素可供使用，不能使用同一个重复元素
				if expected == nums[i]{
					num--
				}
				if expected == nums[j]{
					num--
				}
				if num > 0 {
					solutionMap[[2]int{nums[i],nums[j]}] = struct{}{}
				}
			}
		}
	}
	var result = make([][]int,len(solutionMap))
	i := 0
	for k := range solutionMap{
		result[i] = []int{k[0],k[1],-k[0]-k[1]}
		i++
	}
	return result
}
