package _43_DiameterOfBinaryTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	// 最大直径长度
	var maxDiameter int
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int{
		// 节点为空返回路径长度0
		if root == nil{
			return 0
		}else{
			// 递归左子树路径
			left := dfs(root.Left)
			// 递归右子树路径
			right := dfs(root.Right)
			// 如果当前节点左右子树直径长度大于最大直径长度则更新
			if left + right > maxDiameter{
				maxDiameter = left + right
			}
			return max(left,right)+1
		}
	}
	dfs(root)
	return maxDiameter
}

func max(a,b int) int{
	if a > b{
		return a
	}
	return b
}