package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	input.Scan()
	arr := strings.Split(input.Text()," ")
	nums := []int{}
	for i := 0;i < len(arr);i++ {
		temp,_ := strconv.Atoi(arr[i])
		nums = append(nums,temp)
	}
	// 返回结果
	res := make([][2]int,0)
	for i := 0;i < len(nums);i++{
		res = append(res,[2]int{-1,-1})
	}
	// 栈
	stack := []int{}
	for i := 0;i < len(nums);i++{
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			res[stack[len(stack)-1]][1] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,i)
	}
	stack = []int{}
	for i := len(nums)-1;i >= 0;i--{
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			res[stack[len(stack)-1]][0] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,i)
	}
	for i := 0;i < len(res);i++ {
		fmt.Println(res[i][0],res[i][1])
	}
}

