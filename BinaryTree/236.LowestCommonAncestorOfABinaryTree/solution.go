package _36_LowestCommonAncestorOfABinaryTree

import ."learn-leetcode-with-golang/BinaryTree"

// 递归  时间空间复杂度O（n）
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 节点为nil递归结束
	if root == nil{
		return root
	}
	// root为两节点父节点
	if root.Val == p.Val || root.Val == q.Val{
		return  root
	}
	// 递归左子树查找节点
	left := lowestCommonAncestor(root.Left,p,q)
	// 递归右子树查找节点
	right := lowestCommonAncestor(root.Right,p,q)
	// 如果两节点分别存在左右子树，则返回左右子树父节点
	if left != nil && right != nil{
		return root
	}
	// 两个节点都存在右子树
	if right != nil && left == nil{
		return right
	}
	// 两节点存在左子树
	return left
}

// 存储父节点 时间空间复杂度O（n）
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// 存储所有节点值对应父节点
	parent := map[int]*TreeNode{}
	// 存储访问过的节点
	visited := map[int]bool{}
	// 深度遍历所以节点
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil{
			return
		}
		if root.Left != nil{
			parent[root.Left.Val] = root
			dfs(root.Left)
		}
		if root.Right != nil{
			parent[root.Right.Val] = root
			dfs(root.Right)
		}
	}

	dfs(root)

	// 对p所有的父节点进行标识
	for p != nil{
		visited[p.Val] = true
		p = parent[p.Val]
	}
	// 寻找p 与 q共同的父节点
	for q != nil{
		if visited[q.Val]{
			return q
		}
		q = parent[q.Val]
	}

	return nil
}