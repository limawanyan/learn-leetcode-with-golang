package TraversalMethod

import . "learn-leetcode-with-golang/BinaryTree"

// 后序迭代遍历
func postorderTraversal(root *TreeNode)	[]int {
	if root == nil{
		return nil
	}
	result := make([]int,0)
	stack := make([]*TreeNode,0)
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil{
			stack = append(stack,root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		// 确保不存在右节点或右节点已经弹出后弹出根元素
		if node.Right == nil || node.Right == lastVisit{
			stack = stack[:len(stack)-1]
			result = append(result,node.Val)
			// 标识节点已弹出
			lastVisit = node
		}else{
			root = node.Right
		}
	}
	return result
}