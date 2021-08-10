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
		// 如果当前元素已经大于0，直接跳出循环
		if n1 > 0 {
			break
		}
		// 如果与上一个元素相同跳过
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

// 练习
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		// 进入当前元素已经大于0直接退出循环
		if n1 > 0 {
			break
		}
		// 如果与上一个元素相同跳过
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		l,r := i+1,len(nums)-1
		for l < r {
			n2,n3 := nums[l],nums[r]
			if n1 + n2 + n3 == 0{
				result = append(result,[]int{n1,n2,n3})
				for l < r && nums[l] == n2{
					l++
				}
				for l < r && nums[r] == n3{
					r--
				}
			}else if n1 + n2 + n3 < 0{
				l++
			}else {
				r++
			}
		}
	}
	return result
}
