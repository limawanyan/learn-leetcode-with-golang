package TraversalMethod

import . "learn-leetcode-with-golang/BinaryTree"

// 中序非递归遍历 使用stack保存已访问的元素，用于原路返回
func inorderTraversal(root *TreeNode)  []int {
	result := make([]int,0)
	if root == nil{
		return result
	}
	stack := make([]*TreeNode,0)
	for len(stack) > 0 || root != nil {
		// 存储访问元素，一直向左
		for root != nil{
			stack = append(stack,root)
			root = root.Left
		}
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result,val.Val)
		root = val.Right
	}
	return result
}

// Morris中序遍历
func inorderTraversal2(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil{
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root{
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil{
				predecessor.Right = root
				root = root.Left
			}else{
				res = append(res,root.Val)
				predecessor.Right = nil
				root = root.Right
			}
		}else{
			res = append(res,root.Val)
			root = root.Right
		}
	}
	return
}