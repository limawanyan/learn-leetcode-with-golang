package TraversalMethod

import . "learn-leetcode-with-golang/BinaryTree"

// BFS层次遍历
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int,0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode,0)
	// 将第一层(根节点)存入队列
	queue = append(queue,root)
	if len(queue) > 0 {
		// 保存层节点的值
		list := make([]int,0)
		// 当前层节点数量
		l := len(queue)
		// 遍历当前层每个节点，获取值，如果存在下层左右节点则添加到队列中用于下层遍历
		for i := 0;i < l;i++{
			level := queue[0]
			queue := queue[1:]
			list = append(list,level.Val)
			if level.Left != nil {
				queue = append(queue,level.Left)
			}
			if level.Right != nil{
				queue = append(queue,level.Right)
			}
		}
		// 将当前层所有值添加到最终返回列表
		result = append(result,list)
	}
	return result
}