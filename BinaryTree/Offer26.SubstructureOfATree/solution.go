package Offer26_SubstructureOfATree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil{
		return true
	}
	// 空二叉树不属于任何树子树
	if A == nil || B == nil{
		return false
	}
	//var ret bool
	//if A.Val == B.Val{
	//	ret = helper(A,B)
	//}
	//if !ret{
	//	ret = isSubStructure(A.Left,B)
	//}
	//if !ret{
	//	ret = isSubStructure(A.Right,B)
	//}
	//return ret
	// 判断两二叉树根节点是否相等，如果相等则继续递归左右判断是否包含整个子树
	// 如果不相等，递归左子树
	// 如果A左子树也没有包含B树的，递归右子树
	return helper(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
}

func helper(a *TreeNode,b *TreeNode) bool {
	// b树为空，说明遍历完了，A中包含子树B
	if b == nil{
		return true
	}
	// a树为空，说明A树遍历完了，但B树还有，所以A树不包含子树B
	if a == nil{
		return false
	}
	// a树和b树节点不相等
	if a.Val != b.Val{
		return false
	}
	// 如果相等，继续递归左右子树判断剩余部分是否符合
	return helper(a.Left,b.Left) && helper(a.Right,b.Right)
}