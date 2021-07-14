package _5_ThreeSum

import (
	"sort"
)

// 双指针 + 排序
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0;i < len(nums)-2;i++ {
		n1 :=nums[i]
		// 如果当前元素已经大于0，直接跳出循环
		if n1 > 0{
			break
		}
		// 如果与上一个元素相同跳过
		if i > 0 && n1 == nums[i - 1] {
			continue
		}
		// 左右两个指针
		l,r := i+1,len(nums)-1
		for l < r{
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
			}else{
				r--
			}
		}
	}
	return result
}