package _03_RemoveLinkedListElements

type ListNode struct {
	Val int
	Next *ListNode
}

// 递归 时间复杂度O(n),n为链表长度 空间复杂度O(n),空间复杂度主要取决于递归调用栈
func removeElements(head *ListNode,val int) *ListNode {
	// 终止条件
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next,val)
	// 如果当前节点置等于val，返回下一个节点
	if head.Val	== val {
		return head.Next
	}
	return head
}

// 迭代 时间复杂度O(n) 空间复杂度O(1)
func removeElements2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	// 迭代链表，指针为空时结束
	for temp := dummyHead;temp.Next != nil;{
		if temp.Next.Val == val{
			temp.Next = temp.Next.Next
		}else{
			temp = temp.Next
		}
	}
	return dummyHead.Next
}