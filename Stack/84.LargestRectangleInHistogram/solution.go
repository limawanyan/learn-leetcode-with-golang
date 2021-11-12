package main

import "fmt"

func main()  {
	fmt.Println(largestRectangleArea([]int{1}))
}


func largestRectangleArea(heights []int) int {
	res := 0
	// 单调栈 初始化一个边界元素-1 用于计算第一个矩形面积
	stack := []int{-1}
	// 添加一个伪元素用于计算原始数组最后一个矩形面积
	heights = append(heights,0)
	n := len(heights)
	for i := 0;i < n;i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]]{
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			w := i - stack[len(stack)-1] - 1
			h := heights[mid]
			res = max(res,w * h)
		}
		stack = append(stack,i)
	}
	return res
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
