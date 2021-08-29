package _01_SymmetricTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func isSymmetric(root *TreeNode) bool {
	// 从根节点出发，也可以忽略根节点
	return check(root, root)
}

func check(l, r *TreeNode) bool {
	// 两子树同时遍历完了说明对称
	if l == nil && r == nil {
		return true
	}
	// 其中一个子树先遍历完说明不对称
	if l == nil || r == nil {
		return false
	}
	// 判断当前两对称节点是否相等，递归判断剩余对称节点是否符合要求
	return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
}


// 迭代
func isSymmetric2(root *TreeNode) bool {
	// 初始化队列元素(根节点)，同样也可以忽略根节点
	// var l,r *TreeNode
	// var queue []*TreeNode
	// queue = append(queue,root.Left)
	// queue = append(queue,root.Right)
	l,r := root,root
	var queue []*TreeNode
	queue = append(queue,l)
	queue = append(queue,r)
	for len(queue) > 0{
		// 两两从队列取出元素
		l,r = queue[0],queue[1]
		queue = queue[2:]
		// 判断是否符合对称规则
		if l == nil && r == nil{
			continue
		}
		if l == nil || r == nil{
			return false
		}
		if l.Val != r.Val{
			return false
		}
		// 将下一轮要比较的对称节点放入队列
		queue = append(queue,l.Left)
		queue = append(queue,r.Right)

		queue = append(queue,l.Right)
		queue = append(queue,r.Left)
	}
	return true
}