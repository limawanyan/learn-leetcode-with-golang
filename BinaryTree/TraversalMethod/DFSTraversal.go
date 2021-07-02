package TraversalMethod

import . "learn-leetcode-with-golang/BinaryTree"

// DFS深度搜索-从上到下
func DfsTraversal(root *TreeNode) []int{
	result := make([]int,0)
	dfs(root,&result)
	return result
}

// 深度遍历，结果指针作为参数传入函数内部
func dfs(root *TreeNode,result *[]int)  {
	if root == nil{
		return
	}
	// 保存当前根节点值
	*result = append(*result,root.Val)
	// 递归左节点
	dfs(root.Left,result)
	// 递归右节点
	dfs(root.Right,result)
}

// DFS深度搜索-从下向上（分治法）
func Dfs2Traversal(root *TreeNode) []int{
	result := divideAndConquer(root)
	return result
}

func divideAndConquer(root *TreeNode) []int {
	result := make([]int,0)
	if root == nil{
		return result
	}
	// 分治
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)
	// 合并
	result = append(result,root.Val)
	result = append(result,left...)
	result = append(result,right...)
	return result
}

