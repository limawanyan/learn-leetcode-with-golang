package Offer54_TheKthLargestNode

import (
	. "learn-leetcode-with-golang/BinaryTree"
)

// 递归 时间 空间复杂度 O（N）
func kthLargest(root *TreeNode, k int) int {
	count := make([]int,0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode){
		// dfs遍历右子节点
		if root.Right != nil{
			dfs(root.Right)
		}
		// 节点值存储
		if root != nil{
			count = append(count,root.Val)
		}
		// dfs遍历左子节点
		if root.Left != nil{
			dfs(root.Left)
		}
	}
	dfs(root)
	return count[k-1]
}

// 迭代
func kthLargest2(root *TreeNode, k int) int {
	if root == nil{
		return 0
	}
	stack := []*TreeNode{root}
	nums := []int{}
	for len(stack) != 0{
		for root != nil{
			stack = append(stack,root)
			root = root.Right
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums,root.Val)
		root = root.Left
	}
	return nums[k-1]
}