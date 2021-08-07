package _05_ContructBinaryTreeFormPreorderAndInorderTraversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// len(preorder) == 0 || len(inorder) == 0
	// 前序和中序序列个数始终是一样的，为空时结束
	if len(preorder) == 0 {
		return nil
	}
	// 从中序序列中找到根节点（前序序列的第一个元素）
	var root int
	for k,v := range inorder{
		if v == preorder[0]{
			root = k
			break
		}
	}
	// 递归左右子树
	return &TreeNode{
		Val: preorder[0],
		Left:buildTree(preorder[1:root+1],inorder[0:root]),
		Right: buildTree(preorder[root+1:],inorder[root+1:]),
	}
}

