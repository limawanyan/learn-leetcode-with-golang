package reverse_linked_list

type ListNode struct {
	Val int
    Next *ListNode
}

// 迭代方式
func reverseList(head *ListNode) *ListNode {
	// 存储上一个节点，初始化为空
	var prev *ListNode
	// 当前节点
	curr := head
	for curr != nil {
		// 存储下一个节点
		next := curr.Next
		// 将当前节点指向上一个节点
		curr.Next = prev
		// 将当前节点作为上一个节点
		prev = curr
		// 将下一个节点作为当前节点
		curr = next

		// 赋值简写 curr.Next,prev,curr = prev,curr,curr.Next
	}
	// 返回结果为上一个节点变量，因为当前节点变量指向为空
	return prev
}

// 递归方式
func reverseList2(head *ListNode) *ListNode {
	// 终止条件，当前节点为空或指向到最后一个节点
	if head == nil || head.Next == nil{
		return head
	}
	// 递
	p := reverseList2(head.Next)
	// 归 将当前节点的下一个节点指向当前节点
	head.Next.Next = head
	// 将当前节点指向空
	head.Next = nil
	return p
}

// 递归方式2
func reverseList3(head *ListNode) *ListNode {
	return reverse(nil,head)
}

// 递归传递上一个节点和当前节点
func reverse(prev *ListNode,curr *ListNode) *ListNode{
	// 如果当前节点为空终止递归
	if curr == nil{
		return prev
	}
	// 保存下一个节点
	next := curr.Next
	// 将当前节点指向上一个节点
	curr.Next = prev
	return reverse(curr,next)
}