package _12_PathSum

import . "learn-leetcode-with-golang/BinaryTree"

// 广度优先搜
// 时间复杂度 O（N） N为树的节点数
// 空间复杂度 O（N）
func hasPathSum(root *TreeNode, sum int) bool {
	// 如果根节点为空直接返回false
	if root == nil{
		return false
	}
	// 存储节点队列
	queNode := []*TreeNode{root}
	// 存储节点值
	queVal := []int{root.Val}
	for len(queNode) != 0{
		// 获取队列第一个节点
		now := queNode[0]
		queNode = queNode[1:]
		// 获取节点对应的值
		temp := queVal[0]
		queVal =  queVal[1:]
		// 如果当前节点不存在左右子树了，说明为叶子节点
		if now.Left == nil && now.Right == nil{
			// 判断到这个叶子节点的路劲总和是否相等
			if temp == sum{
				return true
			}
			// 如果不相等继续循环其他路径
			continue
		}
		// 如果左右子节点不为空,将节点和值添加到队列中
		if now.Left != nil{
			queNode = append(queNode,now.Left)
			queVal = append(queVal,now.Left.Val + temp)
		}
		if now.Right != nil{
			queNode = append(queNode,now.Right)
			queVal = append(queVal,now.Right.Val + temp)
		}
	}
	return false
}

// 递归 时间复杂度 O（N） 空间复杂度 O（H）
func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	// 判断是否为叶子节点
	if root.Left == nil && root.Right == nil{
		// 判断sum是否等于当前节点的值
		return sum == root.Val
	}
	// dfs递归左右子节点，在sum的基础上减去当前节点值
	return hasPathSum2(root.Left,sum - root.Val) || hasPathSum2(root.Right,sum - root.Val)
}