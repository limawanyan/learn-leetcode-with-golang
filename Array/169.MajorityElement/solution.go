package _69_MajorityElement

import "sort"

// 摩尔投票法
func majorityElement(nums []int) int {
	major,count := 0,0
	for _,num := range nums{
		if count == 0{
			major = num
		}
		if major == num{
			count++
		}else {
			count--
		}
	}
	return major
}

// 暴力解法 时间复杂度O(n^2)
func majorityElement2(nums []int) int {
	res := 0
	count := 0
	for _,i := range nums{
		temp := 0
		for _,j := range nums{
			if i == j{
				temp++
			}
		}
		if temp > count{
			count = temp
			res = i
		}
	}
	return res
}

// 排序法 时间复杂度O(nlogn)
func majorityElement3(nums []int) int {
	sort.Ints(nums)
	index := len(nums) / 2
	return nums[index]
}

// map
func majorityElement4(nums []int) int {
	hash := make(map[int]int)
	n := len(nums) / 2
	for _,i := range nums{
		hash[i]++
		count,_ := hash[i]
		if count > n{
			return i
		}
	}
	return -1
}