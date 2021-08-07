package _26_InvertBinaryTree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归 DFS 后序遍历
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	//left := mirrorTree(root.Left)
	//right := mirrorTree(root.Right)
	//root.Left = right
	//root.Right = left
	root.Left,root.Right = mirrorTree(root.Right),mirrorTree(root.Left)
	return root
}

// 递归 DFS 前序遍历
func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	root.Left,root.Right = root.Right,root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

// 递归 DFS 中序遍历
func mirrorTree3(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	mirrorTree(root.Left)
	root.Left,root.Right = root.Right,root.Left
	// 这里的左节点已经被替换成了右节点
	mirrorTree(root.Left)
	return root
}

// 迭代 BFS
func mirrorTree4(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 队列
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		// 根节点出队
		node := queue[0]
		queue = queue[1:]
		// 反转左右子树
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}