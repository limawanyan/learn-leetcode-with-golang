package _09_MaximumDepthOfBinaryTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 深度优先
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	return max(maxDepth(root.Left),maxDepth(root.Right)) + 1
}

func max(a,b int) int {
	if a > b{
		return a
	}
	return b
}

// 广度优先
func maxDepth2(root *TreeNode) int {
	if root == nil{
		return 0
	}
	var result int
	// 队列存储节点
	queue := []*TreeNode{}
	// 初始化添加根节点
	queue = append(queue,root)
	for len(queue) > 0{
		ans := len(queue)
		// 节点出队列，添加左右子节点入队列
		for ans > 0{
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil{
				queue = append(queue,node.Left)
			}
			if node.Right != nil{
				queue = append(queue,node.Right)
			}
			ans--
		}
		// 一层所有节点出队列以后,深度++
		result++
	}
	return result
}