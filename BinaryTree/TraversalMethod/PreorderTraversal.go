package TraversalMethod

import (
	"fmt"
	. "learn-leetcode-with-golang/BinaryTree"
)

// 递归前序遍历
func preorderTraversalR(root *TreeNode){
	if root == nil{
		return
	}
	// 根节点
	fmt.Println(root.Val)
	// 左节点
	preorderTraversalR(root.Left)
	// 右节点
	preorderTraversalR(root.Right)
}

// 前序非递归遍历
func preorderTraversal(root *TreeNode) []int{
	if root == nil {
		return nil
	}
	// 最终返回结果
	result := make([]int,0)
	// 堆栈存储中间节点
	stack := make([]*TreeNode,0)
	// 当二叉树左右遍历节点都为空，堆栈无节点退出循环
	for root != nil || len(stack) != 0 {
		// 循环 保存根节点 将节点保存栈中 继续读取左节点
		for root != nil {
			result = append(result,root.Val)
			stack = append(stack,root)
			root = root.Left
		}
		// 上方循环无左节点了 执行stack pop操作 赋值 root 右节点继续遍历
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}