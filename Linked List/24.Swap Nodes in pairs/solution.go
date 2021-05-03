package _4_Swap_Nodes_in_pairs

type ListNode struct {
     Val int
     Next *ListNode
}

// 递归法
func swapPairs(head *ListNode) *ListNode {
	// 当链表只剩一个节点或者没有节点的时候递归终止
	if head == nil || head.Next == nil{
		return head
	}
	// head next swapPairs
	newHead := head.Next
	// 当前节点指向已经反转链表
	head.Next = swapPairs(newHead.Next)
	// 当前节点的下一个节点指向当前节点
	newHead.Next = head
	// 返回处理好的节点
	return newHead
}

// 迭代法
func swapPairs2(head *ListNode) *ListNode{
	dummyHead := &ListNode{0,head}
	temp := dummyHead
	// 没有两两可以交换的节点时终止迭代
	for temp.Next != nil && temp.Next.Next != nil{
		// temp -> node1 -> node2
		node1 := temp.Next
		node2 := temp.Next.Next
		// temp -> node2 -> node1
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}