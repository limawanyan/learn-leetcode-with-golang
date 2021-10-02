package _8_Subsets

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	// 保存最终结果
	result := make([][]int,0)
	// 保存中间结果
	list := make([]int,0)
	backtrack(nums,0,list,&result)
	return result
}

// nums 给定的集合
// pos 下次添加到结合中的元素位置索引
// list 临时结果集合（每次需要复制保存）
// result 最终结果
func backtrack(nums []int,pos int,list []int,result *[][]int)  {
	// 把临时结果复制出来保存到最终结果
	ans := make([]int,len(list))
	copy(ans,list)
	*result = append(*result,ans)
	// 选择、处理结果、再撤销选择
	for i := pos;i < len(nums);i++ {
		// [1][2][3]
		// [1、2][1、3]
		list = append(list,nums[i])
		fmt.Println(list)
		backtrack(nums,i+1,list,result)
		list = list[0 : len(list) - 1]
	}
	fmt.Println("-----")
}

// 迭代位运算
// 时间复杂度 0（n * 2^n） 空间复杂度 O（n）
func subsets2(nums []int) (res [][]int) {
	n := len(nums)
	// 相当于2^n-1
	for mask := 0;mask < 1<<n;mask++{
		set := []int{}
		// nums = [1、2、3]
		for i,v := range nums {
			fmt.Println(mask,">>",i,":",mask >> i,":",mask >> i & 1)
			if mask >> i & 1 > 0 {
				//fmt.Println(i,":",mask >> i,":",mask >> i & 1,":",v)
				set = append(set,v)
			}
		}
		fmt.Println(set)
		fmt.Println("------")

		res = append(res,append([]int(nil),set...))
	}
	return
}

// dfs 回溯
func subsets3(nums []int) (res [][]int) {
	var dfs func(i int,list []int)
	dfs = func(i int, list []int) {
		// i 为当前递归考察的索引元素 索引越界时结束 将当前递归分支结果添加到结果集
		if i == len(nums) {
			res = append(res,append([]int(nil),list...))
			return
		}
		// 选择当前元素
		list = append(list,nums[i])
		dfs(i+1,list)
		// 撤销，不选择当前元素
		list = list[:len(list)-1]
		dfs(i+1,list)
	}

	dfs(0,[]int{})
	return
}