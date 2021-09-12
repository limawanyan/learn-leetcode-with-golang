package _02_BinaryTreeLevelOrderTraversal

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 广度优先BFS搜索 时间与空间复杂度O（n）
func levelOrder(root *TreeNode) [][]int {
	// 最终每层结果
	ret := [][]int{}
	if root == nil{
		return ret
	}
	// 待遍历层所有节点
	q := []*TreeNode{root}
	for i := 0;len(q) > 0;i++{
		// 初始化当前层
		ret = append(ret,[]int{})
		// 保存下一层节点
		p := []*TreeNode{}
		// 循环当前层所以节点获取值添加到返回结果
		for j := 0;j < len(q);j++{
			node := q[j]
			ret[i] = append(ret[i],node.Val)
			// 判断是否有子节点添加到下一层节点
			if node.Left != nil{
				p = append(p,node.Left)
			}
			if node.Right != nil{
				p = append(p,node.Right)
			}
		}
		// 将下一层节点变为待遍历层节点
		q = p
	}
	return ret
}

// DFS递归 时间复杂度O（n） 空间复杂度O（树的高度）
var res [][]int
func levelOrder2(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root,0)
	return res
}

func dfs(root *TreeNode,level int)  {
	if root != nil{
		if level == len(res){
			res = append(res,[]int{})
		}
		res[level] = append(res[level],root.Val)
		dfs(root.Left,level+1)
		dfs(root.Right,level+1)
	}
}

// 锯齿线层序遍历树
func zigzagLevelOrder(root *TreeNode) (res [][]int) {
	if root == nil{
		return res
	}
	queue := []*TreeNode{root}
	isLeft := true
	for len(queue) > 0{
		n := len(queue)
		valus := make([]int,n)
		for j := 0;j < n;j++{
			node := queue[j]
			if node.Left != nil{
				queue = append(queue,node.Left)
			}
			if node.Right != nil{
				queue = append(queue,node.Right)
			}
			if isLeft{
				valus[j] = node.Val
			}else{
				valus[n - j - 1] = node.Val
			}
		}
		queue = queue[n:]
		res = append(res,valus)
		isLeft = !isLeft
	}
	return res
}