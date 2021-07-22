package _8_ValidateBinarySearchTree

import "math"

type TreeNode struct {
	  Val int
	  Left *TreeNode
	  Right *TreeNode
}

// 递归
func isValidBST(root *TreeNode) bool{
	return helper(root,math.MinInt64,math.MaxInt64)
}

func helper(root *TreeNode,lower,upper int) bool {
	// 递归到节点为空结束
	if root == nil{
		return true
	}
	// 左节点：如果左节点大于等于根节点 false
	// 右节点：如果右节点小于等于跟节点 false
	if root.Val <= lower || root.Val >= upper{
		return false
	}
	// 递归左右节点
	return helper(root.Left,lower,root.Val) && helper(root.Right,root.Val,upper)
}

// 中序遍历
func isValidBST2(root *TreeNode) bool {
	// 栈保存节点
	stack := []*TreeNode{}
	inorder := math.MinInt64
	// 当栈所有节点被取出 以及 节点为空时结束循环
	for len(stack) > 0 || root != nil{
		// 先把所有节点压入栈
		for root != nil{
			stack = append(stack,root)
			root = root.Left
		}
		// pop栈，判断是否符合有序二叉树
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder{
			return false
		}
		// 保存根节点值，切换到右子树
		inorder = root.Val
		root = root.Right
	}
	return true
}